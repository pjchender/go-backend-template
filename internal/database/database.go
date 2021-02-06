package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/pjchender/go-backend-template/internal/model"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(dsn string, gormConfig *gorm.Config) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}

func (d *GormDatabase) AutoMigrate() {
	// enable format UUID as PK
	d.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := d.DB.AutoMigrate(
		&model.Organization{},
		&model.User{},
	); err != nil {
		log.Fatal(err.Error())
	}
}

func (d *GormDatabase) DropAllTables() {
	if err := d.DB.Migrator().DropTable(
		"organizations",
		"users",
	); err != nil {
		log.Fatal(err.Error())
	}
}
