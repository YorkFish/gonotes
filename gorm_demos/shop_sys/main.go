package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func castTime(c *gin.Context) {
	fmt.Println(">>> castTime start")

	start := time.Now()
	c.Next()
	cost := time.Since(start)
	fmt.Println("cost:", cost)

	fmt.Println(">>> castTime end")
}

func uploadShowHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "file/upload.html", nil)
}

func uploadDealHandler(c *gin.Context) {
	// 提取用户上传的文件
	fileObj, err := c.FormFile("filename")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "file is invalid",
			"error": err.Error(),
		})
		return
	}

	log.Println(fileObj.Filename)
	dst := fmt.Sprintf("./files/%s", fileObj.Filename)
	// 上传文件到指定的目录，默认 32MiB
	c.SaveUploadedFile(fileObj, dst)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  fmt.Sprintf("upload %s success!", fileObj.Filename),
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/file/upload.html")

	shoppingGroup := r.Group("/shopping")
	{
		shoppingGroup.GET("/index", shopIndexHandler)
		// 添加中间件
		shoppingGroup.GET("/home", castTime, shopHomeHandler)

		// 可以嵌套
		liveGroup := shoppingGroup.Group("/live")
		// 添加中间件
		liveGroup.Use(castTime)
		{
			liveGroup.GET("/index", liveIndexHandler)
			liveGroup.GET("/home", liveHomeHandler)
		}
	}

	r.GET("/upload", uploadShowHandler)
	r.POST("/upload", uploadDealHandler)

	r.Run()
}
