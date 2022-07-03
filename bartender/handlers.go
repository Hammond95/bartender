package main

import (
	"encoding/json"
	"net/http"

	"github.com/Hammond95/bartender/bartender/version"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"message": "Hello World!"},
	)
}

func InfoHandler(c *gin.Context) {
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

func LivenessHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"status": "UP"},
	)
}

func ReadinessHandler(c *gin.Context) {
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

	c.JSON(
		http.StatusOK,
		gin.H{"status": "UP"},
	)
}
