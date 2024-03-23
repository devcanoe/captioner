package user

import (
	"captioner.com.ng/api/server/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserRoute(g *gin.Engine, database *mongo.Client) {

	u := NewUserController(database)

	r := g.Group("/users")

	r.Group("/").GET("/", u.GetUsers).POST("/", u.CreateUser)
	r.Use(middleware.IsAuthenticated())
	r.Group("/:id").GET("/", u.GetUser).PATCH("/", u.UpdateUser).DELETE("/", u.DeleteUser)

}
