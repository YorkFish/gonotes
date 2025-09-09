package routes

import (
	"github.com/gin-gonic/gin"

	"demo/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/index", handlers.IndexHandler)
	r.GET("/post/:year/:month/:day", handlers.PostHandler)

	r.Any("/form", handlers.FormHandler)
	r.Any("/json", handlers.JsonHandler)

	r.GET("/search", handlers.SearchHandler)

	r.Any("/login", handlers.LoginHandler)
	r.Any("/login3", handlers.LoginHandlerV3)
}
