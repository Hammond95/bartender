package v1

import (
	//"encoding/json"
	//"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
)

func SetV1RouteGroupDefinition(router *gin.Engine) {
	apiV1 := router.Group("/v1")
	DefineCocktailRoutes(apiV1)
}
