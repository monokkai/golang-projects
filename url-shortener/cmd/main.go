package main

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/handler"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", handler.HandleShorten)
	r.GET("/:code", handler.HandleRedirect)

	r.Run(":8080")
}
