//Package users define user domain
//dao stand for data access object
package users

import (
	"fmt"
	"strings"

	"github.com/mohitjindal78/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mohitjindal78/bookstore_users-api/utils/date_utils"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

var (
	userDB = make(map[int64]*User)
)

//Get function get the user data
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	//scan will scan the result and populate pointer to given fields
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewInternalServerError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

//Save function save the data
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email id, %s, already exists", user.Email))
		}
		return errors.NewInternalServerError(err.Error())
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	user.Id = userId

	return nil
}
