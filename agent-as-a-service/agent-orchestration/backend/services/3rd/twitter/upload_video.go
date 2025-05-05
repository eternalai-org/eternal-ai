package twitter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/oauth1"
)

// UploadVideo uploads the video in chunks to Twitter
func (c *Client) UploadVideo(videoURL string, additionalOwners []string) (string, error) {
	if videoURL == "" {
		return "", fmt.Errorf("video URL is empty")
	}

	// Open the video file (download from URL if necessary)
	response, err := http.Get(videoURL)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Get the total size of the video file
	totalBytes := response.ContentLength

	// OAuth1 authentication
	config := oauth1.NewConfig(c.ConsumerKey, c.ConsumerSecret)
	token := oauth1.NewToken(c.AccessToken, c.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Step 1: Initialize the upload session
	mediaID, err := c.uploadInit(httpClient, totalBytes, additionalOwners)
	if err != nil {
		return "", err
	}

	// Step 2: Upload video chunks
	err = c.uploadAppend(httpClient, mediaID, response.Body, totalBytes)
	if err != nil {
		return "", err
	}

	// Step 3: Finalize the upload process
	err = c.uploadFinalize(httpClient, mediaID)
	if err != nil {
		return "", err
	}

	return mediaID, nil
}

// uploadInit starts the video upload session
func (c *Client) uploadInit(httpClient *http.Client, totalBytes int64, additionalOwners []string) (string, error) {
	url := "https://upload.twitter.com/1.1/media/upload.json"
	data := make(map[string]string)
	data["command"] = "INIT"
	data["media_type"] = "video/mp4"
	data["total_bytes"] = fmt.Sprintf("%d", totalBytes)
	data["media_category"] = "tweet_video"

	// Create the request body for initialization
	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	for key, value := range data {
		_ = writer.WriteField(key, value)
	}
	if len(additionalOwners) > 0 {
		owners := strings.Join(additionalOwners, ",")
		if err := writer.WriteField("additional_owners", owners); err != nil {
			return "", fmt.Errorf("failed to add additional_owners: %w", err)
		}
	}
	writer.Close()

	// Prepare the POST request
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Decode the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// Extract media ID
	mediaID, ok := result["media_id_string"].(string)
	if !ok {
		return "", fmt.Errorf("failed to find 'data' field in response")
	}

	return mediaID, nil
}

// uploadAppend uploads video chunks to Twitter
func (c *Client) uploadAppend(httpClient *http.Client, mediaID string, videoBody io.Reader, totalBytes int64) error {
	// Define chunk size (4MB per chunk)
	const chunkSize = 4 * 1024 * 1024
	segmentID := 0
	bytesSent := int64(0)

	// Read and send chunks
	for {
		chunk := make([]byte, chunkSize)
		n, err := videoBody.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Slice chunk based on the number of bytes read
		chunk = chunk[:n]

		// Prepare the multipart form for uploading a chunk
		reqBody := new(bytes.Buffer)
		writer := multipart.NewWriter(reqBody)
		part, err := writer.CreateFormFile("media", "video_chunk")
		if err != nil {
			return err
		}

		_, err = part.Write(chunk)
		if err != nil {
			return err
		}

		// Add media_id and segment_index to the form
		writer.WriteField("command", "APPEND")
		writer.WriteField("media_id", mediaID)
		writer.WriteField("segment_index", fmt.Sprintf("%d", segmentID))
		writer.Close()

		// Prepare the POST request to upload the chunk
		req, err := http.NewRequest("POST", "https://upload.twitter.com/1.1/media/upload.json", reqBody)
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		// Send the chunk
		resp, err := httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Update bytes sent and segment ID
		bytesSent += int64(n)
		segmentID++

		// Print progress
		fmt.Printf("Uploaded %d/%d bytes\n", bytesSent, totalBytes)
	}

	return nil
}

// uploadFinalize finalizes the upload process and starts video processing
func (c *Client) uploadFinalize(httpClient *http.Client, mediaID string) error {
	url := "https://upload.twitter.com/1.1/media/upload.json"
	data := make(map[string]string)
	data["command"] = "FINALIZE"
	data["media_id"] = mediaID

	// Create the request body for finalization
	reqBody := new(bytes.Buffer)
	writer := multipart.NewWriter(reqBody)
	for key, value := range data {
		_ = writer.WriteField(key, value)
	}
	writer.Close()

	// Prepare the POST request
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Decode the response
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	// Check if the video has finished processing
	processingInfo, ok := result["processing_info"].(map[string]interface{})
	if ok {
		state := processingInfo["state"].(string)
		if state == "succeeded" {
			return nil
		}

		// If the video is still processing, wait and try again
		checkAfterSecs := processingInfo["check_after_secs"].(float64)
		fmt.Printf("Processing video, checking after %.0f seconds...\n", checkAfterSecs)
		time.Sleep(time.Duration(checkAfterSecs) * time.Second)
		return c.uploadFinalize(httpClient, mediaID)
	}

	return fmt.Errorf("failed to finalize upload")
}
