package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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
	r.Run()
}
