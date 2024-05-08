package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/immanuel-supanova/dev-go/database"
	"github.com/immanuel-supanova/dev-go/models"
	"golang.org/x/crypto/bcrypt"
)

func DeveloperCreate(c *gin.Context) {
	// Get Email and Password
	var data struct {
		Email     string
		Password  string
		Password2 string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	} else if data.Password2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Compare passwords
	if data.Password != data.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return

	}

	// Hash Password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Create Developer
	uuid, _ := uuid.NewV7()
	dev := models.Developer{Email: data.Email, Password: string(hash), IsActive: true, UUID: uuid}
	result := database.DB.Create(&dev)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Developer Created",
	})
}

func DeveloperCreateAdmin(c *gin.Context) {
	// Get user inputs
	var data struct {
		Email     string
		Password  string
		Password2 string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	} else if data.Password2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Compare passwords
	if data.Password != data.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
		return

	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// check if there is an admin
	result := database.DB.Where(&models.Developer{IsAdmin: true}).First(&models.Developer{})

	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "An admin already exists",
		})
		return
	}

	// create admin
	uuid, _ := uuid.NewV7()
	dev2 := models.Developer{Email: data.Email, Password: string(hash), IsActive: true, IsAdmin: true, UUID: uuid}
	result = database.DB.Create(&dev2)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create admin:",
		})
		return

	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Admin Created",
	})
}

// Admin Priveleges
func DeveloperRead(c *gin.Context) {
	// Get Email
	var data struct {
		Email string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	}

	// Get developer data from database
	var dev = models.Developer{Email: data.Email}
	database.DB.Omit("Password").Omit("UUID").Omit("IsAdmin").Find(&dev)

	// Check if User exists
	if dev.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"dev": dev,
	})
}

func DeveloperDelete(c *gin.Context) {
	// Get Email
	var data struct {
		Email    string
		Password string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	}

	// Check if Developer exists
	var dev = models.Developer{Email: data.Email}
	database.DB.First(&dev)

	if dev.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dev",
		})

		return
	}

	// Validate Password
	err := bcrypt.CompareHashAndPassword([]byte(dev.Password), []byte(data.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})

		return
	}

	// Delete User
	result := database.DB.Delete(&dev)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete dev",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Developer Deleted",
	})

}

// Admin Priveleges
func DeveloperList(c *gin.Context) {
	// Get all developers from database
	var devs []models.Developer
	database.DB.Omit("Password").Omit("UUID").Omit("IsAdmin").Find(&devs)

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"devs": devs,
	})
}

func DeveloperResetPassword(c *gin.Context) {
	// Get Email, OldPassword, NewPassword, ConfirmPassword
	var data struct {
		Email           string
		OldPassword     string `json:"oldpassword"`
		NewPassword     string `json:"newpassword"`
		ConfirmPassword string `json:"confirmpassword"`
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})
	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.OldPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Old Password not Provided",
		})
		return
	} else if data.NewPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Password not Provided",
		})
		return
	} else if data.ConfirmPassword == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Confirm Password not Provided",
		})
		return
	}

	// Check if Developer exists
	var dev = models.Developer{Email: data.Email}
	database.DB.First(&dev)

	if dev.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid dev",
		})

		return
	}

	// Check the OldPassword if it matches in the database
	err := bcrypt.CompareHashAndPassword([]byte(dev.Password), []byte(data.OldPassword))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid old password",
		})

		return
	}
	// Check if NewPassword and ConfirmPassword are the same
	if data.NewPassword != data.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Passwords do not match",
		})
	}

	// Hash the NewPassword
	hash, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	// Update the Developer
	result := database.DB.Model(&dev).Updates(models.Developer{
		Password: string(hash),
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update password",
		})

		return
	}
	// Respond
	c.Status(http.StatusOK)
}

func DeveloperEmailChange(c *gin.Context) {
	// Get Email
	var data struct {
		Email    string
		NewEmail string
		Password string
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if data.NewEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "New Email not Provided",
		})
		return
	} else if data.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password not Provided",
		})
		return
	}

	// Check if Developer exists
	var dev = models.Developer{Email: data.Email}
	database.DB.First(&dev)

	if dev.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid developer",
		})

	}

	// Validate Password
	err := bcrypt.CompareHashAndPassword([]byte(dev.Password), []byte(data.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})

		return
	}

	// Update Developer
	result := database.DB.Model(&dev).Updates(models.Developer{
		Email: data.NewEmail,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update developer",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Developer Updated",
	})
}

// Admin Priveleges
func DeveloperIsActiveChange(c *gin.Context) {
	// Get Email
	var data struct {
		Email    string
		isActive bool
	}

	if c.Bind(&data) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read data",
		})

		return

	}

	if data.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email not Provided",
		})
		return
	} else if !data.isActive {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "isActive not Provided",
		})
		return
	}

	// Check if Developer exists
	var dev = models.Developer{Email: data.Email}
	database.DB.First(&dev)

	if dev.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid developer",
		})

	}

	// Check if Developer is admin
	if dev.IsAdmin {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Developer has Admin priveleges",
		})

		return
	}

	// Update Developer
	result := database.DB.Model(&dev).Updates(models.Developer{
		IsActive: data.isActive,
	})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update developer",
		})

		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{
		"Message": "Developer Updated",
	})
}
