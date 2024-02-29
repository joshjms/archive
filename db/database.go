package db

import (
	"fmt"
	"os"

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

	return db, nil
}
