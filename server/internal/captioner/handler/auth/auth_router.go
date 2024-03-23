package auth

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AuthRoute(g *gin.Engine, database *mongo.Client) {
	a := NewAuthController(database)

	r := g.Group("/auth")

	r.POST("/login", a.Login)
}
