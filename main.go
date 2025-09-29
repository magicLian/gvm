package main

import (
	"fmt"
	"gvm/pkg/commands"
	"os"
)

func main() {
	// 检查命令行参数
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}

	// 根据子命令执行不同的操作
	command := os.Args[1]
	switch command {
	case "install":
		if len(os.Args) < 3 {
			fmt.Println("请指定要安装的 Go 版本，例如: gvm install 1.20.0")
			os.Exit(1)
		}
		version := os.Args[2]
		commands.Install(version)
	case "use":
		if len(os.Args) < 3 {
			fmt.Println("请指定要使用的 Go 版本，例如: gvm use 1.20.0")
			os.Exit(1)
		}
		version := os.Args[2]
		commands.Use(version)
	case "ls":
		commands.ListInstalled()
	case "uninstall":
		if len(os.Args) < 3 {
			fmt.Println("请指定要卸载的 Go 版本，例如: gvm uninstall 1.20.0")
			os.Exit(1)
		}
		version := os.Args[2]
		commands.Uninstall(version)
	case "--help":
		help()
	default:
		fmt.Printf("未知命令: %s\n", command)
		help()
		os.Exit(1)
	}
}

// 显示帮助信息
func help() {
	fmt.Println("gvm - Go 版本管理器")
	fmt.Println("\n用法:")
	fmt.Println("  gvm install <version>    安装指定版本的 Go")
	fmt.Println("  gvm use <version>        切换到指定版本的 Go")
	fmt.Println("  gvm list                 列出已安装的 Go 版本")
	fmt.Println("  gvm list-all             列出可安装的 Go 版本")
	fmt.Println("  gvm uninstall <version>  卸载指定版本的 Go")
	fmt.Println("  gvm --help               显示帮助信息")
}
