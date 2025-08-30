package main

import (
	"demo/database"
	"demo/middlewares"
	"demo/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()

	router := gin.Default()

	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middlewares.Logger())

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.ConnMysql()
	}()

	router.Run(":8000")
}
