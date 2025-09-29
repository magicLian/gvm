package commands

import (
	"fmt"
	"gvm/pkg/config"
	"os"
	"path/filepath"
	"sort"
)

// ListInstalled 列出所有已安装的 Go 版本
func ListInstalled() {
	fmt.Println("已安装的 Go 版本:")
	
	// 获取 GVM 根目录
	gvmRoot := config.GetGvmRoot()
	versionsDir := filepath.Join(gvmRoot, "versions")
	
	// 读取 versions 目录
	dirs, err := os.ReadDir(versionsDir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("  尚未安装任何 Go 版本")
		} else {
			fmt.Printf("读取版本目录失败: %v\n", err)
		}
		return
	}
	
	// 收集版本列表
	var versions []string
	for _, dir := range dirs {
		if dir.IsDir() {
			versions = append(versions, dir.Name())
		}
	}
	
	// 排序版本
	sort.Strings(versions)
	
	// 显示当前使用的版本
	currentVersion := GetCurrentVersion()
	
	// 输出版本列表
	for _, version := range versions {
		if version == currentVersion {
			fmt.Printf("  => %s (当前使用)\n", version)
		} else {
			fmt.Printf("    %s\n", version)
		}
	}
}

// GetCurrentVersion 获取当前使用的 Go 版本
func GetCurrentVersion() string {
	gvmRoot := config.GetGvmRoot()
	currentLink := filepath.Join(gvmRoot, "current")
	
	// 检查当前链接是否存在
	target, err := os.Readlink(currentLink)
	if err != nil {
		return ""
	}
	
	// 从路径中提取版本号
	return filepath.Base(target)
}

// ShowCurrent 显示当前使用的 Go 版本
func ShowCurrent() {
	currentVersion := GetCurrentVersion()
	if currentVersion != "" {
		fmt.Printf("当前使用的 Go 版本: %s\n", currentVersion)
	} else {
		fmt.Println("当前没有选择 Go 版本，请使用 'gvm use <version>' 切换版本")
	}
}