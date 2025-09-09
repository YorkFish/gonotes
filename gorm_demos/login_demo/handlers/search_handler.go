package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchHandler(c *gin.Context) {
	name := c.Query("name")
	age := c.DefaultQuery("age", "18")
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
