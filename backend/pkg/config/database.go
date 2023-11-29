package config

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db_name"`
}

func DatabaseConfig() *Database {
	return &Cfg.Database
}
