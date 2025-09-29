package commands

import (
	"fmt"
	"gvm/pkg/config"
	"gvm/pkg/utils"
	"os"
	"path/filepath"
	"runtime"
)

// Install 安装指定版本的 Go
func Install(version string) {
	fmt.Printf("Downloading Go %s...\n", version)

	// 获取操作系统类型和架构
	osType := runtime.GOOS
	arch := runtime.GOARCH

	// 获取安装目录
	gvmRoot := config.GetGvmRoot()
	goVersionDir := filepath.Join(gvmRoot, "versions", version)
	zipFile := filepath.Join(goVersionDir, fmt.Sprintf("go-%s-%s-%s.zip", version, osType, arch))
	isDownload := false

	if err := os.MkdirAll(goVersionDir, 0755); err != nil {
		fmt.Printf("Create version directory failed: %v\n", err)
		return
	}

	// 检查版本是否已安装
	if _, err := os.Stat(zipFile); err == nil {
		fmt.Printf("Go %s is already download\n", version)
		isDownload = true
	}

	if !isDownload {
		// 构建下载 URL
		downloadURL := buildDownloadURL(version, osType, arch)
		if downloadURL == "" {
			fmt.Printf("Unsupported operating system or architecture: %s, %s\n", osType, arch)
			return
		}

		// 下载 Go 安装包
		fmt.Printf("Downloading from %s...\n", downloadURL)
		if err := utils.DownloadFile(downloadURL, zipFile); err != nil {
			fmt.Printf("Download failed: %v\n", err)
			return
		}
	}

	// 解压安装包
	fmt.Println("Unzipping installation package...")
	if err := utils.Unzip(zipFile, goVersionDir); err != nil {
		fmt.Printf("Unzip failed: %v\n", err)
		return
	}

	// 移动文件到安装目录
	goDir := filepath.Join(gvmRoot, "go")
	os.RemoveAll(goDir)
	if err := utils.CopyDirectory(goDir, goVersionDir); err != nil {
		fmt.Printf("Install failed: %v\n", err)
		return
	}

	fmt.Printf("Go %s installed successfully!\n", version)
}

// buildDownloadURL 根据版本、操作系统和架构构建下载 URL
func buildDownloadURL(version, osType, arch string) string {
	// Go 下载链接格式
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
