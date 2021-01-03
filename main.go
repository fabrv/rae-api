package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 200,
			"data":   "success",
		})
	})

	r.GET("/pdd", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"word":       wordOfTheDay(),
			"definition": getDefinition(wordOfTheDay()),
		})
	})

	r.GET("/definicion/:word", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"word":       c.Param("word"),
			"definition": getDefinition("/" + c.Param("word")),
		})
	})
	r.Run(":" + port)
}
