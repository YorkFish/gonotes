package src

import (
	session "demo/middlewares"
	"demo/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", session.SetSession())

	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindByUserId)
	user.PUT("/:id", service.PutUser)

	user.POST("/login", service.LoginUser)
	user.GET("/check", service.CheckUserSession)

	user.Use(session.AuthSession())
	{
		user.DELETE("/:id", service.DeleteUser)
		user.GET("/logout", service.LogoutUser)
	}
}
