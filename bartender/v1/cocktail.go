package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ingredient struct {
	Name     string `json:"name" xml:"Name" yaml:"name"`
	Quantity string `json:"quantity" xml:"Quantity" yaml:"quantity"`
}

func RecipeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "recipe.html", nil)
}

func DefineCocktailRoutes(group *gin.RouterGroup) {
	ckt := group.Group("/cocktail")

	ckt.GET("/ingredients", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ingredients": []Ingredient{
				{"gin", "1/3"},
				{"tonic water", "2/3"},
				{"lime", "some slices"},
			},
		})
	})

	ckt.GET("/recipe", RecipeHandler)
}
