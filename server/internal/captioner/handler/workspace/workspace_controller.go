package workspace

import (
	"net/http"

	response "captioner.com.ng/internal/captioner/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkspaceController struct {
	client *mongo.Client
}

func NewWorkspaceController(client *mongo.Client) *WorkspaceController {
	return &WorkspaceController{
		client: client,
	}
}

func (w *WorkspaceController) GetWorkspace(c *gin.Context) {
	id := c.Param("id")
	workspace, err := NewWorkspaceService(w.client).GetOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
	}
	c.JSON(http.StatusOK, response.HttpResponse[Workspace]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Workspace Successfully Retrieved!", Data: *workspace})
}

func (w *WorkspaceController) GetWorkspaces(c *gin.Context) {
	workspaces, err := NewWorkspaceService(w.client).GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
	}
	c.JSON(http.StatusOK, response.HttpResponse[[]Workspace]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Workspace Successfully Retrieved!", Data: *workspaces})
}

func (w *WorkspaceController) CreateWorkspace(c *gin.Context) {
	var body CreateWorkspace
	c.ShouldBindJSON(&body)

	workspace, err := NewWorkspaceService(w.client).CreateOne(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
	}
	c.JSON(http.StatusOK, response.HttpResponse[Workspace]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Workspace Successfully Retrieved!", Data: *workspace})
}

func (w *WorkspaceController) UpdateWorkspace(c *gin.Context) {
	id := c.Param("id")
	var body UpdateWorkspace
	c.ShouldBindJSON(&body)
	workspace, err := NewWorkspaceService(w.client).UpdateOne(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
	}
	c.JSON(http.StatusOK, response.HttpResponse[Workspace]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Workspace Successfully Retrieved!", Data: *workspace})
}

func (w *WorkspaceController) DeleteWorkspace(c *gin.Context) {
	id := c.Param("id")

	err := NewWorkspaceService(w.client).DeleteOne(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: err.Error(), Data: ""})
	}
	c.JSON(http.StatusNoContent, response.HttpResponse[string]{Status: response.SUCCESS, StatusCode: http.StatusNoContent, Message: "Workspace Successfully Retrieved!", Data: ""})
}
