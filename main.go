package main

import (
	"github.com/labstack/echo/v4"
	"github.com/omkz/golang-echo-blog/controllers"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
  }))
  
  e.Use(middleware.CORS())

	// Routes
  e.GET("/", hello)
  e.GET("/posts", controllers.GetPosts)
  e.POST("/posts", controllers.CreatePost)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
