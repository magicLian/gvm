package commands

import (
	"fmt"
	"gvm/pkg/config"
	"gvm/pkg/utils"
	"os"
	"path/filepath"
	"runtime"
)

// Install Installs the specified version of Go.
func Install(version string) {
	fmt.Printf("Downloading Go %s...\n", version)

	// Get the operating system type and architecture.
	osType := runtime.GOOS
	arch := runtime.GOARCH

	// Get the installation directory.
	gvmRoot := config.GetGvmRoot()
	goVersionDir := filepath.Join(gvmRoot, "versions", version)
	zipFile := filepath.Join(goVersionDir, fmt.Sprintf("go-%s-%s-%s.zip", version, osType, arch))
	isDownload := false

	if err := os.MkdirAll(goVersionDir, 0755); err != nil {
		fmt.Printf("Create version directory failed: %v\n", err)
		return
	}

	// Check if the version is already installed.
	if _, err := os.Stat(zipFile); err == nil {
		fmt.Printf("Go %s is already download\n", version)
		isDownload = true
	}

	if !isDownload {
		// Build the download URL.
		downloadURL := buildDownloadURL(version, osType, arch)
		if downloadURL == "" {
			fmt.Printf("Unsupported operating system or architecture: %s, %s\n", osType, arch)
			return
		}

		// Download the Go installation package.
		fmt.Printf("Downloading from %s...\n", downloadURL)
		if err := utils.DownloadFile(downloadURL, zipFile); err != nil {
			fmt.Printf("Download failed: %v\n", err)
			return
		}
	}

	// Unzip the installation package.
	fmt.Println("Unzipping installation package...")
	if err := utils.Unzip(zipFile, goVersionDir); err != nil {
		fmt.Printf("Unzip failed: %v\n", err)
		return
	}

	// Move the files to the installation directory.
	goDir := filepath.Join(gvmRoot, "go")
	os.RemoveAll(goDir)
	if err := utils.CopyDirectory(goVersionDir, goDir); err != nil {
		fmt.Printf("Install failed: %v\n", err)
		return
	}

	fmt.Printf("Go %s installed successfully!\n", version)
}

// buildDownloadURL Builds download url
func buildDownloadURL(version, osType, arch string) string {
	baseURL := "https://golang.org/dl/go%s.%s-%s.%s"

	var osSuffix, archSuffix, extension string

	switch osType {
	case "windows":
		osSuffix = "windows"
		extension = "zip"
	case "darwin":
		osSuffix = "darwin"
		extension = "tar.gz"
	case "linux":
		osSuffix = "linux"
		extension = "tar.gz"
	default:
		return ""
	}

	switch arch {
	case "amd64", "x86_64":
		archSuffix = "amd64"
	case "arm64":
		archSuffix = "arm64"
	default:
		return ""
	}

	return fmt.Sprintf(baseURL, version, osSuffix, archSuffix, extension)
}
