package main

import (
	"fmt"
	"path/filepath"

	v1 "github.com/Hammond95/bartender/bartender/v1"
	"github.com/gin-gonic/gin"
	hclog "github.com/hashicorp/go-hclog"
)

func SetupServer(logger hclog.Logger, address string, staticAssetsPath string, trustedProxies arrayFlags) *gin.Engine {

	g := gin.Default()

	//var absStaticAssetsPath string
	if staticAssetsPath != "" {
		absStaticAssetsPath, err := filepath.Abs(staticAssetsPath)
		if err != nil {
			logger.Error("Couldn't parse the provided path for static files!")
		}
		// This will panic if the provided path for the templates doesn't have html files.
		g.Static("/assets", filepath.Join(absStaticAssetsPath, "assets"))
		g.LoadHTMLGlob(filepath.Join(absStaticAssetsPath, "templates/*.html"))
	} else {
		logger.Error(fmt.Sprintf("Couldn't parse the provided path for static files, value was %v.", staticAssetsPath))
	}

	if len(trustedProxies) > 0 {
		g.SetTrustedProxies(trustedProxies)
	} else {
		// If we don't use any proxy, we can disable this feature
		g.SetTrustedProxies(nil)
	}

	SetupServerRoutes(g)

	return g
}

func SetupServerRoutes(g *gin.Engine) {
	v1.SetV1RouteGroupDefinition(g)

	g.GET("/hello", HelloHandler)
	g.GET("/info", InfoHandler)
	g.GET("/liveness", LivenessHandler)
	g.GET("/readiness", ReadinessHandler)
}
