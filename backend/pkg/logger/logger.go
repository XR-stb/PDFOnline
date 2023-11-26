package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"backend/pkg/config"
)

func Init() error {
	setLevel()

	return setOutput()
}

func setLevel() {
	if config.Debug() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func setOutput() error {
	logDir := config.LogDir()

	if logDir == "" {
		return nil
	}

	_, err := os.Stat(logDir)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			err = os.MkdirAll(logDir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("create log dir error: %v", err)
			}
		default:
			return fmt.Errorf("get log dir stat error: %v", err)
		}
	}

	filePath := filepath.Join(logDir, "backend.log")
	logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("open log file %s error: %v", filePath, err)
	}

	logrus.SetOutput(logFile)

	return nil
}
