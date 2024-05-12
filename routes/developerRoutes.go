package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/dev-go/controllers"
	"github.com/immanuel-supanova/dev-go/middleware"
)

func DeveloperRoutes(r *gin.Engine) {
	route := r.Group("/dev")

	route.POST("/create", controllers.DeveloperCreate)
	route.POST("/create-admin", controllers.DeveloperCreateAdmin)
	route.GET("/current", controllers.DeveloperCurrent)

	route.DELETE("/delete", middleware.RequireAuth, controllers.DeveloperDelete)
	route.PUT("/reset-password", middleware.RequireAuth, controllers.DeveloperResetPassword)
	route.PUT("/email-change", middleware.RequireAuth, controllers.DeveloperEmailChange)

	route.Use(middleware.RequireAdmin)

	route.GET("/read", controllers.DeveloperRead)
	route.GET("/list", controllers.DeveloperList)
	route.PUT("/isactive", controllers.DeveloperIsActiveChange)
}
