package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// DownloadFile 从指定 URL 下载文件
func DownloadFile(url, filepath string) error {
	// 创建文件
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 发送 HTTP 请求
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败: 状态码 %d", resp.StatusCode)
	}

	// 将响应内容写入文件
	_, err = io.Copy(out, resp.Body)
	return err
}
