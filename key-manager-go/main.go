package main

import (
	"os"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/database"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/routes"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	reflectList := util.BuildListOfReflectType()
	db, err := database.NewDatabaseManager(reflectList)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	gin.SetMode(func() string {
		if os.Getenv("GIN_RUN_MODE") == "debug" {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())
	r := gin.Default()
	routes.ConfigureRoutes(r, db)
	r.Run()
}
