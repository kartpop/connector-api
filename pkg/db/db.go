package db

import (
	"github.com/sede-x/gopoc-connector/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Connector{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
