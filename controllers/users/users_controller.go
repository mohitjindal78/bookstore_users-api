package users

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mohitjindal78/bookstore_users-api/domain/users"
	//"io/ioutil"
	"net/http"
	//"encoding/json"
	"github.com/mohitjindal78/bookstore_users-api/services"
	"github.com/mohitjindal78/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
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

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "\nImplement Me!")
}
