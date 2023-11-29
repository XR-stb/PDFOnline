package config

import "errors"

type Default struct {
	Port          int    `yaml:"port"`
	Debug         bool   `yaml:"debug"`
	LogDir        string `yaml:"log_dir"`
	StaticDir     string `yaml:"static_dir"`
	AdminUsername string `yaml:"admin_user"`
	AdminPassword string `yaml:"admin_password"`
}

func Port() int {
	return Cfg.Default.Port
}

func Debug() bool {
	return Cfg.Default.Debug
}

func LogDir() string {
	return Cfg.Default.LogDir
}

func StaticDir() string {
	return Cfg.Default.StaticDir
}

func AdminUsername() string {
	return Cfg.Default.AdminUsername
}

func AdminPassword() string {
	return Cfg.Default.AdminPassword
}

func (c *Default) verify() error {
	if c.Port == 0 {
		c.Port = 8080
	}

	if c.StaticDir == "" {
		c.StaticDir = "./static"
	}

	if c.AdminUsername == "" {
		c.AdminUsername = "admin"
	}

	if c.AdminPassword == "" {
		c.AdminPassword = generateRandomString()
	}

	if len(c.AdminUsername) < 5 {
		return errors.New("config admin user is too short, min 5 chars")
	}

	if len(c.AdminUsername) > 36 {
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
