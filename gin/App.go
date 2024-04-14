package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	Message  string    `json:"message" binding:"required"`
	Datetime time.Time `time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, Response{
			"pong",
			time.Now(),
		})
	})

	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, Response{
			"post",
			time.Now(),
		})
	})

	r.Run(":8080")
}
