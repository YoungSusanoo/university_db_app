package main

import (
	"os"
	"university_app/internal/applic"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("no env file")
	}

	url := os.Getenv("POSTGRES_URL")
	dbName := os.Getenv("POSTGRES_SCHOOL_DB_NAME")
	app := applic.NewApp(url, dbName)
	app.Run()
}
