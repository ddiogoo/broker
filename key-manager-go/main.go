package main

import (
	"context"
	"os"
	"time"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
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
	routes.ConfigureRoutes(r)
	r.Run()
}
