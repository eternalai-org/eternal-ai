package lighthouse

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-datastore"

	"github.com/gabriel-vasile/mimetype"
	"github.com/ipfs/boxo/blockservice"
	blockstore "github.com/ipfs/boxo/blockstore"
	chunker "github.com/ipfs/boxo/chunker"
	offline "github.com/ipfs/boxo/exchange/offline"
	"github.com/ipfs/boxo/ipld/merkledag"
	"github.com/ipfs/boxo/ipld/unixfs/importer/balanced"
	uih "github.com/ipfs/boxo/ipld/unixfs/importer/helpers"
	dsync "github.com/ipfs/go-datastore/sync"
	"github.com/pkg/errors"
)

type LightHouseResponse struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

const (
	IPFSGateway = "https://gateway.lighthouse.storage/ipfs/"
)

func DownloadDataSimple(hash string) ([]byte, string, error) {
	if strings.Contains(hash, "ipfs://") {
		hash = strings.Replace(hash, "ipfs://", "", 1)
	}
	// https://gateway.lighthouse.storage/ipfs/QmXPGcEHCi1ZmbHFwScuP4ZJ2iv9YjTJMUroUTJUnFXxxj
	urlLink := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash)

	resp, err := http.Get(urlLink)
	if err != nil {
		return nil, "", fmt.Errorf("error when try get reponse :%v", err)
	}
	if resp == nil {
		return nil, "", fmt.Errorf("error when try get reponse ==nil")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("error when try get data url :%v => reponse code :%v", urlLink, resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("error when read body from  reponse :%v", err)
	}

	mtype := mimetype.Detect(body)

	return body, mtype.String(), nil
}

func DownloadDataSimpleWithRetry(hash string) ([]byte, string, error) {
	var err error
	var byteResp []byte
	var mimeType string
	for i := 0; i < 3; i++ {
		byteResp, mimeType, err = DownloadDataSimple(hash)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		return byteResp, mimeType, nil
	}

	return byteResp, mimeType, nil
}

func downloadChunkedData(hash string, start, end int) ([]byte, error) {
	urlLink := fmt.Sprintf("https://gateway.lighthouse.storage/ipfs/%s", hash)

	req, err := http.NewRequest("GET", urlLink, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}

func DownloadChunkedData(hash string, modelDir string) (string, error) {
	// check folder exist first, if not create
	err := os.MkdirAll(modelDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileInfo, err := GetFileInfo(hash)
	if err != nil {
		return "", err
	}

	fileSize, err := strconv.ParseInt(fileInfo.FileSizeInBytes, 10, 64)
	if err != nil {
		return "", err
	}

	filePath := path.Join(modelDir, fmt.Sprintf("model.zip"))
	if _, err := os.Stat(filePath); err == nil {
		return filePath, nil
	}

	var chunkSize int64 = 1024 * 1024 * 10 // 10MB
	var start int64 = 0
	var end int64 = chunkSize

	parts := fileSize / chunkSize
	log.Println("parts: ", parts)
	part := 0
	for start < fileSize {
		if end > fileSize {
			end = fileSize
		}
		log.Println("part: ", part, "start: ", start, "end: ", end)
		data, err := downloadChunkedData(hash, int(start), int(end))
		if err != nil {
			return "", err
		}

		err = WriteFile(filePath, data, os.ModePerm)
		if err != nil {
			return "", err
		}
		start = end + 1
		end += chunkSize
		part++
	}
	return filePath, nil
}

func WriteFile(name string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}

type FileInfo struct {
	FileSizeInBytes string `json:"fileSizeInBytes"`
	Cid             string `json:"cid"`
	Encryption      bool   `json:"encryption"`
	FileName        string `json:"fileName"`
	MimeType        bool   `json:"mimeType"`
	TxHash          string `json:"txHash"`
	Error           struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func GetFileInfo(hash string) (*FileInfo, error) {
	urlLink := fmt.Sprintf("https://api.lighthouse.storage/api/lighthouse/file_info?cid=%s", hash)

	resp, err := http.Get(urlLink)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var respBody FileInfo

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if respBody.Error.Code != 0 {
		return &respBody, errors.New(respBody.Error.Message)
	}

	return &respBody, nil
}

func UploadDataWithRetry(apikey, fileName string, data []byte) (string, error) {
	var err error
	var hash string
	for i := 0; i < 3; i++ {
		hash, err = UploadData(apikey, fileName, data)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		return hash, nil
	}
	return hash, err
}

func UploadData(apikey, fileName string, data []byte) (string, error) {
	urlLink := "https://node.lighthouse.storage/api/v0/add"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}
	if _, err = fw.Write(data); err != nil {
		return "", err
	}

	w.Close()

	req, err := http.NewRequest(
		"POST",
		urlLink,
		&b,
	)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", w.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apikey))

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var respBody LightHouseResponse

	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", err
	}

	return respBody.Hash, nil
}

func getCurrentDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func Cid(data []byte) string {
	ds := dsync.MutexWrap(datastore.NewNullDatastore())
	bs := blockstore.NewBlockstore(ds)
	bs = blockstore.NewIdStore(bs)
	bsrv := blockservice.New(bs, offline.Exchange(bs))
	dsrv := merkledag.NewDAGService(bsrv)
	ufsImportParams := uih.DagBuilderParams{
		Maxlinks:   uih.BlockSizeLimit, // Default max of 174 links per block
		RawLeaves:  false,
		CidBuilder: cid.V0Builder{},
		Dagserv:    dsrv,
		NoCopy:     false,
	}
	reader := bytes.NewReader(data)
	ufsBuilder, err := ufsImportParams.New(chunker.NewSizeSplitter(reader, chunker.DefaultBlockSize)) // 256KiB chunks
	if err != nil {
		return cid.Undef.String()
	}
	nd, err := balanced.Layout(ufsBuilder)
	if err != nil {
		return cid.Undef.String()
	}
	return nd.Cid().String()
}

func fileExistOnNetwork(data []byte) (string, bool, error) {
	cidCompute := Cid(data)
	if cidCompute == "" {
		return "", false, nil
	}
	fileInfo, err := GetFileInfo(cidCompute)
	if err != nil {
		if fileInfo != nil {
			if fileInfo.Error.Code == 404 {
				return "", false, nil
			}
		}
		return "", false, err
	}

	return cidCompute, true, nil
}

func UploadFile(apikey, fileName string, filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	/*cid, exist, err := fileExistOnNetwork(data)
	if err != nil {
		return "", err
	}
	if exist {
		return cid, nil
	}*/

	urlLink := "https://node.lighthouse.storage/api/v0/add"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fw, err := w.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}
	if _, err = fw.Write(data); err != nil {
		return "", err
	}

	w.Close()

	req, err := http.NewRequest(
		"POST",
		urlLink,
		&b,
	)
	if err != nil {
		return "", fmt.Errorf("err when init upload request err:%v", err)
	}

	client := &http.Client{}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", w.FormDataContentType())
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apikey))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("err when execute upload request err:%v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("err when read body err:%v", err)
	}
	var respBody LightHouseResponse
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return "", fmt.Errorf("err when parse json body:%v,err:%v", body, err)
	}

	return respBody.Hash, nil
}

func DownloadToFile(hash string, filePath string) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := os.Stat(filePath); err == nil {
		return nil
	}
	data, _, err := DownloadDataSimple(hash)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		f.Close()
		os.Remove(filePath)
		return err
	}
	err = f.Close()
	if err != nil {
		os.Remove(filePath)
	}
	return nil
}
