package controllers

import (
	"github.com/omkz/golang-echo-blog/models"

	"log"
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetPosts(c echo.Context) error {
	result := models.PostAll()
	return c.JSON(http.StatusOK,result)
}

func CreatePost(c echo.Context) error {
	var post models.Post

	err := c.Bind(&post)

	if err != nil {
		log.Print(err)
	}

	err = models.PostCreate(&post)

	if err != nil {
		log.Print(err)
	}

	return c.JSON(http.StatusOK, post)

}