package initialize

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBase() {
	var err error
	dsn := "host=localhost user=postgres password=2212 dbname=gin_jwt port=5432"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error en la base de datos")
	}
}
