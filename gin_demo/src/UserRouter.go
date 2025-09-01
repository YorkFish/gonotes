package src

import (
	session "demo/middlewares"
	"demo/pojo"
	"demo/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users", session.SetSession())

	// MySql
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)
	// user.GET("/", service.FindAllUsers)
	// user.GET("/:id", service.FindByUserId)
	user.PUT("/:id", service.PutUser)

	user.POST("/login", service.LoginUser)
	user.GET("/check", service.CheckUserSession)

	user.Use(session.AuthSession())
	{
		user.DELETE("/:id", service.DeleteUser)
		user.GET("/logout", service.LogoutUser)
	}

	// Redis
	user.GET("/", service.CatchUserAllDecorator(service.RedisAllUser, "user_all", pojo.User{}))
	user.GET("/:id", service.CatchOneUserDecorator(service.RedisOneUser, "id", "user_%s", pojo.User{}))
}
