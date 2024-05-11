package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/dev-go/database"
	"github.com/immanuel-supanova/dev-go/jwtauth"
	"github.com/immanuel-supanova/dev-go/models"
	"gorm.io/gorm"
)

func RequireAuth(c *gin.Context) {
	// Get the token from the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	// Decode/validate it
	appid, err := jwtauth.DecodeAccessToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
		return
	}

	// check if app exists
	var app = models.Application{ID: appid}
	appResult := database.DB.First(&app)

	if appResult.Error != nil {
		if errors.Is(appResult.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "App has not been found"})

			return

		} else {
			// Handle other database errors
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})

		}
		return
	}

	// Check if application is active
	if !app.IsActive {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev has been suspended"})
		return
	}

	// Get the developer associated with the app
	devid := app.DevID
	var dev = models.Developer{UUID: devid}

	// Check if developer exists
	result := database.DB.First(&dev)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev has not been found"})

			return

		} else {
			// Handle other database errors
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})

		}
		return
	}

	// Check if developer is active
	if !dev.IsActive {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev has been suspended"})
		return
	}

	// Continue
	c.Next()

}

func RequireAdmin(c *gin.Context) {
	// Get the token from the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	// Decode/validate it
	appid, err := jwtauth.DecodeAccessToken(tokenString)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
		return
	}

	// check if app exists
	var app = models.Application{ID: appid}
	appResult := database.DB.First(&app)

	if appResult.Error != nil {
		if errors.Is(appResult.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "App has not been found"})

			return

		} else {
			// Handle other database errors
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})

		}
		return
	}

	// Check if application is active
	if !app.IsActive {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev has been suspended"})
		return
	}

	// Get the developer associated with the app
	devid := app.DevID
	var dev = models.Developer{UUID: devid}

	// Check if developer exists
	result := database.DB.First(&dev)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev has not been found"})
			return
		} else {
			// Handle other database errors
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
		}
		return
	}

	// Check if developer is active
	if !dev.IsActive {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev has been suspended"})
	}

	// Check if developer is admin
	if !dev.IsAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Dev does not have admin priveleges"})
	}

	// Continue
	c.Next()

}
