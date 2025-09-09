package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"demo/models"
)

func LoginHandler(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}

func LoginHandlerV2(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		// PUT, DELETE 等都会走这里，增加了不必要的消耗
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	}
}

func LoginHandlerV3(c *gin.Context) {
	// TODO: 去数据库中校验

	if c.Request.Method == "POST" {
		var u models.UserInfo
		// 根据请求头中的 Content-Type 解析
		err := c.ShouldBind(&u)
		// 若解析数据出问题
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"username": u.Username,
			"password": u.Password,
		})
	} else {
		c.HTML(http.StatusOK, "login.html", nil)
	}
}
