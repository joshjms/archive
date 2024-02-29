package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/joshjms/archive/db"
	"github.com/joshjms/archive/filesystem"
)

func main() {
	log.Print("Starting application")

	// load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// connect to database
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	fs := filesystem.New(db)
	fs.Init()

	fmt.Println("Connected to database")
}
