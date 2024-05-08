package devgo

import (
	"os"

	"github.com/immanuel-supanova/dev-go/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() {
	var err error

	dsn := os.Getenv("DB")
	database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

}
