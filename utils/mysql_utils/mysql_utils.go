package mysql_utils

import (
	"strings"

	"github.com/mohitjindal78/bookstore_users-api/logger"

	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	errorNoRows      = "no rows in result set"
)

//ParseError function handle mysql error.
func ParseError(msg string, err error) *errors.RestErr {
	if strings.Contains(err.Error(), errorNoRows) {
		return errors.NewInternalServerError("user not found")
	}

	if strings.Contains(err.Error(), indexUniqueEmail) {
		return errors.NewBadRequestError("email id already exists")
	}

	logger.Error(msg, err)
	return errors.NewInternalServerError("database error")
}
