package commands

import (
	"fmt"
	"gvm/pkg/config"
	"os"
	"path/filepath"
)

// Uninstall 卸载指定版本的 Go
func Uninstall(version string) {
	fmt.Printf("正在卸载 Go %s...\n", version)
	
	// 获取 GVM 根目录
	gvmRoot := config.GetGvmRoot()
	goVersionDir := filepath.Join(gvmRoot, "versions", version)
	
	// 检查版本是否已安装
	if _, err := os.Stat(goVersionDir); os.IsNotExist(err) {
		fmt.Printf("Go %s 尚未安装\n", version)
		return
	}
	
	// 检查是否正在使用该版本
	currentVersion := GetCurrentVersion()
	if currentVersion == version {
		fmt.Printf("无法卸载正在使用的版本: %s\n", version)
		fmt.Println("请先使用 'gvm use' 切换到其他版本")
		return
	}
	
	// 删除版本目录
	err := os.RemoveAll(goVersionDir)
	if err != nil {
		fmt.Printf("卸载失败: %v\n", err)
		return
	}
	
	fmt.Printf("Go %s 卸载成功!\n", version)
}