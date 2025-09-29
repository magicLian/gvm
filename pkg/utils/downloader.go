package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
)

// DownloadFile 下载文件并显示进度条
func DownloadFile(url, destPath string) error {
	// 确保目标目录存在
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("create directory failed: %v", err)
	}

	// 创建文件
	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}
	defer out.Close()

	// 创建带超时的 HTTP 客户端
	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("send HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: status code %d", resp.StatusCode)
	}

	// 创建进度条
	bar := progressbar.NewOptions64(
		resp.ContentLength,
		progressbar.OptionSetDescription("Downloading..."),
		progressbar.OptionSetWidth(30),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)

	// 写入文件 + 更新进度
	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	if err != nil {
		return fmt.Errorf("write response content failed: %v", err)
	}

	fmt.Println("\n✅ Download completed:", destPath)
	return nil
}
