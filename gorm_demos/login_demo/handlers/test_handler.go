package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"demo/models"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

func PostHandler(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")
	day := c.Param("day")
	c.JSON(http.StatusOK, gin.H{
		"year":  year,
		"month": month,
		"day":   day,
	})
}

func FormHandler(c *gin.Context) {
	objA := models.FormA{}
	objB := models.FormB{}
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"errA": errA.Error(),
			"errB": errB.Error(),
		})
	}
}

func JsonHandler(c *gin.Context) {
	objA := models.FormA{}
	objB := models.FormB{}
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"errA": errA.Error(),
			"errB": errB.Error(),
		})
	}
}
