package http

import (
	"io"
	"log"
	"os"
)

func Download(uri string, file string, username, password interface{}) bool {
	out, err := os.Create(file)
	if err != nil {
		log.Printf("Create destination file error: %s", err.Error())
		return false
	}
	defer out.Close()

	body, err := Get(uri, username, password)
	if err != nil {
		log.Printf("Get source file error: %s", err.Error())
		return false
	}
	_, err = io.Copy(out, body)
	if err != nil {
		log.Printf("Copy file error: %s", err.Error())
		return false
	}
	return true
}
