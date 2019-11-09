package configuration

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)


func Parse(data string, config interface{}) error {
	if _, err := toml.Decode(data, config); err != nil {
		log.Print("Parse toml data failed.")
		return err
	}
	return nil
}

func ParseFromFile(file string, config interface{}) error {
	if _, err := toml.DecodeFile(file, config); err != nil {
		fmt.Print("Parse toml file failed.")
		return err
    }
    return nil
}