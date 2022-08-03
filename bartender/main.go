package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/Hammond95/bartender/bartender/api"
	"github.com/Hammond95/bartender/bartender/version"
	hclog "github.com/hashicorp/go-hclog"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	staticAssetsPath            = flag.String("static-assets-path", "./static", "The path were gin tonic can find his static files.")
	address                     = flag.String("address", ":8888", "TCP host+port where you want to start gin tonic.")
	trustedProxies   arrayFlags = arrayFlags{}
)

func main() {
	flag.Var(&trustedProxies, "trusted", "specify network addresses or network CIDRs from where request headers related to client IP can be trusted.")
	flag.Parse()
	log := hclog.Default()

	log.Info(
		fmt.Sprintf(
			"Starting Bartender @[commit: %s, build time: %s, release: %s]",
			version.Commit, version.BuildTime, version.Release,
		),
	)

	host, port, err := net.SplitHostPort(*address)
	if err != nil {
		log.Error("failed to parse address (%q): %v", *&address, err)
		log.Error("address was parsed as host = %v, port = %v", host, port)
		os.Exit(1)
	}

	mongoContext, mongoContextCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer mongoContextCancel()

	mongoClient, err := mongo.Connect(
		mongoContext,
		options.Client().ApplyURI(
			fmt.Sprintf("mongodb://localhost:%v", os.Getenv("MONGODB_PORT")),
		),
		// TODO: Provide Built Config
		// will implement in config.go
		options.Client().SetAuth(
			options.Credential{
				Username: os.Getenv("MONGODB_USERNAME"),
				Password: os.Getenv("MONGODB_PASSWORD"),
			},
		),
	)
	if err != nil {
		panic(err)
	}
	if err := mongoClient.Ping(mongoContext, nil); err != nil {
		panic(err)
	}

	env := api.HandlersEnv{MongoDB: mongoClient, Logger: log}

	g := SetupServer(env, *address, *staticAssetsPath, trustedProxies)

	srv := &http.Server{
		Addr:    *address,
		Handler: g,
	}

	go func() {
		// service connections
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error(fmt.Sprintf("Failed to run server: %v", err))
			return
		}
	}()

	// TODO: Better management of signals
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("Server Shutdown: %v", err))
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info("timeout of 5 seconds.")
	}
	log.Info("Server exiting.")
}
