package devgo

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/immanuel-supanova/dev-go/middleware"
	"github.com/immanuel-supanova/dev-go/routes"
)

func devgoRouter() http.Handler {
	debug := os.Getenv("DEBUG")

	if debug == "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.LoggerMiddleware(gin.LoggerConfig{}, "devgo"))

	routes.DeveloperRoutes(r)
	routes.ApplicationRoutes(r)

	return r
}

var DevGoServer = &http.Server{
	Addr:         ":3000",
	Handler:      devgoRouter(),
	ReadTimeout:  5 * time.Second,
	WriteTimeout: 10 * time.Second,
}
