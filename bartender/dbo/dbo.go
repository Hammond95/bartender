package dbo

import (
	"database/sql"

	models "github.com/Hammond95/bartender/bartender/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCocktails(mongo *mongo.Client) ([]models.Cocktail, error) {
	cocktails := [...]models.Cocktail{
		{
			Name:        "Gin Tonic",
			Image:       sql.NullString{},
			Ingredients: []models.Ingredient{},
			IsAlcoholic: true,
			Tags:        []string{"bitter", "classic"},
		},
		{},
	}

	//mongo.Database("local").Collection("cocktails")

	return cocktails[:], nil
}

func GetCocktail(mongo *mongo.Client, name string) (models.Cocktail, error) {
	return models.Cocktail{
		Name:  name,
		Image: sql.NullString{},
		Ingredients: []models.Ingredient{
			{Name: "gin", Type: "liquor", Image: sql.NullString{}},
			{Name: "tonic water", Type: "beverage", Image: sql.NullString{}},
			{Name: "lime", Type: "fruit", Image: sql.NullString{}},
		},
		IsAlcoholic: true,
		Tags:        []string{"bitter", "classic"},
	}, nil
}

func GetCocktailIngredients(mongo *mongo.Client, name string) ([]models.Ingredient, error) {
	cocktail, err := GetCocktail(mongo, name)
	return cocktail.Ingredients, err
}
