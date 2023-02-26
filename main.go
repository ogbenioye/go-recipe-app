package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Keywords     []string  `json:"keywords"`
	Ingridents   []string  `json:"ingridents"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func IndexHandler(c *gin.Context) {
	username := c.Params.ByName("username")
	c.JSON(200, gin.H{
		"message": "Hello," + username,
	})
}

func main() {
	router := gin.Default()
	router.GET("/:username", IndexHandler)

	router.Run()
}
