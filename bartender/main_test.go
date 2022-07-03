package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/Hammond95/bartender/bartender/api"
	"github.com/gin-gonic/gin"
	hclog "github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestHelloHandler(t *testing.T) {
	log := hclog.Default()

	env := api.HandlersEnv{Logger: log}

	mockResponse := `{"message":"Hello World!"}`
	r := SetUpRouter()
	r.GET("/hello", env.HelloHandler)
	req, _ := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}
