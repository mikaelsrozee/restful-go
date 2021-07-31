package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file", err)
    }

    // Init app
    app := App{}
    app.Init(os.Getenv("DB_NAME"))
}
