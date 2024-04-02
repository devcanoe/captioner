package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"captioner.com.ng/internal/captioner/core/dto"
	"captioner.com.ng/internal/captioner/core/usecase"
	"captioner.com.ng/pkg/constants"
	"captioner.com.ng/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthHandler struct {
		user *usecase.UserUsecase
		v    *validator.Validate
	}

	IAuthHandler interface {
		Signin(c *gin.Context)
		VerifyEmail(c *gin.Context)
		Signup(c *gin.Context)
		ForgotPassword(c *gin.Context)
		ResetPassword(c *gin.Context)
	}
)

var _ IAuthHandler = (*AuthHandler)(nil)

func NewAuthHandler(db *mongo.Database, v *validator.Validate) *AuthHandler {
	return &AuthHandler{
		user: usecase.InitUserUsecase(db),
		v:    v,
	}
}

func (a *AuthHandler) Signin(c *gin.Context) {
	//GET JSON CONTENT
	var body dto.SigninUserRequest
	c.ShouldBindJSON(&body)

	//VALIDATE JSON CONTENT
	if err := a.v.Struct(body); err != nil {
		c.AbortWithError(400, err)
		return
	}

	//GET USER
	param := []constants.Identifier{
		{Key: "email", Value: body.Email},
	}

	result, status, err := a.user.GetUser(c.Request.Context(), param)

	//RESULT USER
	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	c.JSON(status, result)

}

func (a *AuthHandler) VerifyEmail(c *gin.Context) {
	//GET JSON CONTENT
	var body dto.VerifyEmailRequest
	c.ShouldBindJSON(&body)

	//VALIDATE JSON CONTENT
	if err := a.v.Struct(body); err != nil {
		c.AbortWithError(400, err)
		return
	}

	//GET USER
	param := []constants.Identifier{
		{Key: "email", Value: body.Email},
	}

	_, status, err := a.user.GetUser(c.Request.Context(), param)

	if err == mongo.ErrNoDocuments {
		payload := utils.JWTPayload{
			Email: body.Email,
			Type:  utils.EMAIL_TOKEN,
		}

		token, err := utils.CreateToken(payload, time.Now().Add(utils.EMAIL_TOKEN_EXPIRE*time.Second))
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		c.JSON(200, token)
		return
	}
	if err != nil {
		c.AbortWithError(status, err)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "user already exist"})
}

func (a *AuthHandler) Signup(c *gin.Context) {
	//GET JSON BODY
	var body dto.CreateUserRequest
	c.ShouldBindJSON(&body)
	if err := a.v.Struct(body); err != nil {
		c.AbortWithError(400, err)
		return
	}

	//GET TOKEN AND VALIDATE
	token, err := utils.VerifyToken(c.Query("token"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	token_type, _ := token.Get("type")

	if token_type.(string) != utils.EMAIL_TOKEN {
		c.AbortWithError(400, errors.New("invalid token type"))
		return
	}

	//CREATE USER
	result, status, err := a.user.CreateUser(c.Request.Context(), body)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(status, result)
}

func (a *AuthHandler) ForgotPassword(c *gin.Context) {
	//GET JSON CONTENT
	var body dto.VerifyEmailRequest
	c.ShouldBindJSON(&body)

	//VALIDATE JSON CONTENT
	if err := a.v.Struct(body); err != nil {
		c.AbortWithError(400, err)
		return
	}

	//GET USER
	param := []constants.Identifier{
		{Key: "email", Value: body.Email},
	}

	_, status, err := a.user.GetUser(c.Request.Context(), param)

	if err != nil {
		c.AbortWithError(status, err)
		return
	}

	//CREATE TOKEN
	payload := utils.JWTPayload{
		Email: body.Email,
		Type:  utils.RESET_TOKEN,
	}

	token, err := utils.CreateToken(payload, time.Now().Add(utils.RESET_TOKEN_EXPIRE*time.Second))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.JSON(200, token)
}

func (a *AuthHandler) ResetPassword(c *gin.Context) {
	//GET JSON
	var body dto.UpdateUserPasswordRequest
	c.ShouldBindJSON(&body)

	//VERIFY TOKEN
	token, err := utils.VerifyToken(c.Query("token"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	fmt.Println(token)

	//VERIFY JSON
	if err := a.v.Struct(body); err != nil {
		c.AbortWithError(400, err)
		return
	}

	//GET EMAIL & TOKEN TYPE FROM TOKEN
	email, _ := token.Get("email")
	token_type, _ := token.Get("type")

	if token_type.(string) == utils.RESET_TOKEN {
		//UPDATE USER WITH TOKEN
		filter := []constants.Identifier{{Key: "email", Value: email}}
		user, status, err := a.user.GetUser(c.Request.Context(), filter)
		if err != nil {
			c.AbortWithError(status, err)
			return
		}
		fmt.Println(user)

		body.ID = user.ID

		result, status, err := a.user.UpdateUserPassword(c.Request.Context(), body)
		if err != nil {
			c.AbortWithError(status, err)
			return
		}
		fmt.Println(result)

		c.JSON(status, result)
		return
	}

	c.JSON(400, gin.H{"message": "token is invalid"})
}
