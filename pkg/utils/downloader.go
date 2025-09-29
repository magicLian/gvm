package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile Download file from specified URL.
func DownloadFile(url, filepath string) error {
	// Create file.
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}
	defer out.Close()

	// Send HTTP request.
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("send HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check response status.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: Status code %d", resp.StatusCode)
	}

	// Write response content to file.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("write response content to file failed: %v", err)
	}
	return nil
}
