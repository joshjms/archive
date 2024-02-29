package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/joshjms/archive/api"
	"github.com/joshjms/archive/db"
	fs "github.com/joshjms/archive/filesystem"
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

	fmt.Println("Connected to database")

	fs := fs.New(db)
	fs.Init()

	api.Init(fs)
}
