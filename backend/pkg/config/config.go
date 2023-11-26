package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type Config struct {
	Port   int    `yaml:"port"`
	Debug  bool   `yaml:"debug"`
	LogDir string `yaml:"log_dir"`

	StaticDir string `yaml:"static_dir"`
	Dsn       string `yaml:"dsn"`
}

func Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	return yaml.NewDecoder(f).Decode(&Cfg)
}

func Port() int {
	return Cfg.Port
}

func Debug() bool {
	return Cfg.Debug
}

func LogDir() string {
	return Cfg.LogDir
}

func StaticDir() string {
	return Cfg.StaticDir
}

func Dsn() string {
	return Cfg.Dsn
}
