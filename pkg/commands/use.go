package commands

import (
	"fmt"
	"gvm/pkg/config"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Use Switch to the specified Go version.
func Use(version string) {
	fmt.Printf("Switching to Go %s...\n", version)

	// Get GVM root directory.
	gvmRoot := config.GetGvmRoot()
	goVersionDir := filepath.Join(gvmRoot, "versions", version)

	// Check if version is installed.
	if _, err := os.Stat(goVersionDir); os.IsNotExist(err) {
		fmt.Printf("Go %s is not installed. Please install it first using 'gvm install %s'\n", version, version)
		return
	}

	// Create or update current version link.
	currentDir := filepath.Join(gvmRoot, "current")

	// Delete old link (if exists).
	if _, err := os.Lstat(currentDir); err == nil {
		if err = os.Remove(currentDir); err != nil {
			fmt.Printf("Failed to remove old link: %v\n", err)
			return
		}
	}

	// Create new link.
	goVersionBinDir := filepath.Join(goVersionDir, "go", "bin")
	if err := createSymlink(goVersionBinDir, currentDir); err != nil {
		fmt.Printf("Failed to create new link: %v\n", err)
		return
	}

	fmt.Printf("Go %s is now set as current version.\n", version)
}

// createSymlink Create link to the specified target.
func createSymlink(target, link string) error {
	osType := runtime.GOOS

	if osType == "windows" {
		// Windows use mklink command.
		cmd := exec.Command("cmd", "/c", "mklink", "/D", link, target)
		return cmd.Run()
	} else {
		// Unix-like systems use os.Symlink.
		return os.Symlink(target, link)
	}
}
