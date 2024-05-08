package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/immanuel-supanova/dev-go/database"
	"github.com/immanuel-supanova/dev-go/models"
)

func ApplicationCreate(c *gin.Context) {
	var data struct {
		Name  string
		DevID string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name not Provided",
		})
		return
	}

	if data.DevID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "DevID not Provided",
		})
		return
	}

	devuuid, err := uuid.Parse(data.DevID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create Application
	uuid, _ := uuid.NewV7()
	app := models.Application{Name: data.Name, IsActive: true, UUID: uuid, DevID: devuuid}
	result := database.DB.Create(&app)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Application",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Application Created",
	})
}

func ApplicationRead(c *gin.Context) {
	var data struct {
		UUID string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UUID not Provided",
		})
		return
	}

	appuuid, err := uuid.Parse(data.UUID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Get developer data from database
	var app = models.Application{UUID: appuuid}
	database.DB.Find(&app)

	// Check if User exists
	if app.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"app": app,
	})

}

func ApplicationList(c *gin.Context) {
	// Get all applications from database
	var apps []models.Application
	database.DB.Find(&apps)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"apps": apps,
	})
}

func ApplicationUpdate(c *gin.Context) {
	var data struct {
		UUID string
		Name string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UUID not Provided",
		})
		return
	}

	appuuid, err := uuid.Parse(data.UUID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if App exists
	var app = models.Application{UUID: appuuid}
	database.DB.First(&app)

	// Update Application
	result := database.DB.Model(&app).Updates(models.Application{
		Name: data.Name,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update application",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Application Updated",
	})

}

func ApplicationIsActive(c *gin.Context) {
	var data struct {
		UUID     string
		IsActive bool
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UUID not Provided",
		})
		return
	}

	appuuid, err := uuid.Parse(data.UUID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if App exists
	var app = models.Application{UUID: appuuid}
	database.DB.First(&app)

	// Update Application
	result := database.DB.Model(&app).Updates(models.Application{
		IsActive: data.IsActive,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update application",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Application Updated",
	})

}

func ApplicationDelete(c *gin.Context) {
	var data struct {
		UUID string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.UUID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "UUID not Provided",
		})
		return
	}

	appuuid, err := uuid.Parse(data.UUID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if App exists
	var app = models.Application{UUID: appuuid}
	database.DB.First(&app)

	// Delete App
	result := database.DB.Delete(&app)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete app",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Application Deleted",
	})
}
