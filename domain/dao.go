package domain

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"urlTracker/utils"

	"github.com/monaco-io/request"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link_Click_Struct struct {
	Date   string
	Clicks int
}
type Response struct {
	Link_Clicks []Link_Click_Struct
}

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
	result, err := usersC.InsertOne(ctx, bson.M{
		"userId":   user.UserId,
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		restErr := utils.InternalErr("can't insert user to the database.")
		return nil, restErr
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	user.Password = ""
	return user, nil
}

func crateUrlShortner(url string) string {
	var result interface{}
	token := "38419ca28f8f4a6495303228b148c45c58937c89"

	client := request.Client{
		URL:    "https://api-ssl.bitly.com/v4/bitlinks",
		Method: "POST",
		JSON:   bson.M{"long_url": url},
		Bearer: token,
	}
	if err := client.Send().Scan(&result).Error(); err != nil {
		log.Fatal(err)
	}

	str := client.Send().String()
	x := map[string]string{}

	json.Unmarshal([]byte(str), &x)
	return x["id"]
}

func getClicksBitlink(shortUrl string) *[]Link_Click_Struct {

	token := "38419ca28f8f4a6495303228b148c45c58937c89"
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-ssl.bitly.com/v4/bitlinks/"+shortUrl+"/clicks?unit=minute&units=-1", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var x Response
	json.Unmarshal(bodyText, &x)
	return &x.Link_Clicks
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
	if err != nil {
		restErr := utils.InternalErr("can't insert linkData to the database.")
		return nil, restErr
	}
	linkData.ID = result.InsertedID.(primitive.ObjectID)
	return linkData, nil
}

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

func UrlClicked(shortUrl string) (bool, *utils.RestErr) {

	LinkLogC := db.Collection("LinkLog")
	resp := getClicksBitlink(shortUrl)

	for i := 0; i < len(*resp); i++ {

		for j := 0; j < (*resp)[i].Clicks; j++ {
			ctx, _ := context.WithTimeout(context.Background(), time.Second*20)
			result, err := LinkLogC.InsertOne(ctx, bson.M{
				"logId":    randstr.Hex(16),
				"shortUrl": shortUrl,
				"date":     (*resp)[i].Date,
			})
			if err != nil {
				restErr := utils.InternalErr("can't insert linkData to the database.")
				return false, restErr
			}
			fmt.Println(result)
		}
	}

	return true, nil
}

func FindLinkData(shortUrl string) (*LinkData, *utils.RestErr) {
	var linkData LinkData
	linkDatasC := db.Collection("LinkData")
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err := linkDatasC.FindOne(ctx, bson.M{"shortUrl": shortUrl}).Decode(&linkData)
	if err != nil {
		restErr := utils.NotFound("LinkData not found.")
		return nil, restErr
	}
	return &linkData, nil
}

func FindLinksData() (*[]LinkData, *utils.RestErr) {
	var results []LinkData
	linkDatasC := db.Collection("LinkData")
	cursor, err := linkDatasC.Find(context.TODO(), bson.D{})
	if err != nil {
		restErr := utils.NotFound("LinkData not found.")
		return nil, restErr
	}
	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem LinkData
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	cursor.Close(context.TODO())
	return &results, nil
}

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
