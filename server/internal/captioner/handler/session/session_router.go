package session

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SessionRoute(g *gin.Engine, database *mongo.Client) {
	s := NewSessionController(database)

	r := g.Group("/sessions")

	r.Group("/").GET("/", s.GetAll).POST("/", s.CreateOne)
	r.Group("/:id").GET("/", s.GetOne).PATCH("/", s.UpdateOne).DELETE("/", s.DeleteOne)
}
