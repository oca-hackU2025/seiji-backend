package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// インデックスルートに簡単なメッセージを返す
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from Gin!",
		})
	})

	r.Run() // デフォルトで :8080 ポートで起動
}
