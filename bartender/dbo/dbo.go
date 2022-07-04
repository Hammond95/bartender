package dbo

import (
	"context"
	"database/sql"

	models "github.com/Hammond95/bartender/bartender/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCocktails(mongo *mongo.Client) ([]models.Cocktail, error) {
	cocktailsCollection := mongo.Database("local").Collection("cocktails")

	// TODO: Add Pagination
	cursor, err := cocktailsCollection.Find(context.TODO(), &bson.M{})
	if err != nil {
		panic(err)
	}

	defer cursor.Close(context.TODO())

	var results []models.Cocktail
	cursor.All(context.TODO(), &results)
	if err != nil {
		// TODO: Find a way to manage different errors of mongodb with appropriate http codes
		// e.g. Unauthorized should be associated with 401/403
		//      Mongo not reachable or failing with 500/503
		panic(err)
	}

	return results[:], err
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
