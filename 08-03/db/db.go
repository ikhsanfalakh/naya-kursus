package db

import (
	"0803/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Con *gorm.DB

func Connect() {
	dsn := ""
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Con = db

	err = db.AutoMigrate(models.Sales{})
	if err != nil {
		panic(err)
	}
}
