//Package services package will interact with all external system
package services

import (
	"github.com/mohitjindal78/bookstore_users-api/domain/users"
	"github.com/mohitjindal78/bookstore_users-api/utils/date_utils"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
)

//GetUser Func updates
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

//CreateUser Func updates
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

//UpdateUser Func updates
func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
		if user.Password != "" {
			current.Password = user.Password
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser Func updates
func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

//Search Func searches
func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
