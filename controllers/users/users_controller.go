package users

import (
	//"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mohitjindal78/bookstore_users-api/domain/users"

	//"io/ioutil"
	"net/http"
	//"encoding/json"
	"strconv"

	"github.com/mohitjindal78/bookstore_users-api/services"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
)

//getUserId func get user id
func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, errors.NewBadRequestError("use id should be a number")
	}
	return userId, nil
}

//Create function creates
func Create(c *gin.Context) {
	var user users.User
	//c.ShouldBindJSON do all of the below commented code
	//fmt.Println("user: ", user)
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//TODO: Need to handle Error
	//fmt.Println("Err :", err.Error())
	//return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//TODO Need to handle json Error
	//fmt.Println("Err :", err.Error())
	//return
	//}
	//fmt.Println("Response :", string(bytes))
	//fmt.Println("user: ", user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

//Get function get user
func Get(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

//Update func updates
func Update(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError(err.Error())
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	result, updateErr := services.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

//Delete func updates
func Delete(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}

	if deleteErr := services.DeleteUser(userId); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
