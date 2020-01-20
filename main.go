package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
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

  // Routes
  e.GET("/", hello)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}