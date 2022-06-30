package main

import (
	"fmt"
	"net/http"
	"sync"

	"time"

	"eeyoo.com/kee/pkg/downloader"
	"github.com/gin-gonic/gin"
)

func init() {
	downloader.Hi()
}

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })

	r.Run()
}

type Person struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Id       int
	Birthday time.Time
}

// 语义
func hello() {
	fmt.Println("hello world!")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}

type Hello struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type World struct {
	num sync.Mutex
}
