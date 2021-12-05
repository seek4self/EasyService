package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var dirs = []string{
	"/api/v1",
	"/global",
	"/initialize",
	"/model/config",
	"/middleware",
	"/service",
}

func InitServiceDir(name string) {
	for _, dir := range dirs {
		Mkdir(name + dir)
	}
}

func Mkdir(name string) {
	err := os.MkdirAll(name, 0666)
	if err != nil {
		fmt.Println("mkdir err", err)
		os.Exit(1)
	}
}

func ReadDir() (dirs, files []string) {
	// dirs = make([]string, 0)
	// files = make([]string, 0)
	err := filepath.Walk("templates", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("walk dir err", err)
		os.Exit(1)
	}
	return
}
