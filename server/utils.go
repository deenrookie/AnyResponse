package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getText(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	content, err := ioutil.ReadAll(file)
	text := string(content)
	return text
}

func walkDir(dirPth string) []string {
	files := make([]string, 0)
	suffix := ".txt"
	_ = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		if strings.HasSuffix(fi.Name(), suffix) {
			files = append(files, filename)
		}
		return nil
	})

	return files
}

func md5v1(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func timeNow() string {
	return fmt.Sprintf("%v", time.Now().UnixNano())
}