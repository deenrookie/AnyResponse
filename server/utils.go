package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
