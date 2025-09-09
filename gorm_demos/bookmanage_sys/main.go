package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// newBookHandler 返回添加书籍页面的处理函数
func newBookHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "book/new_book.html", nil)
}

// createBookHandler 添加书籍
func createBookHandler(c *gin.Context) {
	// 1. 从表单获取数据
	title := c.PostForm("title")
	price := c.PostForm("price")
	priceVal, err := strconv.ParseFloat(price, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "price is invalid",
			"error": err.Error(),
		})
		return
	}

	// 2. 插入数据到数据库
	err = insertBook(title, priceVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "insert failed",
			"error": err.Error(),
		})
		return
	}

	// 3. 重定向
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func bookListHandler(c *gin.Context) {
	// 查数据
	bookList, err := queryAllBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "query failed",
			"error": err.Error(),
		})
		return
	}

	// return json
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 0,
	// 	"data": bookList,
	// })

	// return html
	c.HTML(http.StatusOK, "book/book_list.html", gin.H{
		"code": 0,
		"data": bookList,
	})
}

func bookDetailHandler(c *gin.Context) {
	// 1. 获取书籍 id
	// 2. 去数据库拿到具体的书籍信息
	// 3. 返回 JSON 格式的数据
	bookId := c.Param("id")
	if len(bookId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "id is invalid",
		})
		return
	}

	bookIdVal, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"msg":   "id is invalid",
			"error": err.Error(),
		})
		return
	}

	bookObj, err := queryBookInfo(bookIdVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  1,
			"msg":   "query failed",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": bookObj,
	})
}

func editBookHandler(c *gin.Context) {
	bookId := c.Query("id")
	if len(bookId) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 11,
			"msg":  "id is invalid",
		})
		return
	}
	fmt.Println("bookId:", bookId)

	idVal, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  12,
			"msg":   "id is invalid",
			"error": err.Error(),
		})
		return
	}
	fmt.Println("idVal:", idVal)

	if c.Request.Method == "POST" {
		// 1. 获取用户提交的数据
		// 2. 去数据库中找对应数据更新
		// 3. 跳转回 /book/list

		title := c.PostForm("title")
		price := c.PostForm("price")
		fmt.Printf("title: %v, price%v\n", title, price)
		priceVal, err := strconv.ParseFloat(price, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  1,
				"msg":   "price is invalid",
				"error": err.Error(),
			})
			return
		}

		err = editBook(idVal, title, priceVal)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":  1,
				"msg":   "edit failed",
				"error": err.Error(),
			})
			return
		}

		c.Redirect(http.StatusMovedPermanently, "/book/list")
	} else { // GET
		bookObj, err := queryOneBook(idVal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  1,
				"msg":   "query failed",
				"error": err.Error(),
			})
			return
		}

		c.HTML(http.StatusOK, "book/book_edit.html", bookObj)
	}
}

func deleteBookHandler(c *gin.Context) {
	idStr := c.Query("id")
	idVal, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "id is invalid",
			"error": err.Error(),
		})
		return
	}

	fmt.Println("idVal:", idVal)
	err = deleteBook(idVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  1,
			"msg":   "delete failed",
			"error": err.Error(),
		})
		return
	}

	// 删除成功，跳转到书籍列表页
	c.Redirect(http.StatusMovedPermanently, "/book/list")
}

func main() {
	// 程序启动就连接数据库
	err := initDB()
	if err != nil {
		panic(err)
	}

	// gin 默认使用了 log, recover 中间件
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/book/new", newBookHandler)
	r.POST("/book/new", createBookHandler)

	r.GET("/book/list", bookListHandler)
	// 组合查询
	r.GET("/book/:id", bookDetailHandler)

	// Any
	r.Any("book/edit", editBookHandler)

	r.GET("/book/delete", deleteBookHandler)

	r.Run()
}
