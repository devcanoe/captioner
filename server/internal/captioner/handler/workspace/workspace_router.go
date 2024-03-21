package workspace

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func WorkspaceRoute(g *gin.Engine, database *mongo.Client) {
	w := NewWorkspaceController(database)

	r := g.Group("/workspaces")

	r.Group("/:id").GET("/", w.GetWorkspace)
}
