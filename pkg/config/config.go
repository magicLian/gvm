package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetGvmRoot Get GVM root directory.
func GetGvmRoot() string {
	var gvmRoot string

	// Find GVM root directory from environment variable.
	gvmRoot = os.Getenv("GVM_ROOT")
	if gvmRoot == "" {
		// If environment variable not found, determine GVM root directory based on operating system.
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Get user home directory failed: %v\n", err)
			os.Exit(1)
		}

		switch runtime.GOOS {
		case "windows":
			// Windows use APPDATA environment variable.
			appData := os.Getenv("APPDATA")
			if appData != "" {
				gvmRoot = filepath.Join(appData, "gvm")
			} else {
				// If APPDATA not found, use user home directory.
				gvmRoot = filepath.Join(homeDir, "gvm")
			}
		default:
			// Unix-like systems use .gvm directory.
			gvmRoot = filepath.Join(homeDir, ".gvm")
		}
	}

	// Ensure GVM root directory exists.
	if err := os.MkdirAll(gvmRoot, 0755); err != nil {
		fmt.Printf("Create gvm root directory failed: %v\n", err)
		os.Exit(1)
	}
	return gvmRoot
}

// SetDefaultVersion Set default Go version.
func SetDefaultVersion(version string) {
	gvmRoot := GetGvmRoot()
	configFile := filepath.Join(gvmRoot, "config")

	// Write default version to config file.
	err := os.WriteFile(configFile, []byte(fmt.Sprintf("default=%s\n", version)), 0644)
	if err != nil {
		fmt.Printf("Write config file failed: %v\n", err)
	}
}

// GetDefaultVersion Get default Go version.
func GetDefaultVersion() string {
	gvmRoot := GetGvmRoot()
	configFile := filepath.Join(gvmRoot, "config")

	// Read config file.
	content, err := os.ReadFile(configFile)
	if err != nil {
		return ""
	}

	// Parse config file, find default version.
	// Simple implementation, actual may need more complex parsing.
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "default=") {
			// Trim prefix "default=" and return default version.
			return strings.TrimPrefix(line, "default=")
		}
	}

	return ""
}
