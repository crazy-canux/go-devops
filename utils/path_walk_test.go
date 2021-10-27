package utils

import (
	"fmt"
	"testing"
)

var dirExclude = []string{".git"}

var fileExclude = []string{
	".gitignore",
	"README.txt",
	"auto_update.sh",
	"update_index.py",
	"update_sigs.py",
	"update_sigs.log",
}

func TestPathWalk(t *testing.T) {
	files, err := PathWalk("/home/canux/Documents/hastebin", dirExclude, fileExclude)
	if err != nil {
		t.Error("list dir failed.")
	} else {
		for _, f := range files {
			fmt.Println(f)
		}
	}
}
