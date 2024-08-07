package routes

import (
	"net/http"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/dto"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/generator"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/model"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/mongodb"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, client *mongodb.MongoClient) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/generate", func(c *gin.Context) {
		var keyDto dto.KeyDto
		if err := c.BindJSON(&keyDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid body",
			})
			return
		}
		key := model.NewKey(keyDto.Email, generator.GenerateApiKey(), "/ping")
		res, err := client.InsertOne(key)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error during generate a key",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "successfully during generate key",
			"key":     key,
			"id":      res,
		})
	})
}
