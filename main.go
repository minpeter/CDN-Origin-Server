package main

import (
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	store := persistence.NewInMemoryStore(time.Second)
	router.GET("/images/:image", cache.CachePage(store, time.Minute, images))
	router.POST("/upload", upload)
	router.Run(":8888")
}
