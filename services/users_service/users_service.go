package users_service

import (
	"github.com/navisot/bookstore-users-api/domains/user_model"
	"github.com/navisot/bookstore-users-api/utils/errors"
)

func GetUser(userId int64) (*user_model.User, *errors.RestErr) {

	getUser := user_model.User{Id:userId}

	if err := getUser.Get(); err != nil {
		return nil, err
	}

	return &getUser, nil
}

func CreateUser(user user_model.User) (*user_model.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil

}
