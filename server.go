package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/js", "./js")

	router.GET("/", func(context *gin.Context) {
		context.File("./index.html")
	})

	router.GET("/wasm", func(context *gin.Context) {
		context.Header("Content-Type", "application/wasm")
		context.File("./lib.wasm")
	})

	err := router.Run("localhost:9000")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("service running")
	}
}
