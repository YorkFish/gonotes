package main

import (
	"demo/database"
	"demo/middlewares"
	"demo/pojo"
	"demo/src"

	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func setupLogging() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogging()

	router := gin.Default()

	// 注册 validator func
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("userpasd", middlewares.UserPasd)
		v.RegisterStructValidation(middlewares.UserList, pojo.Users{})
	}

	router.Use(gin.Recovery(), middlewares.Logger())

	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	go func() {
		database.ConnMysql()
	}()

	router.Run(":8000")
}
