package dbo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	models "github.com/Hammond95/bartender/bartender/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func paginate(collection *mongo.Collection, ginContext *gin.Context) (*[]models.Cocktail, error) {
	var results []models.Cocktail
	var cursor *mongo.Cursor
	var err error = nil
	var lastId primitive.ObjectID

	perPage, err := strconv.ParseInt(ginContext.DefaultQuery("perPage", "30"), 10, 64)
	fmt.Printf("PER PAGE = %v", perPage)
	lastIdParam := ginContext.Query("lastId")

	if lastIdParam == "" {
		lastId = primitive.NilObjectID
	} else {
		lastId, err = primitive.ObjectIDFromHex(lastIdParam)
		fmt.Printf("LAST ID = %v, %v\n", lastIdParam, lastId.String())
	}

	// TODO: Add a check to avoid perPage=0 or perPage>BIG
	var opts = options.Find().SetSort(&bson.D{{Key: "_id", Value: 1}}).SetLimit(int64(perPage))

	if lastId == primitive.NilObjectID {
		// First Page
		cursor, err = collection.Find(context.TODO(), &bson.M{}, opts)
	} else {
		filter := bson.M{"_id": bson.M{"$gt": lastId}}
		cursor, err = collection.Find(context.TODO(), filter, opts)
	}

	for {
		if cursor.TryNext(context.TODO()) {
			var result models.Cocktail
			if err := cursor.Decode(&result); err != nil {
				//TODO: change with logger
				panic(err)
			}
			//fmt.Println(result)
			results = append(results, result)
			continue
		}
		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}
		if cursor.ID() == 0 {
			break
		}
	}

	//cursor.All(context.TODO(), &results)
	if len(results) > 0 {
		//lastId = results[len(results)-1].Id
		fmt.Println(ginContext.GetQueryArray("lastId"))
		fmt.Println(ginContext.Params)
	}

	//fmt.Println(results)

	return &results, err
}

func GetCocktails(mongo *mongo.Client, ginContext *gin.Context) ([]models.Cocktail, error) {
	cocktailsCollection := mongo.Database("local").Collection("cocktails")

	// TODO: Add Real Pagination
	results, err := paginate(cocktailsCollection, ginContext)
	if err != nil {
		panic(err)
	}

	/*
		if err != nil {
			// TODO: Find a way to manage different errors of mongodb with appropriate http codes
			// e.g. Unauthorized should be associated with 401/403
			//      Mongo not reachable or failing with 500/503
			panic(err)
		}*/

	return (*results)[:], err
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
