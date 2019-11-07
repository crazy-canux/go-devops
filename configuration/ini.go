package configuration

import (
	"github.com/go-ini/ini"
	"log"
)

func LoadIni(filename string) (*ini.File, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		log.Printf("Load %s failed.", filename)
		return &ini.File{}, err
	}
	return cfg, nil
}

func GetKey(cfg *ini.File, section, key string) (*ini.Key, error) {
	sec, err := cfg.GetSection(section)
	if err != nil {
		log.Printf("Section %s not exist.", section)
		return &ini.Key{}, err
	}
	k, err := sec.GetKey(key)
	if err != nil {
		log.Printf("Key %s not exist.", key)
		return &ini.Key{}, err
	}
	return k, nil
}

func SetKey(cfg *ini.File, section, key, value string, create bool) error {
	sec, err := cfg.GetSection(section)
	if err != nil {
		log.Printf("Section %s not exist.", section)
		return nil
	}
	k, err := sec.GetKey(key)
	if err != nil && !create {
		log.Printf("Key %s not exist.", key)
		return err
	} else if err != nil && create {
		log.Println("Add new key.")
		_, err = sec.NewKey(key, value)
		if err != nil {
			log.Printf("Create new key %s failed.", key)
			return nil
		}
	} else {
		k.SetValue(value)
	}
	return nil
}
