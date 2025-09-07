package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// v1
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// v2
	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"github.com/thinkerou/favicon"
)

var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/gintest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func main() {
	// 连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close() // v2 不需要这一步
	DB.AutoMigrate(&Todo{}) // 模型关联

	r := gin.Default()

	r.Use(favicon.New("favicon.ico"))
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {
			// 1. get data
			var todo Todo
			c.BindJSON(&todo)
			// 2. save to table
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
			// 3. return response
		})

		// 查看所有
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询 todo 表中所有的数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		// 查看某一个
		// v1Group.GET("/todo/:id", func(c *gin.Context) {})

		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "wrong id"})
				return
			}

			var todo Todo
			if err := DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}

			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}

			if err := DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"id": id,
					"message": "be deleted",
				})
			}
		})
	}

	r.Run()
}
