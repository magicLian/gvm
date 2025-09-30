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
	if len(dirs) == 0 {
		fmt.Println("  No Go versions installed yet.")
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
	sort.Sort(sort.Reverse(sort.StringSlice(versions)))

	// Show current version.
	currentVersion := GetCurrentVersion()

	fmt.Println("Installed Go versions:")

	// Output version list.
	for _, version := range versions {
		if version == currentVersion {
			fmt.Printf("  => %s (current) ✅ \n", version)
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

	// current -> /opt/go/versions/1.25.1/go
	// filepath.Dir(target) -> /opt/go/versions/1.25.1
	// filepath.Base(filepath.Dir(target)) -> "1.25.1"
	return filepath.Base(filepath.Dir(target))
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
