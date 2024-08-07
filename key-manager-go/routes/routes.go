package routes

import (
	"net/http"
	"strings"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/dto"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/generator"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/model"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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

	r.POST("/check", func(c *gin.Context) {
		var checkPermissionDto dto.CheckPermissionDto
		if err := c.BindJSON(&checkPermissionDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid body",
			})
			return
		}
		result, err := client.FindOne(checkPermissionDto)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "key not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "error checking permission",
			})
			return
		}

		if strings.Contains(result.Route, checkPermissionDto.Route) {
			c.JSON(http.StatusOK, gin.H{
				"message": "permission granted",
			})
		} else {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "permission denied",
			})
		}
	})
}
