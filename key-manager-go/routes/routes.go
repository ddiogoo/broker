package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

/*
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
*/
