package session

import (
	"net/http"

	response "captioner.com.ng/internal/captioner/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionController struct {
	client  *mongo.Client
	service *SessionService
}

func NewSessionController(client *mongo.Client) *SessionController {
	return &SessionController{
		client:  client,
		service: NewSessionService(client),
	}
}
func (s *SessionController) GetAll(c *gin.Context) {
	sessions, err := s.service.GetSessions()

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: "Could't get Users", Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[[]Session]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Users Successfully Retrieved", Data: *sessions})

}

func (s *SessionController) GetOne(c *gin.Context) {
	id := c.Param("id")

	session, err := s.service.GetSession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: "Could't get Users", Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[Session]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Users Successfully Retrieved", Data: *session})

}

func (s *SessionController) CreateOne(c *gin.Context) {
	var body CreateSession
	c.ShouldBindJSON(&body)

	body.Device = c.GetHeader("user-agent")
	body.IP = c.ClientIP()

	session, err := s.service.CreateSession(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: "Could't get Users", Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[Session]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Users Successfully Retrieved", Data: *session})

}

func (s *SessionController) UpdateOne(c *gin.Context) {
	id := c.Param("id")
	var body UpdateSession
	c.ShouldBindJSON(&body)

	session, err := s.service.UpdateSession(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: "Could't get Users", Data: ""})
		return
	}
	c.JSON(http.StatusOK, response.HttpResponse[Session]{Status: response.SUCCESS, StatusCode: http.StatusOK, Message: "Users Successfully Retrieved", Data: *session})
}

func (s *SessionController) DeleteOne(c *gin.Context) {
	id := c.Param("id")

	err := s.service.DeleteSession(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.HttpResponse[string]{Status: response.ERROR, StatusCode: http.StatusInternalServerError, Message: "Could't get Users", Data: ""})
		return
	}
	c.JSON(http.StatusNoContent, "")
}
