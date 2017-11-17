package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	user := route.Group("/user")
	user.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})
	route.Run("127.0.0.1:8899")
}
