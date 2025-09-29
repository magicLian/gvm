package commands

import (
	"fmt"
	"gvm/pkg/config"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Use 切换到指定版本的 Go
func Use(version string) {
	fmt.Printf("正在切换到 Go %s...\n", version)

	// 获取 GVM 根目录
	gvmRoot := config.GetGvmRoot()
	goVersionDir := filepath.Join(gvmRoot, "versions", version)

	// 检查版本是否已安装
	if _, err := os.Stat(goVersionDir); os.IsNotExist(err) {
		fmt.Printf("Go %s 尚未安装，请先使用 'gvm install %s' 安装\n", version, version)
		return
	}

	// 创建或更新当前版本链接
	currentLink := filepath.Join(gvmRoot, "current")

	// 删除旧的链接（如果存在）
	if _, err := os.Lstat(currentLink); err == nil {
		err = os.Remove(currentLink)
		if err != nil {
			fmt.Printf("删除旧链接失败: %v\n", err)
			return
		}
	}

	// 创建新链接
	err := createSymlink(goVersionDir, currentLink)
	if err != nil {
		fmt.Printf("创建链接失败: %v\n", err)
		return
	}

	// 更新配置文件中的默认版本
	if len(os.Args) > 3 && os.Args[3] == "--default" {
		config.SetDefaultVersion(version)
		fmt.Printf("已将 Go %s 设置为默认版本\n", version)
	}

	fmt.Printf("成功切换到 Go %s\n", version)
	fmt.Println("请重新打开终端或执行 'source ~/.bashrc' (根据您的shell类型) 使更改生效")
}

// createSymlink 创建符号链接，适配不同操作系统
func createSymlink(target, link string) error {
	osType := runtime.GOOS

	if osType == "windows" {
		// Windows 上使用 mklink 命令
		cmd := exec.Command("cmd", "/c", "mklink", "/D", link, target)
		return cmd.Run()
	} else {
		// Unix-like 系统上使用 os.Symlink
		return os.Symlink(target, link)
	}
}
