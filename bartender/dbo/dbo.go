package dbo

import (
	models "github.com/Hammond95/bartender/bartender/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCocktails(mongo *mongo.Client) ([]models.Cocktail, error) {
	cocktails := [...]models.Cocktail{{}, {}}

	mongo.Database("local").Collection("cocktails")

	return cocktails[:], nil
}
