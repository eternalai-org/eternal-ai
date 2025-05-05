package pkg

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"solo/internal/model"
	"strings"
	"time"
)

type AllowedCode struct {
	Code map[string]string
}

type RelyErrorMessage struct {
	Code    interface{} `json:"code"`
	Message *string     `json:"message"`
	Error   interface{} `json:"error"`
}

type RelyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAllowedCode() AllowedCode {
	ac := new(AllowedCode)
	codes := make(map[string]string)
	codes["200_ok"] = "200 OK"
	codes["201_created"] = "201 Created"
	codes["202_accepted"] = "202 Accepted"
	ac.Code = codes
	return *ac
}

func HttpRequest(fullUrl string, method string, headers map[string]string, reqBody interface{}) ([]byte, *http.Header, int, error) {
	fullUrl = strings.TrimSpace(fullUrl)
	bff := new(bytes.Buffer)
	if reqBody != nil {
		byteData, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, 0, err
		}
		bff = bytes.NewBuffer(byteData)
	}

	req, err := http.NewRequest(method, fullUrl, bff)
	if err != nil {
		return nil, nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, 0, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	isAllowed := isAllowed(res.Status)
	if !isAllowed {
		data := &RelyErrorMessage{}
		err = json.Unmarshal(body, data)
		if err != nil {
			return nil, nil, res.StatusCode, err
		}

		dataErrorString, ok := data.Error.(string)
		if ok {
			return nil, &res.Header, res.StatusCode, errors.New(dataErrorString)
		}

		dataError := &RelyError{}
		byteArray, err := json.Marshal(data.Error)
		if err != nil {
			return nil, nil, res.StatusCode, err
		}

		err = json.Unmarshal(byteArray, dataError)
		if err != nil {
			if data.Message != nil {
				return nil, &res.Header, res.StatusCode, errors.New(*data.Message)
			}
		}
		return nil, &res.Header, res.StatusCode, errors.New(dataError.Message)

	}

	if err != nil {
		return nil, nil, res.StatusCode, err
	}

	return body, &res.Header, res.StatusCode, nil
}

func HttpStreamRequest(fullUrl string, method string, headers map[string]string, reqBody interface{}, strChan chan model.StreamingData) (*model.LLMInferResponse, *http.Header, int, error) {
	dataResp := &model.LLMInferResponse{}
	dataResp.Choices = make([]model.LLMInferChoice, 1)
	var err error

	defer func() {
		if err != nil {
			strChan <- model.StreamingData{
				Data:     &model.LLMInferStreamResponse{},
				Stop:     true,
				Err:      err,
				StreamID: -1,
			}

			return
		}

		strChan <- model.StreamingData{
			Data:     &model.LLMInferStreamResponse{},
			Stop:     true,
			Err:      nil,
			StreamID: -1,
		}
		close(strChan)
	}()

	fullUrl = strings.TrimSpace(fullUrl)
	bff := new(bytes.Buffer)
	if reqBody != nil {
		byteData, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, 0, err
		}
		bff = bytes.NewBuffer(byteData)
	}

	req, err := http.NewRequest(method, fullUrl, bff)
	if err != nil {
		return nil, nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{
		Timeout: 120 * time.Second, // Set the timeout duration
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, 0, err
	}

	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Error: received status code %d", resp.StatusCode))
		return nil, nil, 0, err
	}

	// Read the response body in a buffered manner
	scanner := bufio.NewScanner(resp.Body)
	msg := ""
	index := 0
	for scanner.Scan() {
		streamData := scanner.Text()
		streamData = strings.ReplaceAll(streamData, "data: ", "")
		if streamData != "[DONE]" && streamData != "" {
			_dataResp := &model.LLMInferStreamResponse{}
			err = json.Unmarshal([]byte(streamData), _dataResp)
			if err != nil {
				fmt.Println("[Err]", err)
				return nil, nil, 0, err
			}

			if len(_dataResp.Choices) > 0 {
				msgToChan := _dataResp.Choices[0].Delta.Content
				msg += msgToChan
				strChan <- model.StreamingData{
					StreamID: index,
					Data:     _dataResp,
					Stop:     false,
					Err:      nil,
				}
			}

			if index == 0 {
				dataResp.Id = _dataResp.Id
				dataResp.Object = _dataResp.Object
				dataResp.Created = _dataResp.Created
				dataResp.Model = _dataResp.Model
				dataResp.Choices[0].Message.Role = _dataResp.Choices[0].Delta.Role
			}

			index++

			// Process each line of the response
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
		return nil, nil, 0, err
	}

	// Sleep for a while before the next request
	//time.Sleep(5 * time.Second) // Adjust as needed
	dataResp.Choices[0].Message.Content = msg

	return dataResp, &resp.Header, resp.StatusCode, nil
}

func HttpRequestFullResponse(fullUrl string, method string, headers map[string]string, reqBody interface{}) ([]byte, *http.Header, int, error) {
	fullUrl = strings.TrimSpace(fullUrl)
	bff := new(bytes.Buffer)
	if reqBody != nil {
		byteData, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, 0, err
		}
		bff = bytes.NewBuffer(byteData)
	}

	req, err := http.NewRequest(method, fullUrl, bff)
	if err != nil {
		return nil, nil, 0, err
	}

	if len(headers) > 0 {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, 0, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, &res.Header, res.StatusCode, err
	}
	return body, &res.Header, res.StatusCode, nil
}

func JsonRequest(fullUrl string, method string, headers map[string]string, reqBody io.Reader) ([]byte, *http.Header, int, error) {
	// headers["accept"] = "application/json"
	// headers["content-type"] = "application/json"

	return HttpRequest(fullUrl, method, headers, reqBody)
}

func isAllowed(code string) bool {
	ac := NewAllowedCode()
	code = strings.ReplaceAll(code, " ", "_")
	code = strings.ToLower(code)
	getCode, ok := ac.Code[code]
	if !ok {
		return false
	}

	_ = getCode
	return true
}
