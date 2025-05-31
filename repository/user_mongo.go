package repository

import (
	"context"
	"errors"

	"team-maker-api/domain"
	"team-maker-api/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type userMongoRepo struct{}

func NewUserRepository() domain.UserRepository {
	return &userMongoRepo{}
}

func (r *userMongoRepo) FindByUsernameAndPassword(username, password string) (*domain.User, error) {
	collection := utils.DB.Collection("users")

	var user domain.User
	err := collection.FindOne(context.TODO(), bson.M{"username": username, "password": password}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found or invalid credentials")
	}
	return &user, nil
}
