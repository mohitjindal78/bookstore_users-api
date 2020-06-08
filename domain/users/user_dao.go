//Package users define user domain
//dao stand for data access object
package users

import (
	"fmt"

	"github.com/mohitjindal78/bookstore_users-api/utils/date_utils"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

//Get function get the user data
func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
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

//Save function save the data
func (user *User) Save() *errors.RestErr {
	currentUser := userDB[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()

	userDB[user.Id] = user
	return nil
}
