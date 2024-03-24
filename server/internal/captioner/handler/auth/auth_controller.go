package auth

import (
	"net/http"

	"captioner.com.ng/api/types"
	"captioner.com.ng/internal/captioner/handler/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	client   *mongo.Client
	service  *AuthService
	validate *validator.Validate
}

func NewAuthController(client *mongo.Client) *AuthController {
	return &AuthController{
		client:   client,
		service:  NewAuthService(client),
		validate: validator.New(),
	}
}

func (a *AuthController) Login(c *gin.Context) {

	var body Signin
	c.ShouldBindJSON(&body)
	if err := a.validate.Struct(body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	body.IP = c.ClientIP()
	body.Device = c.GetHeader("user-agent")

	users, sessions, err := a.service.Login(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, types.HttpResponse[string]{Status: types.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
		return
	}
	c.SetCookie("x-refresh", sessions.RefreshToken, types.REFRESH_TOKEN_EXPIRE, "*", "*", true, true)
	c.SetCookie("x-session", sessions.SessionToken, types.SESSION_TOKEN_EXPIRE, "*", "*", true, true)
	c.JSON(http.StatusOK, types.HttpResponse[user.User]{Status: types.SUCCESS, StatusCode: http.StatusOK, Message: "Users Successfully Retrieved", Data: *users})
}
