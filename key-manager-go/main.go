package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/test"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(func() string {
		if os.Getenv("GIN_RUN_MODE") == "debug" {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var test test.TestDto
		if err := c.BindJSON(&test); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid body",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "body is ok",
		})
		fmt.Println(test.Email)
	})
	r.Run()
}
