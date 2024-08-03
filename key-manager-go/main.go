package main

import (
	"log"
	"os"
	"reflect"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/builder"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/database"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/model"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	test()
	keyManagerDb, err := database.NewKeyManagerDatabase()
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

func test() {
	typs := []reflect.Type{}
	keyType := reflect.TypeOf(&model.Key{}).Elem()
	typs = append(typs, keyType)
	manager := builder.NewDatabaseQueryBuilder(typs)
	manager.ChargeTypes()
	manager.CreateAllTable()
}
