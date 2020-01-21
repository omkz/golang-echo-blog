package controllers

import (
	"github.com/omkz/golang-echo-blog/models"
	"net/http"
	"github.com/labstack/echo"
)

func GetPosts(c echo.Context) error {
	result := models.PostAll()
	return c.JSON(http.StatusOK,result)
}