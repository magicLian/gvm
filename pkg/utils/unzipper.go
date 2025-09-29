package utils

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Unzip 解压文件，支持 zip 和 tar.gz 格式
func Unzip(source, destination string) error {
	// 检查文件扩展名
	if strings.HasSuffix(source, ".zip") {
		return unzipFile(source, destination)
	} else if strings.HasSuffix(source, ".tar.gz") || strings.HasSuffix(source, ".tgz") {
		return untarFile(source, destination)
	} else {
		return fmt.Errorf("不支持的文件格式: %s", source)
	}
}

// unzipFile 解压 zip 文件
func unzipFile(source, destination string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		path := filepath.Join(destination, file.Name)

		// 创建目录
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		// 确保父目录存在
		os.MkdirAll(filepath.Dir(path), 0755)

		// 解压文件
		destFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		sourceFile, err := file.Open()
		if err != nil {
			destFile.Close()
			return err
		}

		_, err = io.Copy(destFile, sourceFile)
		sourceFile.Close()
		destFile.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

// untarFile 解压 tar.gz 文件
func untarFile(source, destination string) error {
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(destination, hdr.Name)

		// 检查文件类型
		switch hdr.Typeflag {
		case tar.TypeDir:
			// 创建目录
			os.MkdirAll(path, 0755)
		case tar.TypeReg:
			// 确保父目录存在
			os.MkdirAll(filepath.Dir(path), 0755)
			
			// 解压文件
			destFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, hdr.FileInfo().Mode())
			if err != nil {
				return err
			}
			defer destFile.Close()
			
			_, err = io.Copy(destFile, tarReader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}