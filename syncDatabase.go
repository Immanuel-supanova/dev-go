package devgo

import (
	"github.com/immanuel-supanova/dev-go/database"
	"github.com/immanuel-supanova/dev-go/models"
)

func SyncDatabase() {
	database.DB.AutoMigrate(&models.Developer{})
	database.DB.AutoMigrate(&models.Application{})
	database.DB.AutoMigrate(&models.Log{})

}
