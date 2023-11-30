package config

import (
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type Config struct {
	Default  Default  `yaml:"default"`
	Database Database `yaml:"database"`
	SMTP     SMTP     `yaml:"smtp"`
}

func Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = yaml.NewDecoder(f).Decode(&Cfg)
	if err != nil {
		return err
	}

	return Cfg.Default.verify()
}

func generateRandomString() string {
	strs := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := make([]byte, 16)
	for i := 0; i < 16; i++ {
		ret[i] = strs[rand.Intn(len(strs))]
	}
	return string(ret)
}
