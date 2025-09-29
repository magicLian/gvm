package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MoveDirectory Move source directory to destination.
func MoveDirectory(source, destination string) error {
	// Ensure parent directory of destination exists.
	parentDir := filepath.Dir(destination)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return fmt.Errorf("Create parent directory failed: %v", err)
	}

	// Use os.Rename to move directory.
	if err := os.Rename(source, destination); err != nil {
		// If Rename fails (e.g., cross-file system), manually copy files.
		return CopyDirectory(source, destination)
	}

	return nil
}

// CopyDirectory Copy source directory to destination.
func CopyDirectory(source, destination string) error {
	// Ensure destination directory exists.
	if err := os.MkdirAll(destination, 0755); err != nil {
		return fmt.Errorf("Create destination directory failed: %v", err)
	}

	// Read source directory.
	dirEntries, err := os.ReadDir(source)
	if err != nil {
		return fmt.Errorf("Read source directory failed: %v", err)
	}

	for _, entry := range dirEntries {
		srcPath := filepath.Join(source, entry.Name())
		destPath := filepath.Join(destination, entry.Name())

		// Check if entry is a directory.
		if entry.IsDir() {
			// Recursively copy subdirectory.
			if err := CopyDirectory(srcPath, destPath); err != nil {
				return err
			}
		} else {
			// Copy file.
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return fmt.Errorf("Open source file failed: %v", err)
			}
			defer srcFile.Close()

			destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, entry.Type())
			if err != nil {
				return fmt.Errorf("Create destination file failed: %v", err)
			}
			defer destFile.Close()

			// Copy file content.
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return fmt.Errorf("Copy file content failed: %v", err)
			}
		}
	}

	return nil
}
