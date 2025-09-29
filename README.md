# GVM - Go 版本管理器

GVM 是一个用于管理多个 Go 版本的工具。 它支持 Windows、Linux 和 macOS 系统。

## 功能特点

- 安装多个 Go 版本
- 切换当前使用的 Go 版本
- 列出已安装的 Go 版本
- 卸载不需要的 Go 版本

## 安装方法
下载对应系统的gvm可执行文件，并放置于环境变量中。

## 使用方法

```bash
# 安装指定版本的 Go
gvm install 1.24.0

# 切换到指定版本的 Go
gvm use 1.24.0

# 列出已安装的 Go 版本
gvm list

# 列出可安装的 Go 版本
gvm list-all

# 卸载指定版本的 Go
gvm uninstall 1.24.0

# 显示帮助信息
gvm help
```

## 编译项目

```bash
go build -o gvm main.go
```

## 核心目录功能

- `$GVM_ROOT`：GVM 工具的使用目录。
- `$GVM_ROOT/versions`：Go 版本的安装目录，每个版本的 Go 都在一个独立的子目录中。
- `$GVM_ROOT/archives`: Go 版本的压缩包文件。
- `$GVM_ROOT/current`: Go 当前可执行文件目录。

current目录为软链接，指向versions中当前使用的Go版本的可执行文件目录。