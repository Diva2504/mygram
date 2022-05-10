package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/takadev15/mygram-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userName = os.Getenv("DB_USER")
	dbName   = os.Getenv("DB_NAME")
	dbPass   = os.Getenv("DB_PASSWORD")
	dbPort   = os.Getenv("DB_PORT")
  dbHost   = os.Getenv("DB_HOST")
	db       *gorm.DB
	err      error
)

func DBInit() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, userName, dbPass, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Databases Error", err.Error())
	}
	log.Printf("Databases Connected")
	db.Debug().AutoMigrate(models.Photo{})
}

func GetDB() *gorm.DB {
  return db
}
