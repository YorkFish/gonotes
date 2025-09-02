package service

import (
	"log"
	"net/http"

	db "demo/database"
	"demo/middlewares"
	"demo/pojo"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

// ===
// MySql

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

	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

// Delete User
func DeleteUser(c *gin.Context) {
	userDeleted := pojo.DeleteUser(c.Param("id"))
	if !userDeleted {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// Put User
func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}

	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}

// Create Userlist
func CreateUserList(c *gin.Context) {
	users := pojo.Users{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(400, "Error:%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// Login User
func LoginUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := pojo.CheckUserPassword(name, password)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}

	middlewares.SaveSession(c, user.Id)
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login Successfully",
		"User":     user,
		"Sessions": middlewares.GetSession(c),
	})
}

// Logout User
func LogoutUser(c *gin.Context) {
	middlewares.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Successfully",
	})
}

// check user session
func CheckUserSession(c *gin.Context) {
	sessionId := middlewares.GetSession(c)
	if sessionId == 0 {
		c.JSON(http.StatusUnauthorized, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Check Session Successfully",
		"User":    sessionId,
	})
}

// ===
// Redis

// redis one user
func RedisOneUser(c *gin.Context) {
	id := c.Param("id")
	if id == "0" {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	user := pojo.User{}
	db.DBConnect.Find(&user, id)
	c.Set("dbResult", user)
}

// redis all user
func RedisAllUser(c *gin.Context) {
	users := []pojo.User{}
	db.DBConnect.Find(&users)
	c.Set("dbUserAll", users)
}

// ===
// MongoDB

// MongoDB create user
func MongoDBCreateOneUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	newUser := pojo.MgoCreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "Create User Successfully",
		"User":    newUser,
	})
}

// MongoDB findall user
func MongoDBFindAllUser(c *gin.Context) {
	users := pojo.MgoFindAllUser()
	c.JSON(http.StatusOK, gin.H{
		"message": "Find All User Successfully",
		"User":    users,
	})
}

// MongoDB find user by id
func MongoDBFindOneUser(c *gin.Context) {
	user := pojo.MgoFindById(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Find User Successfully",
		"User":    user,
	})
}

// MongoDB put user
func MongoDBUpdateUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	user = pojo.MgoPutUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update User Successfully",
		"User":    user,
	})
}

// MongoDB delete user
func MongoDBDeleteUser(c *gin.Context) {
	user := pojo.MgoDeleteUser(c.Param("id"))
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}
