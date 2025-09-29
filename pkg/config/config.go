package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetGvmRoot 获取 GVM 根目录
func GetGvmRoot() string {
	var gvmRoot string

	// 从环境变量获取 GVM 根目录
	gvmRoot = os.Getenv("GVM_ROOT")
	if gvmRoot == "" {
		// 如果环境变量不存在，根据操作系统确定 GVM 根目录
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Get user home directory failed: %v\n", err)
			os.Exit(1)
		}

		switch runtime.GOOS {
		case "windows":
			// Windows 上使用 APPDATA 环境变量
			appData := os.Getenv("APPDATA")
			if appData != "" {
				gvmRoot = filepath.Join(appData, "gvm")
			} else {
				// 如果 APPDATA 不存在，使用用户主目录
				gvmRoot = filepath.Join(homeDir, "gvm")
			}
		default:
			// Unix-like 系统上使用 .gvm 目录
			gvmRoot = filepath.Join(homeDir, ".gvm")
		}
	}

	// 确保 GVM 根目录存在
	if err := os.MkdirAll(gvmRoot, 0755); err != nil {
		fmt.Printf("Create gvm root directory failed: %v\n", err)
		os.Exit(1)
	}
	return gvmRoot
}

// SetDefaultVersion 设置默认的 Go 版本
func SetDefaultVersion(version string) {
	gvmRoot := GetGvmRoot()
	configFile := filepath.Join(gvmRoot, "config")

	// 写入默认版本到配置文件
	err := os.WriteFile(configFile, []byte(fmt.Sprintf("default=%s\n", version)), 0644)
	if err != nil {
		fmt.Printf("Write config file failed: %v\n", err)
	}
}

// GetDefaultVersion 获取默认的 Go 版本
func GetDefaultVersion() string {
	gvmRoot := GetGvmRoot()
	configFile := filepath.Join(gvmRoot, "config")

	// 读取配置文件
	content, err := os.ReadFile(configFile)
	if err != nil {
		return ""
	}

	// 解析配置文件，查找默认版本
	// 简单实现，实际可能需要更复杂的解析
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "default=") {
			return strings.TrimPrefix(line, "default=")
		}
	}

	return ""
}
