package main

import (
	"github.com/gin-gonic/gin"

	"tech-friday/internal/health"
)

func main() {
	r := gin.Default()

	r.GET("/health", health.Handler)

	r.Run(":8080")
}
