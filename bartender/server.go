package main

import (
	"fmt"
	"path/filepath"

	api "github.com/Hammond95/bartender/bartender/api"
	"github.com/gin-gonic/gin"
)

func SetupServer(
	env api.HandlersEnv,
	address string,
	staticAssetsPath string,
	trustedProxies arrayFlags,
) *gin.Engine {

	g := gin.Default()

	//var absStaticAssetsPath string
	if staticAssetsPath != "" {
		absStaticAssetsPath, err := filepath.Abs(staticAssetsPath)
		if err != nil {
			env.Logger.Error("Couldn't parse the provided path for static files!")
		}
		// This will panic if the provided path for the templates doesn't have html files.
		g.Static("/assets", filepath.Join(absStaticAssetsPath, "assets"))
		g.LoadHTMLGlob(filepath.Join(absStaticAssetsPath, "templates/*.html"))
	} else {
		env.Logger.Error(fmt.Sprintf("Couldn't parse the provided path for static files, value was %v.", staticAssetsPath))
	}

	if len(trustedProxies) > 0 {
		g.SetTrustedProxies(trustedProxies)
	} else {
		// If we don't use any proxy, we can disable this feature
		g.SetTrustedProxies(nil)
	}

	SetupServerRoutes(env, g)

	return g
}

func SetupServerRoutes(env api.HandlersEnv, g *gin.Engine) {
	api.SetV1RouteGroupDefinition(g)

	g.GET("/hello", env.HelloHandler)
	g.GET("/info", env.InfoHandler)
	g.GET("/liveness", env.LivenessHandler)
	g.GET("/readiness", env.ReadinessHandler)
}
