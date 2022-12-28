package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	userAgent := ""

	engine.Use(func(c * gin.Context) {
		userAgent = c.GetHeader("User-Agent")
		c.Next()
	})

	engine.GET("/", func(c * gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
			"User-Agent": userAgent,
		})
	})
	engine.Run(":3000")
}
