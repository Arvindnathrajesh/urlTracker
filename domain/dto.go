package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId   int                `json:"userId" bson:"userId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
	Phone    string             `json:"phone" bson:"phone,omitempty"`
	Password string             `json:"password" bson:"password,omitempty"`
}

type LinkData struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Url      string             `json:"url" bson:"url,omitempty"`
	ShortUrl string             `json:"shortUrl" bson:"shortUrl,omitempty"`
	UseCase  string             `json:"useCase" bson:"useCase,omitempty"`
	UserId   int                `json:"userId" bson:"userId,omitempty"`
}

type LinkLog struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LogId    string             `json:"logId" bson:"logId,omitempty"`
	ShortUrl string             `json:"shortUrl" bson:"shortUrl,omitempty"`
	Date     string             `json:"date" bson:"date,omitempty"`
}

type LinkCityLog struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LinkLogId string             `json:"linkLogId" bson:"linkLogId,omitempty"`
	ShortUrl  string             `json:"shortUrl" bson:"shortUrl,omitempty"`
	City      string             `json:"city" bson:"city,omitempty"`
	Region    string             `json:"region" bson:"region,omitempty"`
	Country   string             `json:"country" bson:"country,omitempty"`
}
