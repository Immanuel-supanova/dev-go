package devgo

import (
	"log"

	"github.com/joho/godotenv"
)

func Config() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environments")
	}
}
