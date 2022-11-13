package services

import (
	"../domain"
	"../utils"
)

func CreateUser(user *domain.User) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func FindUser(email string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Find(email)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}

func UpdateUser(email string, field string, value string) (*domain.User, *utils.RestErr) {
	user, restErr := domain.Update(email, field, value)
	if restErr != nil {
		return nil, restErr
	}
	user.Password = ""
	return user, nil
}
