package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Hammond95/bartender/bartender/version"
	"github.com/gin-gonic/gin"
	hclog "github.com/hashicorp/go-hclog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type HandlersEnv struct {
	mongodb *mongo.Client
	logger  hclog.Logger
}

func (env *HandlersEnv) HelloHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"message": "Hello World!"},
	)
}

func (env *HandlersEnv) InfoHandler(c *gin.Context) {
	info := struct {
		App       string `json:"app"`
		BuildTime string `json:"buildTime"`
		Commit    string `json:"commit"`
		Release   string `json:"release"`
	}{
		"Bartender", version.BuildTime, version.Commit, version.Release,
	}

	_, err := json.Marshal(info)
	if err != nil {
		c.AbortWithError(
			http.StatusInternalServerError,
			err,
		)
	}

	c.JSON(
		http.StatusOK,
		info,
	)
}

func (env *HandlersEnv) LivenessHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"status": "UP"},
	)
}

func (env *HandlersEnv) ReadinessHandler(c *gin.Context) {
	// If we add calls to external resources (e.g. a database),
	// we should also perform readiness
	// checks to those resources.

	/* TODO:
	Idea: Each resource should provide a method .IsReady() to be called by this handler.
	If any resource is not ready, the Handler should return a not OK code = http.StatusServiceUnavailable.

	if isReady == nil || !isReady {
		c.AbortWithError(
			http.StatusServiceUnavailable,
			err,
		)
	}
	*/
	err := env.mongodb.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		c.AbortWithError(
			http.StatusServiceUnavailable,
			err,
		)
	}

	c.JSON(
		http.StatusOK,
		gin.H{"status": "UP"},
	)
}
