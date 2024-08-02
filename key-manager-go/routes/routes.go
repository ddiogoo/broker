package routes

import (
	"fmt"
	"net/http"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/ctx"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/dto"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/generator"
	"github.com/ddiogoo/broker/tree/master/key-manager-go/model"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, db *ctx.KeyManagerDatabase) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("", func(c *gin.Context) {
		var keyDto dto.KeyDto
		if err := c.BindJSON(&keyDto); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid body to KeyDto",
			})
			return
		}
		key := model.NewKey(keyDto.Email, generator.GenerateApiKey(), "/ping")
		fmt.Println("Email: " + key.Email)
		fmt.Println("ApiKey: " + key.ApiKey)
		fmt.Println("Routes: " + key.Routes)
	})
}
