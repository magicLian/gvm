package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
)

// DownloadFile Download file from url to destPath with progress bar
func DownloadFile(fileUrl, destPath string) error {
	// Ensure the directory exists
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("create directory failed: %v", err)
	}

	client := &http.Client{
		Timeout:   60 * time.Second,
	}
	resp, err := client.Get(fileUrl)
	if err != nil {
		return fmt.Errorf("send HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: status code %d", resp.StatusCode)
	}

	// Create progress bar
	bar := progressbar.NewOptions64(
		resp.ContentLength,
		progressbar.OptionSetDescription("Downloading..."),
		progressbar.OptionSetWidth(30),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// Create file
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}
	defer out.Close()

	// Write response content to file + update progress bar
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return fmt.Errorf("write response content failed: %v", err)
	}

	fmt.Println("\n✅ Download completed:", destPath)
	return nil
}
