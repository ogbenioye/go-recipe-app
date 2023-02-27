package main

import (
	"net/http"
	"time"

	"github.com/rs/xid"

	"github.com/gin-gonic/gin"
)

var recipes []Recipe

type Recipe struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Keywords     []string  `json:"keywords"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt"`
}

func init() {
	recipes = make([]Recipe, 0)
}

func ListAllRecipeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	recipe.Id = xid.New().String()
	recipe.PublishedAt = time.Now()

	recipes = append(recipes, recipe)

	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipe", NewRecipeHandler)
	router.GET("/recipe", ListAllRecipeHandler)
	router.Run()
}
