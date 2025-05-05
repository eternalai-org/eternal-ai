package zip_hf_model_to_light_house

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"solo/pkg/lighthouse"
	"strconv"
	"time"
)

var (
	ZipChunkSize = 200 //MB
	BASH_EXEC    = "/usr/bin/bash"
	THREADS      = runtime.NumCPU() - 1
)

type HFModelZipFile struct {
	File string `json:"file"`
	Hash string `json:"hash"`
}

type HFModelInLightHouse struct {
	Model     string           `json:"model"`
	NumOfFile int              `json:"num_of_file"`
	Files     []HFModelZipFile `json:"files"`
}

func ExecuteCommand(fileCmd string) ([]byte, error) {
	commandId := strconv.FormatInt(time.Now().UnixMicro(), 10)
	fileLog := fmt.Sprintf("/tmp/log_%v.txt", commandId)
	execCmd := fmt.Sprintf("%v %v  2>&1 | /usr/bin/tee %v", BASH_EXEC, fileCmd, fileLog)
	fileExec := fmt.Sprintf("/tmp/bash_%v.sh", commandId)

	if err := os.WriteFile(fileExec, []byte(execCmd), 0644); err != nil {
		return nil, err
	}

	command := exec.Command(BASH_EXEC, fileExec)
	out, err := command.Output()
	if err != nil {
		return out, err
	}
	return os.ReadFile(fileLog)
}

func getScriptZipFile(modelFolder string, hfDir string) (string, error) {
	filePath := fmt.Sprintf("/tmp/hf-zip-model-%v.sh", modelFolder)
	if err := removeFileIfExists(filePath); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	commands := []string{
		fmt.Sprintf("cd %v", hfDir),
		fmt.Sprintf("sudo rm -Rf %v.zip.part-*", modelFolder),
		fmt.Sprintf("sudo tar -cf - %v | sudo pigz --best -p %v | sudo split -b %vM - %v.zip.part-", modelFolder, THREADS, ZipChunkSize, modelFolder),
	}

	for _, cmd := range commands {
		if _, err := file.WriteString(cmd + " \n "); err != nil {
			return "", fmt.Errorf("error writing to file: %v", err)
		}
	}

	return filePath, nil
}

func removeFileIfExists(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("error removing file: %v", err)
		}
	}
	return nil
}

func getScriptUnZipFile(modelFolder string, hfDir string) (string, error) {
	model := fmt.Sprintf("hf-unzip-model-%v.sh", modelFolder)
	filePath := filepath.Join("/tmp", model)
	if err := removeFileIfExists(filePath); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file:%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("cd %v \n ", hfDir))
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}

	_, err = file.WriteString(fmt.Sprintf("sudo cat %v.zip.part-* | sudo pigz -p %v -d | sudo tar -xf -", modelFolder, 2))
	if err != nil {
		return "", fmt.Errorf("error write file:%v", err)
	}
	return filePath, nil
}

