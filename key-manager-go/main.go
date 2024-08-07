package main

import (
	"context"
	"os"
	"time"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/mongodb"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Env
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	// MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongodb.NewMongoClient(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(); err != nil {
			panic(err)
		}
	}()
	err = client.Ping()
	if err != nil {
		panic(err)
	}
	// Gin
	gin.SetMode(func() string {
		if os.Getenv("GIN_RUN_MODE") == "debug" {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())
	r := gin.Default()
	routes.ConfigureRoutes(r, client)
	r.Run()
}
