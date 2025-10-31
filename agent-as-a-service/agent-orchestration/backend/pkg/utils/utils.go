package utils

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Environment string

const Production Environment = "production"

func IsEnvProduction(env string) bool {
	return env == string(Production)
}

func CheckRedirect(link string) (string, bool, error) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Prevents automatic redirects, we'll handle them manually
		},
	}

	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		return "", false, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 && resp.StatusCode < 400 { // Check for redirect status codes
		redirectURL, err := resp.Location()
		if err != nil {
			return "", false, err
		}
		return redirectURL.String(), true, nil
	}

	return "", false, nil // No redirect
}

func IsDownloadLink(link string) (bool, error) {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return false, err
	}

	resp, err := http.Head(parsedURL.String()) // Use HEAD to avoid downloading the whole file
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// contentType := resp.Header.Get("Content-Type")
	// contentDisposition := resp.Header.Get("Content-Disposition")

	// Check Content-Type (common download types)
	// if strings.HasPrefix(contentType, "application/") ||
	// 	strings.HasPrefix(contentType, "audio/") ||
	// 	strings.HasPrefix(contentType, "video/") ||
	// 	strings.HasPrefix(contentType, "image/") ||
	// 	contentType == "application/octet-stream" { // generic binary download
	// 	return true, nil
	// }

	// Check Content-Disposition (indicates attachment/download)
	// if strings.Contains(contentDisposition, "attachment") {
	// 	return true, nil
	// }

	// Check for common file extensions in the URL path.
	// if strings.HasSuffix(parsedURL.Path, ".zip") ||
	// 	strings.HasSuffix(parsedURL.Path, ".rar") ||
	// 	strings.HasSuffix(parsedURL.Path, ".exe") ||
	// 	strings.HasSuffix(parsedURL.Path, ".dmg") ||
	// 	strings.HasSuffix(parsedURL.Path, ".pdf") ||
	// 	strings.HasSuffix(parsedURL.Path, ".mp3") ||
	// 	strings.HasSuffix(parsedURL.Path, ".mp4") ||
	// 	strings.HasSuffix(parsedURL.Path, ".jpg") ||
	// 	strings.HasSuffix(parsedURL.Path, ".avif") ||
	// 	strings.HasSuffix(parsedURL.Path, ".png") {
	// 	return true, nil
	// }
	if strings.HasSuffix(parsedURL.Path, ".pdf") {
		return true, nil
	}

	return false, nil
}

func FollowRedirects(urlStr string) (string, error) {
	client := &http.Client{} // default client follows redirects
	resp, err := client.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Request.URL.String(), nil
}

type FileInfo struct {
	Size int64
	Name string
}

func getFileNameFromHeaders(headers http.Header, urlStr string) (string, error) {
	contentDisposition := headers.Get("Content-Disposition")
	if contentDisposition != "" {
		// Attempt to extract filename from Content-Disposition
		parts := strings.Split(contentDisposition, ";")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if strings.HasPrefix(part, "filename=") {
				filename := strings.Trim(part[len("filename="):], `"`) // Remove quotes
				return filename, nil
			}
		}
	}

	// Fallback: extract filename from URL path
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	return path.Base(parsedURL.Path), nil
}

func GetFileInfoByURL(urlStr string) (*FileInfo, error) {
	resp, err := http.Head(urlStr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	contentLength := resp.ContentLength
	if contentLength < 0 {
		return nil, fmt.Errorf("content length not available")
	}

	filename, err := getFileNameFromHeaders(resp.Header, urlStr)
	if err != nil {
		return nil, err
	}

	return &FileInfo{
		Size: contentLength,
		Name: filename,
	}, nil
}

func GetFileSizeByURL(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	contentLength := resp.ContentLength
	if contentLength < 0 {
		return 0, fmt.Errorf("content length not available")
	}

	return contentLength, nil
}

func CreateValidLink(domainURL, link string) (string, error) {
	parsedLink, err := url.Parse(link)
	if err != nil {
		return "", fmt.Errorf("invalid link: %w", err)
	}
	parsedLink.Fragment = ""

	if parsedLink.Scheme != "" && parsedLink.Host != "" {
		return link, nil
	}

	parsedDomain, err := url.Parse(domainURL)
	if err != nil {
		return "", fmt.Errorf("invalid domain URL: %w", err)
	}

	if parsedDomain.Scheme == "" {
		return "", fmt.Errorf("domain URL must include a scheme (http or https)")
	}

	if strings.HasPrefix(link, "/") || !strings.Contains(link, "//") {
		return parsedDomain.Scheme + "://" + parsedDomain.Host + parsedLink.Path + "?" + parsedLink.RawQuery + parsedLink.Fragment, nil
	}

	return parsedDomain.Scheme + "://" + parsedDomain.Host + "/" + link, nil
}

func RemoveFragment(urlString string) (string, error) {
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}
	u.Fragment = "" // Remove the fragment (the part after '#')
	return u.String(), nil
}

func filenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func CreateFolderIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		log.Printf("Created directory: %s", path)
	} else if err != nil {
		return err
	} else {
		log.Printf("Directory %s already exists", path)
	}
	return nil
}

func DownloadFileByUrl(link string, localPath string) (string, string, error) {
	parsedUrl, err := url.ParseRequestURI(link)
	if err != nil {
		return "", "", err
	}

	// Get the data
	client := &http.Client{}
	req, _ := http.NewRequest("GET", link, nil)

	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", "", errors.New("received non 200 response code")
	}

	fileName := strings.Replace(resp.Header.Get("Content-Disposition"), "attachment; filename=", "", -1)
	fileName = strings.Replace(fileName, "attachment;filename=", "", -1)
	if fileName == "" {
		filePath := parsedUrl.Path
		fileName = path.Base(filePath)
	}

	_ = CreateFolderIfNotExists(filepath.Join(localPath, filenameWithoutExtension(fileName)))

	pathFile := filepath.Join(localPath, filenameWithoutExtension(fileName), fileName)
	// Create the file
	out, err := os.Create(pathFile)
	if err != nil {
		return "", "", err
	}
	defer out.Close()
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", "", err
	}

	return pathFile, fileName, nil
}
