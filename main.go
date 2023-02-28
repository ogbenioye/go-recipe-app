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

func DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}

	recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recipe deleted successfully",
	})
}

func UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")

	var recipe Recipe

	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	index := -1

	for i := 0; i < len(recipes); i++ {
		if recipes[i].Id == id {
			index = i
		}
	}

	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Recipe not found",
		})
		return
	}

	// recipe.Id = id
	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)
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
	router.PUT("/recipe/:id", UpdateRecipeHandler)
	router.DELETE("/recipe/:id", DeleteRecipeHandler)
	router.Run()
}
