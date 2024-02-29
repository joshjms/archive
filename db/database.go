package db

import (
	"fmt"
	"log"
	"os"

	fs "github.com/joshjms/archive/filesystem/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Initializes the database connection
func Init() (*gorm.DB, error) {
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&fs.File{}); err != nil {
		log.Print("Error migrating file")
		return nil, err
	}

	if err := db.AutoMigrate(&fs.Directory{}); err != nil {
		log.Print("Error migrating directory")
		return nil, err
	}

	return db, nil
}
