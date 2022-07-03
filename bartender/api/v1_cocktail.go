package api

import (
	"database/sql"
	"net/http"

	"github.com/Hammond95/bartender/bartender/dbo"
	models "github.com/Hammond95/bartender/bartender/models"
	"github.com/gin-gonic/gin"
)

func RecipeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "recipe.html", nil)
}

func (env *HandlersEnv) GetCocktailsHandler(c *gin.Context) {

	cocktails, err := dbo.GetCocktails(env.MongoDB)
	if err != nil {
		panic(err)
	}

	c.JSON(
		http.StatusOK,
		gin.H{"cocktails": cocktails},
	)
}

func V1DefineCocktailRoutes(group *gin.RouterGroup) {
	ckt := group.Group("/cocktail")

	ckt.GET("/ingredients", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ingredients": []models.Ingredient{
				{Name: "gin", Type: "liquor", Image: sql.NullString{}},
				{Name: "tonic water", Type: "beverage", Image: sql.NullString{}},
				{Name: "lime", Type: "fruit", Image: sql.NullString{}},
			},
		})
	})

	ckt.GET("/recipe", RecipeHandler)
}
