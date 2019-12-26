package main

import "github.com/gin-gonic/gin"

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello": "World",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
