package router

import (
	"captioner.com.ng/internal/captioner/adapters/app/handlers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Router struct {
	router *gin.RouterGroup
	user   *handlers.MeHandler
	auth   *handlers.AuthHandler
}

func NewRouterGroup(r *gin.Engine, c *mongo.Database, v *validator.Validate) *Router {
	return &Router{
		router: &r.RouterGroup,
		user:   handlers.NewMeHandler(c, v),
		auth:   handlers.NewAuthHandler(c, v),
	}
}
