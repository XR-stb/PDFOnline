package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"

	"backend/pkg/api"
	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/logger"
	"backend/pkg/provider"
	"backend/pkg/static"
	"backend/pkg/user"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "./etc/config.yaml", "config file path")
}

func main() {
	flag.Parse()

	if err := config.Parse(configPath); err != nil {
		logrus.Fatalf("parse config file error: %v", err)
	}

	if err := logger.Init(); err != nil {
		logrus.Fatalf("init logger error: %v", err)
	}

	if err := database.Init(); err != nil {
		logrus.Fatalf("init database error: %v", err)
	}

	if err := user.CreateInternalUser(); err != nil {
		logrus.Fatalf("create internal user error: %v", err)
	}

	if err := static.Init(); err != nil {
		logrus.Fatalf("init static directory error: %v", err)
	}

	if err := provider.Init(); err != nil {
		logrus.Fatalf("init provider error: %v", err)
	}

	if err := api.NewRoute().Run(fmt.Sprintf(":%d", config.Port())); err != nil {
		logrus.Fatalf("run api server error: %v", err)
	}
}
