package router

import (
	"captioner.com.ng/internal/captioner/handler/user"
	"captioner.com.ng/internal/captioner/handler/workspace"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Connect(g *gin.Engine, database *mongo.Client) {
	user.UserRoute(g, database)
	workspace.WorkspaceRoute(g, database)
}
