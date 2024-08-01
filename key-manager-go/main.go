package main

import (
	"log"
	"os"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/ctx"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	keyManagerDb, err := ctx.NewKeyManagerDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer keyManagerDb.Close()
	err = keyManagerDb.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("Successfully connection with PostgreSQL.")
	gin.SetMode(func() string {
		if os.Getenv("GIN_RUN_MODE") == "debug" {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())
	r := gin.Default()
	routes.ConfigureRoutes(r, keyManagerDb)
	r.Run()
}
