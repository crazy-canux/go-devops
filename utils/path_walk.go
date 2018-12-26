package utils

import (
	"path/filepath"
	"log"
	"os"
)

var files []string

func PathWalk(dir string, dir_ex, file_ex []string) ([]string, error) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Processing %s failed.", info.Name())
			return nil
		}

		if info.IsDir() {
			for _, d := range dir_ex {
				if info.Name() == d {
					return filepath.SkipDir
				} else {
					return nil
				}
			}
		}

		for _, f := range file_ex {
			if info.Name() == f {
				return nil
			} else {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		log.Println("traverse dir failed.")
		return files, err
	}
	return files, nil
}