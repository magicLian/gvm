package commands

import (
	"fmt"
	"gvm/pkg/config"
	"os"
	"path/filepath"
	"sort"
)

// ListInstalled Lists all installed Go versions.
func ListInstalled() {
	fmt.Println("Installed Go versions:")

	// Get GVM root directory.
	gvmRoot := config.GetGvmRoot()
	versionsDir := filepath.Join(gvmRoot, "versions")

	// Read versions directory.
	dirs, err := os.ReadDir(versionsDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("  No Go versions installed yet.")
		} else {
			fmt.Printf("Failed to read versions directory: %v\n", err)
		}
		return
	}

	// Collect version list.
	var versions []string
	for _, dir := range dirs {
		if dir.IsDir() {
			versions = append(versions, dir.Name())
		}
	}

	// Sort versions.
	sort.Strings(versions)

	// Show current version.
	currentVersion := GetCurrentVersion()

	// Output version list.
	for _, version := range versions {
		if version == currentVersion {
			fmt.Printf("  => %s (current)\n", version)
		} else {
			fmt.Printf("    %s\n", version)
		}
	}
}

// GetCurrentVersion Get current version
func GetCurrentVersion() string {
	gvmRoot := config.GetGvmRoot()
	currentLink := filepath.Join(gvmRoot, "current")

	// Check if current link exists.
	target, err := os.Readlink(currentLink)
	if err != nil {
		return ""
	}

	// Extract version number from path.
	return filepath.Base(target)
}

// ShowCurrent Show current version.
func ShowCurrent() {
	currentVersion := GetCurrentVersion()
	if currentVersion != "" {
		fmt.Printf("Current Go version: %s\n", currentVersion)
	} else {
		fmt.Println("No Go version is currently selected. Please use 'gvm use <version>' to select a version.")
	}
}
