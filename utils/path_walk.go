package utils

import (
	"log"
	"os"
	"path/filepath"
)

var files []string

// Return all files with absolute path in `dir`, but exclude files in `dirEx`, and exclude file in `fileEx`.
func PathWalk(dir string, dirEx, fileEx []string) ([]string, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Processing %s failed.", info.Name())
			return nil
		}

		if info.IsDir() {
			if In(info.Name(), dirEx) {
				return filepath.SkipDir
			} else {
				return nil
			}
		}

		if In(info.Name(), fileEx) {
			return nil
		} else {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		log.Println("traverse dir failed.")
		return files, err
	}
	return files, nil
}

// If `name` in `slice`, return true, else return false.
func In(name string, slice []string) bool {
	for _, s := range slice {
		if name == s {
			return true
		}
	}
	return false
}

// If extention of `file` in list, return true, else return false.
func ExtIn(file string, list []string) bool {
	for _, ext := range list {
		if filepath.Ext(file) == ext {
			return true
		}
	}
	return false
}
