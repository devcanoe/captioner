package handlers

import (
	"captioner.com.ng/internal/captioner/core/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	MeHandler struct {
		user *usecase.UserUsecase
	}

	IMeHandler interface {
	}
)

func NewMeHandler(db *mongo.Database, v *validator.Validate) *MeHandler {
	return &MeHandler{
		user: usecase.InitUserUsecase(db),
	}

}

func (u *MeHandler) HealthCheck(c *gin.Context) {
	c.JSON(200, "hello")
}
