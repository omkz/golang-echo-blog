package controllers

import (
	"github.com/omkz/golang-echo-blog/models"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	result := models.PostAll()
	return c.JSON(http.StatusOK,result)
}