package main

import (
	. "demo/src"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	AddUserRouter(v1)

	router.Run(":8000")
}
