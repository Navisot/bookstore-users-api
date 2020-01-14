package user_model

import (
	"fmt"
	"github.com/navisot/bookstore-users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func GetDemoUsers(userId int64) *User {
	return usersDB[userId]
}

func SetDemoUsers(user *User) {
	usersDB[user.Id] = user
}


func (user *User) Get() *errors.RestErr {

	result := GetDemoUsers(user.Id)

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {

	currentUser := GetDemoUsers(user.Id)

	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s is already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	SetDemoUsers(user)
	return nil
}


