package main

import (
	"context"
	"fmt"

	"demo/database"
	"demo/models"

	// both can be used
	// "gopkg.in/mgo.v2/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertOneUser() {
	userinfo := models.User{
		Id:       1,
		Name:     "york",
		Password: "123456",
		Age:      18,
		Email:    "york@gmail.com",
	}

	result, err := database.QmgoConnection.InsertOne(context.TODO(), userinfo)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func InsertUsers() {
	userinfo := []models.User{
		{
			Id:       2,
			Name:     "fish",
			Password: "123456",
			Age:      18,
			Email:    "fish@gmail.com",
		},
		{
			Id:       3,
			Name:     "jessy",
			Password: "123456",
			Age:      18,
			Email:    "jessy@gmail.com",
		},
		{
			Id:       4,
			Name:     "alice",
			Password: "123456",
			Age:      18,
			Email:    "alice@gmail.com",
		},
	}

	result, err := database.QmgoConnection.InsertMany(context.TODO(), userinfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", result)
}

func FindOne() {
	user := models.User{}
	database.QmgoConnection.Find(context.TODO(), bson.M{"name": "york"}).One(&user)
	fmt.Printf("%+v\n", user)
}

func FindAll() {
	users := []models.User{}

	database.QmgoConnection.Find(context.TODO(), bson.M{}).All(&users)
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}
}

func Aggregate() {
	matchStage := bson.D{{"$match", []bson.E{{"age", bson.D{{"$gt", 10}}}}}}
	groupStage := bson.D{{"$group", bson.D{
		{"_id", "$name"},
		{"email", bson.D{{"$push", "$email"}}},
		{"age", bson.D{{"$sum", "$age"}}},
	}}}

	var showEithInfo []bson.M
	err := database.QmgoConnection.Aggregate(context.TODO(), []bson.D{matchStage, groupStage}).All(&showEithInfo)
	if err != nil {
		panic(err)
	}

	for _, info := range showEithInfo {
		fmt.Printf("%+v\n", info)
	}
}

func main() {
	// InsertOneUser()
	// InsertUsers()

	// FindOne()
	// FindAll()

	Aggregate()
}
