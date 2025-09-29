package commands

import (
	"fmt"
	"gvm/pkg/config"
	"os"
	"path/filepath"
)

// Uninstall Uninstall specified Go version.
func Uninstall(version string) {
	fmt.Printf("正在卸载 Go %s...\n", version)
	
	// Get GVM root directory.
	gvmRoot := config.GetGvmRoot()
	goVersionDir := filepath.Join(gvmRoot, "versions", version)
	
	// Check if version is installed.
	if _, err := os.Stat(goVersionDir); os.IsNotExist(err) {
		fmt.Printf("Go %s is not installed.\n", version)
		return
	}
	
	// Check if version is currently used.
	currentVersion := GetCurrentVersion()
	if currentVersion == version {
		fmt.Printf("Cannot uninstall the currently used version: %s\n", version)
		fmt.Println("Please switch to another version first using 'gvm use'.")
		return
	}
	
	// Remove version directory.
	err := os.RemoveAll(goVersionDir)
	if err != nil {
		fmt.Printf("Uninstallation failed: %v\n", err)
		return
	}
	
	fmt.Printf("Go %s uninstalled successfully!\n", version)
}