func getListZipFile(modelFolder string, hfDir string) ([]string, error) {
	filePath := fmt.Sprintf("/tmp/list-zip-model-%v.sh", modelFolder)
	if err := removeFileIfExists(filePath); err != nil {
		return nil, err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("error creating file:%v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("rm /tmp/list_file_%v.txt \n", modelFolder))
	if err != nil {
		return nil, fmt.Errorf("error write file:%v", err)
	}

	_, err = file.WriteString(fmt.Sprintf("cd %v \n", hfDir))
	if err != nil {
		return nil, fmt.Errorf("error write file:%v", err)
	}

	_, err = file.WriteString(fmt.Sprintf("sudo ls %v.zip.part-* > /tmp/list_file_%v.txt \n", modelFolder, modelFolder))
	if err != nil {
		return nil, fmt.Errorf("error write file:%v", err)
	}

	output, err := ExecuteCommand(fmt.Sprintf("/tmp/list-zip-model-%v.sh ", modelFolder))
	if err != nil {
		return nil, fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	file, err = os.Open(fmt.Sprintf("/tmp/list_file_%v.txt", modelFolder))
	if err != nil {
		return nil, fmt.Errorf("error opening file:%v", err)
	}

	scanner := bufio.NewScanner(file)
	listFile := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		listFile = append(listFile, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file:,%v", err)
	}
	return listFile, nil
}

func uploadListZipFileToLightHouse(modelFolder string, hfDir string, apiKey string) (*HFModelInLightHouse, error) {
	listFile, err := getListZipFile(modelFolder, hfDir)
	if err != nil {
		return nil, err
	}

	if len(listFile) == 0 {
		return nil, fmt.Errorf("no files pattern %v.zip.part-*  found in folder %v", modelFolder, hfDir)
	}

	result := HFModelInLightHouse{
		Model:     modelFolder,
		NumOfFile: len(listFile),
		Files:     make([]HFModelZipFile, 0),
	}

	for i, file := range listFile {
		log.Println("Start upload model ", modelFolder, "chunk", i, "file", file)
		for j := 0; j < 10; i++ {
			cid, err := lighthouse.UploadFile(apiKey, file, fmt.Sprintf("%v/%v", hfDir, file))
			if err != nil {
				log.Println("Error when upload model ", modelFolder, "retry", j, "chunk", i, "file", file, "err", err)
				time.Sleep(2 * time.Minute)
				continue
			} else {
				log.Println("Finish upload model ", modelFolder, "chunk", i, "file", file, "==> hash", cid)
				result.Files = append(result.Files, HFModelZipFile{File: file, Hash: cid})
				break
			}
		}
	}

	return &result, nil
}

func uploadHFModelResultToLightHouse(info *HFModelInLightHouse, apiKey string) (string, error) {
	data, _ := json.Marshal(info)
	cid, err := lighthouse.UploadData(apiKey, info.Model, data)
	if err != nil {
		return "", err
	}
	return cid, nil
}

func getHFModelResultFromLightHouse(hash string) (*HFModelInLightHouse, error) {
	data, _, err := lighthouse.DownloadDataSimple(hash)
	if err != nil {
		return nil, err
	}

	result := &HFModelInLightHouse{}
	if err = json.Unmarshal(data, result); err != nil {
		return nil, err
	}
	return result, nil
}

func downloadZipFileFromLightHouse(info *HFModelInLightHouse, hfDir string) error {
	for _, file := range info.Files {
		log.Println("Start download ", "file", file.File, "hash", file.Hash, "hfDir", hfDir)
		for {
			filePath := filepath.Join(hfDir, file.File)
			err := lighthouse.DownloadToFile(file.Hash, filePath)
			if err != nil {
				log.Println("Error when try down file from light house", "file", file.File, "hash", file.Hash, "err", err.Error())
				time.Sleep(2 * time.Minute)
				continue
			}
			break
		}
	}
	log.Println("Success download all zip file:", "model", info.Model)
	return nil
}

func DownloadHFModelFromLightHouse(hash string, hfDir string) error {
	info, err := getHFModelResultFromLightHouse(hash)
	if err != nil {
		return fmt.Errorf("error when get model info from light house hash : %v err :%v ", hash, err)
	}

	if err = downloadZipFileFromLightHouse(info, hfDir); err != nil {
		return fmt.Errorf("error when download zip chunk file:%v ", err)
	}

	scriptFile, err := getScriptUnZipFile(info.Model, hfDir)
	if err != nil {
		return fmt.Errorf("error when get unzip script file:%v ", err)
	}
	log.Println("Start unzip list files")

	output, err := ExecuteCommand(scriptFile)
	if err != nil {
		return fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	log.Println("Success unzip model ", info.Model)
	unzipFolder := filepath.Join(hfDir, info.Model)
	files, err := os.ReadDir(filepath.Join(hfDir, info.Model))
	if err != nil {
		return fmt.Errorf("error when read dir:%v , err:%v", unzipFolder, err.Error())
	}

	for _, file := range files {
		fmt.Printf("%s/%s\n", info.Model, file.Name())
	}
	return nil
}

func ZipAndUploadHFModelFromLightHouse(modelFolder string, hfDir string, apiKey string) (string, error) {
	scriptFile, err := getScriptZipFile(modelFolder, hfDir)
	if err != nil {
		return "", fmt.Errorf("error when get script zip file:%v ", err)
	}

	log.Println("Start compress model")
	output, err := ExecuteCommand(scriptFile)
	if err != nil {
		return "", fmt.Errorf("error when execute file:%v , output:%v", err, string(output))
	}

	log.Println("Finish compress model . Start upload model")
	result, err := uploadListZipFileToLightHouse(modelFolder, hfDir, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload list zip file to light house :%v ", err)
	}

	hash, err := uploadHFModelResultToLightHouse(result, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload model result to light house :%v ", err)
	}
	return hash, nil
}

func UploadHFModelFromLightHouse(modelFolder string, hfDir string, apiKey string) (string, error) {
	log.Println("Start upload model")
	result, err := uploadListZipFileToLightHouse(modelFolder, hfDir, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload list zip file to light house :%v ", err)
	}

	hash, err := uploadHFModelResultToLightHouse(result, apiKey)
	if err != nil {
		return "", fmt.Errorf("error when upload model result to light house :%v ", err)
	}
	return hash, nil
}
