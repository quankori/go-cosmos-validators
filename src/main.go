package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quankori/go-cosmos-validators/src/internal/redis"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		// r := make(chan string)
		go func() {
			for {
				result := redis.GetMessage("key")
				fmt.Println(result[0])
				time.Sleep(time.Second)
			}
		}()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/pong/:text", func(c *gin.Context) {
		redis.PushMessage("key", c.Param("text"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
