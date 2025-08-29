package service

import (
	"log"
	"net/http"
	"strconv"

	"demo/pojo"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

// Get User
func FindAllUsers(c *gin.Context) {
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

// Get User by Id
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user)
}

// Post User
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	userList = append(userList, user)
	c.JSON(http.StatusOK, "Successfully posted")
}

// Delete User
func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	for i, user := range userList {
		log.Println(user)
		if user.Id == userId {
			userList = append(userList[:i], userList[i+1:]...)
			c.JSON(http.StatusOK, "Successfully deleted")
			return
		}
	}

	c.JSON(http.StatusNotFound, "Error")
}

// Put User
func PutUser(c *gin.Context) {
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
	}

	userId, _ := strconv.Atoi(c.Param("id"))
	for i, user := range userList {
		if user.Id == userId {
			userList[i] = beforeUser
			c.JSON(http.StatusOK, "Successfully updated")
			return
		}
	}

	c.JSON(http.StatusNotFound, "Error")
}
