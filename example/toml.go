package main

import (
	"fmt"
	"github.com/crazy-canux/go-devops/configuration"
)

var fileName = "/home/canux/Src/go/src/github.com/crazy-canux/go-devops/data/configuration/test.toml"

type Config struct {
	title string
	constraint Constraint `toml:"constraint"`
}

type Constraint struct {
	Name string
	Version string
}

func main() {
	var config Config
	//config := &Config{}

	err := configuration.ParseFromFile(fileName, &config)
	if err != nil {
		fmt.Errorf("parse from file failed: %v", err)
	}
	fmt.Printf("result: %v", config.title)
}
