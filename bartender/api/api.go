package api

import (
	//"encoding/json"
	//"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/mongo"
)

type HandlersEnv struct {
	MongoDB *mongo.Client
	Logger  hclog.Logger
}

func SetV1RouteGroupDefinition(router *gin.Engine, env HandlersEnv) {
	apiV1 := router.Group("/v1")
	V1DefineCocktailRoutes(apiV1, env)
}
