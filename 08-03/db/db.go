package db

import (
	"0803/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Con *gorm.DB

func Connect() {
	dsn := "host=db.blrhjjsznbvxrcrbctsw.supabase.co user=postgres password=SZGNBCaFfZb9exph dbname=project-naya port=5432 sslmode=disable TimeZone=Asia/Jakarta"
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
