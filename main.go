package main

import "github.com/gin-gonic/gin"

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
