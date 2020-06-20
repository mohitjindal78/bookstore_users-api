//Package users define user domain
//dao stand for data access object
package users

import (
	"github.com/mohitjindal78/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mohitjindal78/bookstore_users-api/utils/date_utils"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
	"github.com/mohitjindal78/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name = ?, last_name = ?, email = ? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

var (
	userDB = make(map[int64]*User)
)

//Get function get the user data
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	//scan will scan the result and populate pointer to given fields
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

//Save function save the data
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId

	return nil
}

//Update function save the data
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

//Delete function get the user data
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
