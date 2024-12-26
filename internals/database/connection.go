package database

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {

	once.Do(func(){
		var err error
		dsn := "host=" + os.Getenv("") + " user=" + os.Getenv("") + " password=" + os.Getenv("") + " dbname=" + os.Getenv("") + " port=" + os.Getenv("") + " sslmode=" + os.Getenv("") + " TimeZone=" + os.Getenv("")
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Database conection failed", err)
		}
	})

	return db
}

