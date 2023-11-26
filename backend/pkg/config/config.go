package config

import (
	"errors"
	"math/rand"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

type Config struct {
	Port   int    `yaml:"port"`
	Debug  bool   `yaml:"debug"`
	LogDir string `yaml:"log_dir"`

	StaticDir     string `yaml:"static_dir"`
	Dsn           string `yaml:"dsn"`
	AdminUser     string `yaml:"admin_user"`
	AdminPassword string `yaml:"admin_password"`
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

	return verify(Cfg)
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

func AdminUser() string {
	return Cfg.AdminUser
}

func AdminPassword() string {
	return Cfg.AdminPassword
}

func verify(c *Config) error {
	if c.Port == 0 {
		c.Port = 8080
	}

	if c.StaticDir == "" {
		c.StaticDir = "./static"
	}

	if c.Dsn == "" {
		return errors.New("config dsn is required")
	}

	if c.AdminUser == "" {
		c.AdminUser = "admin"
	}

	if c.AdminPassword == "" {
		c.AdminPassword = generateRandomString()
	}

	if len(c.AdminUser) < 5 {
		return errors.New("config admin user is too short, min 5 chars")
	}

	if len(c.AdminUser) > 36 {
		return errors.New("config admin user is too long, max 36 chars")
	}

	if len(c.AdminPassword) < 6 {
		return errors.New("config admin password is too short, min 6 chars")
	}

	if len(c.AdminPassword) > 36 {
		return errors.New("config admin password is too long, max 36 chars")
	}

	return nil
}

func generateRandomString() string {
	strs := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := make([]byte, 16)
	for i := 0; i < 16; i++ {
		ret[i] = strs[rand.Intn(len(strs))]
	}
	return string(ret)
}
