package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// MoveDirectory 移动目录
func MoveDirectory(source, destination string) error {
	// 确保目标目录的父目录存在
	parentDir := filepath.Dir(destination)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		return fmt.Errorf("创建父目录失败: %v", err)
	}

	// 使用 os.Rename 移动目录
	if err := os.Rename(source, destination); err != nil {
		// 如果 Rename 失败（比如跨文件系统），则手动复制文件
		return CopyDirectory(source, destination)
	}

	return nil
}

// CopyDirectory 复制目录（当 os.Rename 失败时使用）
func CopyDirectory(source, destination string) error {
	// 确保目标目录存在
	if err := os.MkdirAll(destination, 0755); err != nil {
		return err
	}

	// 读取源目录
	dirEntries, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	for _, entry := range dirEntries {
		srcPath := filepath.Join(source, entry.Name())
		destPath := filepath.Join(destination, entry.Name())

		// 检查是否为目录
		if entry.IsDir() {
			// 递归复制子目录
			if err := CopyDirectory(srcPath, destPath); err != nil {
				return err
			}
		} else {
			// 复制文件
			srcFile, err := os.Open(srcPath)
			if err != nil {
				return err
			}
			defer srcFile.Close()

			destFile, err := os.OpenFile(destPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, entry.Type())
			if err != nil {
				return err
			}
			defer destFile.Close()

			// 复制文件内容
			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
