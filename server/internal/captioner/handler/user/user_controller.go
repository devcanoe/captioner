package user

import (
	"net/http"

	response "captioner.com.ng/internal/captioner/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{
		client: client,
	}
}

func (u *UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := NewUserService(u.client).GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusNotFound, Message: "User Not Found!", Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[User]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "User Successfully Retrieved!", Data: *user})
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := NewUserService(u.client).GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: "Could't get Users", Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[[]User]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Users Successfully Retrieved", Data: users})

}

func (u *UserController) CreateUser(c *gin.Context) {
	var userBody CreateUser
	c.ShouldBindJSON(&userBody)

	user, err := NewUserService(u.client).CreateUser(userBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[User]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "User Successfully Created", Data: *user})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var userBody UpdateUser
	c.ShouldBindJSON(&userBody)
	id := c.Param("id")

	user, err := NewUserService(u.client).UpdateUser(id, userBody)
	if err != nil {
		c.JSON(http.StatusNotFound, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusNotFound, Message: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusAccepted, response.HttpResponse[User]{Status: response.SUCCESS, StatusCode: http.StatusAccepted, Message: "User Successfully Updated", Data: *user})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := NewUserService(u.client).DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
		return
	}
	c.JSON(http.StatusNoContent, response.HttpResponse[string]{Status: response.SUCCESS, StatusCode: http.StatusNoContent, Message: "Users Successfully Created", Data: ""})
}
