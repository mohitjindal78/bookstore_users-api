//Package users define user domain
//dao stand for data access object
package users

import (
	"fmt"

	"github.com/mohitjindal78/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
	"github.com/mohitjindal78/bookstore_users-api/utils/mysql_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name = ?, last_name = ?, email = ?, status = ?, password = ? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE status=?;"
)

var (
	userDB = make(map[int64]*User)
)

//Get function get the user data
func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return mysql_utils.ParseError("error while preparing get user", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	//scan will scan the result and populate pointer to given fields
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		return mysql_utils.ParseError("error while sca get user", err)
	}

	return nil
}

//Save function save the data
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return mysql_utils.ParseError("error while preparing save user", err)
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		return mysql_utils.ParseError("error while exec save user", err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError("error while retrieving lastinsert id user", err)
	}
	user.Id = userId

	return nil
}

//Update function save the data
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return mysql_utils.ParseError("error while preparing update user", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.Id)
	if err != nil {
		return mysql_utils.ParseError("error while exec update user", err)
	}
	return nil
}

//Delete function get the user data
func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return mysql_utils.ParseError("error while preparing delete user", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError("error while exec delete user", err)
	}
	return nil
}

//FindByStatus function get the user data
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, mysql_utils.ParseError("error while preparing findbyuser user", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError("error while query findbyuser user", err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			return nil, mysql_utils.ParseError("error while scan findbyuser user", err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users found with given status %s", status))
	}
	return results, nil
}
