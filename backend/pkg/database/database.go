package database

import (
	"backend/pkg/config"
	"backend/pkg/database/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var db *gorm.DB

func Init() error {
	var err error
	db, err = gorm.Open(mysql.Open(config.Dsn()), &gorm.Config{
		TranslateError: true,
		Logger: logger.New(logrus.StandardLogger(), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: false,
			Colorful:                  false,
		}),
	})
	if err != nil {
		return fmt.Errorf("open database failed: %v", err)
	}

	if err = AutoMigrate(db); err != nil {
		return err
	}

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(models.PDF{}, models.User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}

func Instance() *gorm.DB {
	if db == nil {
		if err := Init(); err != nil {
			logrus.Fatalf("init database error: %v", err)
		}
	}

	return db
}

// Use ONLY FOR TEST
func Use(mockDB *gorm.DB) {
	db = mockDB
}
