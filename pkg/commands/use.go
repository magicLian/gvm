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
	currentLink := filepath.Join(gvmRoot, "current")

	// Delete old link (if exists).
	if _, err := os.Lstat(currentLink); err == nil {
		err = os.Remove(currentLink)
		if err != nil {
			fmt.Printf("Failed to remove old link: %v\n", err)
			return
		}
	}

	// Create new link.
	err := createSymlink(goVersionDir, currentLink)
	if err != nil {
		fmt.Printf("Failed to create new link: %v\n", err)
		return
	}

	// Update default version in config file.
	if len(os.Args) > 3 && os.Args[3] == "--default" {
		config.SetDefaultVersion(version)
		fmt.Printf("Go %s is now set as default version.\n", version)
	}

	fmt.Printf("Go %s is now set as current version.\n", version)
	fmt.Println("Please restart your terminal or run 'source ~/.bashrc' (depending on your shell type) to apply the changes.")
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
