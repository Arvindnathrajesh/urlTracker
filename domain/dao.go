package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"../utils"

	"github.com/monaco-io/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"userId":   user.UserId,
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	fmt.Println(result)
	fmt.Println(err)
	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

// Params:  map[string]string{"hello": "world"},

func crateUrlShortner(url string) string {
	var result interface{}

	client := request.Client{
		URL:    "https://api-ssl.bitly.com/v4/bitlinks",
		Method: "POST",
		JSON:   bson.M{"long_url": url},
		Bearer: "38419ca28f8f4a6495303228b148c45c58937c89",
	}
	if err := client.Send().Scan(&result).Error(); err != nil {
		return "Create URL failed"
	}

	str := client.Send().String()
	x := map[string]string{}

	json.Unmarshal([]byte(str), &x)
	fmt.Println(x)
	return x["id"]
}

func CreateLinkData(linkData *LinkData) (*LinkData, *utils.RestErr) {
	usersC := db.Collection("LinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	shortUrl := crateUrlShortner(linkData.Url)
	result, err := usersC.InsertOne(ctx, bson.M{
		"url":      linkData.Url,
		"shortUrl": shortUrl,
		"useCase":  linkData.UseCase,
		"userId":   linkData.UserId,
	})
	fmt.Println(err)
	if err != nil {
		restErr := utils.InternalErr("can't insert linkData to the database.")
		return nil, restErr
	}
	linkData.ID = result.InsertedID.(primitive.ObjectID)
	return linkData, nil
}

// func CreateUserLinkData(url string, userPhone string) (*UserLinkData, *utils.RestErr) {
// 	userLinkDatasC := db.Collection("UserLinkData")
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
// 	result, err := userLinkDatasC.InsertOne(ctx, bson.M{
// 		"url":        url,
// 		"userPhone":  userPhone,
// 		"clickCount": 1,
// 	})
// 	if err != nil {
// 		restErr := utils.InternalErr("can't insert userLinkData to the database.")
// 		return nil, restErr
// 	}
// 	var userLinkData UserLinkData
// 	errFind := userLinkDatasC.FindOne(ctx, bson.M{"url": url, "userPhone": userPhone}).Decode(&userLinkData)
// 	if errFind != nil {
// 		restErr := utils.NotFound("userLinkData not found.")
// 		return nil, restErr
// 	}
// 	userLinkData.ID = result.InsertedID.(primitive.ObjectID)
// 	return &userLinkData, nil
// }

func Find(email string) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := usersC.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	return &user, nil
}

// func UrlClicked(url string, userPhone string) (*UserLinkData, *utils.RestErr) {
// 	var userLinkData UserLinkData
// 	userLinkDatasC := db.Collection("UserLinkData")

// 	fmt.Println(url)
// 	fmt.Println(userPhone)
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
// 	err := userLinkDatasC.FindOne(ctx, bson.M{"url": url, "userPhone": userPhone}).Decode(&userLinkData)

// 	fmt.Println(userLinkData)
// 	if err != nil {
// 		CreateUserLinkData(url, userPhone)
// 	} else {
// 		UpdateUserLinkData(&userLinkData)
// 	}
// 	return &userLinkData, nil
// }

func FindLinkData(shortUrl string) (*LinkData, *utils.RestErr) {
	var linkData LinkData
	linkDatasC := db.Collection("LinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := linkDatasC.FindOne(ctx, bson.M{"shortUrl": shortUrl}).Decode(&linkData)
	if err != nil {
		restErr := utils.NotFound("LinkData not found.")
		return nil, restErr
	}
	// fmt.Println(shortUrl)
	// fmt.Println(linkData)
	return &linkData, nil
}

// func FindUserLinkData(url string, userPhone string) (*UserLinkData, *utils.RestErr) {
// 	var userLinkData UserLinkData
// 	userLinkDatasC := db.Collection("UserLinkData")
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
// 	err := userLinkDatasC.FindOne(ctx, bson.M{"url": url, "userPhone": userPhone}).Decode(&userLinkData)
// 	if err != nil {
// 		restErr := utils.NotFound("UserLinkData not found.")
// 		return nil, restErr
// 	}
// 	return &userLinkData, nil
// }

// func UpdateUserLinkData(userLinkData *UserLinkData) (*UserLinkData, *utils.RestErr) {
// 	UserLinkDatasC := db.Collection("UserLinkData")
// 	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
// 	result, err := UserLinkDatasC.UpdateOne(ctx, bson.M{"url": userLinkData.Url, "userPhone": userLinkData.UserPhone}, bson.M{"$set": bson.M{"clickCount": userLinkData.ClickCount + 1}})
// 	if err != nil {
// 		restErr := utils.InternalErr("can not update.")
// 		return nil, restErr
// 	}
// 	if result.MatchedCount == 0 {
// 		restErr := utils.NotFound("userLinkData not found.")
// 		return nil, restErr
// 	}
// 	if result.ModifiedCount == 0 {
// 		restErr := utils.BadRequest("no such field")
// 		return nil, restErr
// 	}
// 	userLinkData, restErr := FindUserLinkData(userLinkData.Url, userLinkData.UserPhone)
// 	if restErr != nil {
// 		return nil, restErr
// 	}
// 	return userLinkData, restErr
// }

func Update(email string, field string, value string) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	result, err := usersC.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{field: value}})
	if err != nil {
		restErr := utils.InternalErr("can not update.")
		return nil, restErr
	}
	if result.MatchedCount == 0 {
		restErr := utils.NotFound("user not found.")
		return nil, restErr
	}
	if result.ModifiedCount == 0 {
		restErr := utils.BadRequest("no such field")
		return nil, restErr
	}
	user, restErr := Find(email)
	if restErr != nil {
		return nil, restErr
	}
	return user, restErr
}
