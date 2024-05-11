package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/dev-go/controllers"
	"github.com/immanuel-supanova/dev-go/middleware"
)

func DeveloperRoutes(r *gin.Engine) {
	r.POST("/create", controllers.DeveloperCreate)
	r.POST("/create-admin", controllers.DeveloperCreateAdmin)
	r.DELETE("/delete", controllers.DeveloperDelete)
	r.PUT("/reset-password", controllers.DeveloperResetPassword)
	r.PUT("/email-change", controllers.DeveloperEmailChange)

	r.Use(middleware.RequireAdmin)

	r.GET("/read", controllers.DeveloperRead)
	r.GET("/list", controllers.DeveloperList)
	r.PUT("/isactive", controllers.DeveloperIsActiveChange)
}
