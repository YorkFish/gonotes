package pojo

import (
	db "demo/database"
	"log"
	"strconv"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       int    `json:"UserId" binding:"required"` // UserId => user_id
	Name     string `json:"UserName" binding:"gt=5"`   // Name => name, UserName => user_name
	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"email"`
}

type Users struct {
	UserList     []User `json:"UserList" bingding:"required,gt=0,lt=3"`
	UserListSize int    `json:"UserListSize"`
}

// ===
// MySql

// Get
func FindAllUsers() []User {
	var users []User
	db.DBConnect.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User
	db.DBConnect.Where("id = ?", userId).First(&user)
	return user
}

// Post
func CreateUser(user User) User {
	db.DBConnect.Create(&user)
	return user
}

// Delete
func DeleteUser(userId string) bool {
	user := User{}
	result := db.DBConnect.Where("id = ?", userId).Delete(&user)
	log.Println(result)
	return result.RowsAffected > 0
}

// Update
func UpdateUser(userId string, user User) User {
	db.DBConnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}

// Check
func CheckUserPassword(name string, password string) User {
	user := User{}
	db.DBConnect.Where("name = ? and password = ?", name, password).First(&user)
	return user
}

// ===
// MongoDB

func MgoCreateUser(user User) User {
	db.MgoConnect.Insert(user)
	return user
}

func MgoFindAllUser() []User {
	users := []User{}
	db.MgoConnect.Find(nil).All(&users)
	return users
}

func MgoFindById(id string) User {
	userId, _ := strconv.Atoi(id)
	user := User{}
	db.MgoConnect.Find(bson.M{"id": userId}).One(&user)
	return user
}

func MgoPutUser(id string, user User) User {
	userId, _ := strconv.Atoi(id)
	updateUserId := bson.M{"id": userId}
	updateData := bson.M{"$set": user}
	err := db.MgoConnect.Update(updateUserId, updateData)
	if err != nil {
		log.Println(err)
		return User{}
	}
	return user
}

func MgoDeleteUser(id string) bool {
	userId, _ := strconv.Atoi(id)
	err := db.MgoConnect.Remove(bson.M{"id": userId})
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
