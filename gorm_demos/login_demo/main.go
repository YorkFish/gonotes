package main

import (
	"github.com/gin-gonic/gin"

	"demo/routes"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routes.RegisterRoutes(r)

	r.Run()
}
