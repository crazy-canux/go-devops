package tests

import (
	"path/filepath"
	"os"
	"flag"
	"crypto/md5"
	"log"
	"fmt"
	"io"
)

func main() {
	dir := flag.String("path", ".", "Specify the directory")
	flag.Parse()

	if ! filepath.IsAbs(*dir) {
		log.Fatal("Path must be a absolute diractory")
	}

    var fileList []string

    walkErr := filepath.Walk(
		*dir,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if ! f.IsDir() {
				fileList = append(fileList, path)
			}
			return nil
		})
    if walkErr != nil {
    	log.Fatalf("Walk path failed: %s", walkErr)
	}

	for _, filename := range fileList {
		f, fErr := os.Open(filename)
		if fErr != nil {
			log.Fatal("Open file error: ", fErr)
			return
		}
		defer f.Close()

		md5Ctx := md5.New()
		if _, err := io.Copy(md5Ctx, f); err != nil {
			log.Fatal(err)
		}
		cipher := md5Ctx.Sum(nil)

		fmt.Print(filename + ": ")
		fmt.Printf("%x\n", cipher)
	}
}

    
