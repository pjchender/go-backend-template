package main

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/pjchender/go-backend-template/global"
	"github.com/pjchender/go-backend-template/internal/database"
	"github.com/pjchender/go-backend-template/pkg/setup"
)

func init() {
	var err error

	// setupEnv should invoke before setupSetting()
	err = setup.Env()
	if err != nil {
		log.Fatalf("init.setupEnv failed: %v", err)
	}

	err = setup.Settings()
	if err != nil {
		log.Fatalf("init.setupSetting failed: %v", err)
	}
}

func main() {
	db, err := database.New(global.DatabaseSetting.DSN, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	db.DropAllTables()
}
