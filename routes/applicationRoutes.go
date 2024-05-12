package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/dev-go/controllers"
	"github.com/immanuel-supanova/dev-go/middleware"
)

func ApplicationRoutes(r *gin.Engine) {
	route := r.Group("/app")

	route.POST("/create", controllers.ApplicationCreate)

	route.GET("/read", middleware.RequireAuth, controllers.ApplicationRead)
	route.PUT("/update", middleware.RequireAuth, controllers.ApplicationUpdate)
	route.DELETE("/delete", middleware.RequireAuth, controllers.ApplicationDelete)

	route.POST("/token", controllers.ApplicationGetToken)
	route.GET("/token-refresh", controllers.ApplicationGetAccessToken)

	route.Use(middleware.RequireAdmin)
	route.GET("/list", controllers.ApplicationList)
	route.PUT("/isactive", controllers.ApplicationUpdate)
}
