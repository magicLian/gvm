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
	goArchivesDir := filepath.Join(gvmRoot, "archives")
	goCurrentDir := filepath.Join(gvmRoot, "current")
	zipFile := filepath.Join(goArchivesDir, buildDownloadFileName(version, osType, arch))
	isDownload := false

	if err := os.MkdirAll(goArchivesDir, 0755); err != nil {
		fmt.Printf("Create archives directory failed: %v\n", err)
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

	// Create current symlink.
	if err := os.RemoveAll(goCurrentDir); err != nil {
		fmt.Printf("Remove current symlink failed: %v\n", err)
		return
	}
	goVersionGODir := filepath.Join(goVersionDir, "go")
	if err := os.Symlink(goVersionGODir, goCurrentDir); err != nil {
		fmt.Printf("Create current symlink failed: %v\n", err)
		return
	}

	fmt.Printf("Go %s installed successfully!\n", version)
}

// buildDownloadFileName Builds download file name
func buildDownloadFileName(version, osType, arch string) string {
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

	return fmt.Sprintf("go%s.%s-%s.%s", version, osSuffix, archSuffix, extension)
}

// buildDownloadURL Builds download url
func buildDownloadURL(version, osType, arch string) string {
	baseURL := "https://mirrors.aliyun.com/golang/%s"
	downloadFileName := buildDownloadFileName(version, osType, arch)
	if downloadFileName == "" {
		return ""
	}
	return fmt.Sprintf(baseURL, downloadFileName)
}
