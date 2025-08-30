package pojo

import (
	db "demo/database"
	"log"
)

type User struct {
	Id       int    `json:"UserId"`   // UserId => user_id
	Name     string `json:"UserName"` // Name => name, UserName => user_name
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

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
	return result.RowsAffected != 0
}

// Update
func UpdateUser(userId string, user User) User {
	db.DBConnect.Model(&user).Where("id = ?", userId).Updates(user)
	return user
}
