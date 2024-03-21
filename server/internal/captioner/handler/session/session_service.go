package session

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SessionService struct {
	client   *mongo.Client
	validate *validator.Validate
	repo     *SessionRepository
}

func NewSessionService(client *mongo.Client) *SessionService {
	return &SessionService{
		client:   client,
		validate: validator.New(),
		repo:     NewSessionRepository(client),
	}
}

func (s *SessionService) GetSessions() (*[]Session, error) {
	sessions, err := s.repo.GetAllSession()

	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (s *SessionService) GetSession(id string) (*Session, error) {
	sessionID, fail := primitive.ObjectIDFromHex(id)

	if fail != nil {
		return nil, fail
	}

	session, err := s.repo.GetOneSession(sessionID)

	if err != nil {
		return nil, err
	}
	return session, nil

}

func (s *SessionService) CreateSession(params CreateSession) (*Session, error) {
	v := s.validate

	invalid := v.Struct(params)
	if invalid != nil {
		return nil, invalid
	}
	var newParams = CreateSession{
		UserID:       params.UserID,
		RefreshToken: "string",
		SessionToken: "string",
		Device:       params.Device,
		IP:           params.IP,
	}

	session, err := s.repo.CreateOneSession(newParams)

	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) UpdateSession(id string, params UpdateSession) (*Session, error) {
	sessionID, fail := primitive.ObjectIDFromHex(id)
	if fail != nil {
		return nil, fail
	}
	v := s.validate

	invalid := v.Struct(params)
	if invalid != nil {
		return nil, invalid
	}

	session, err := s.repo.UpdateOneSession(sessionID, params)

	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *SessionService) DeleteSession(id string) error {
	sessionID, fail := primitive.ObjectIDFromHex(id)

	if fail != nil {
		return fail
	}

	err := s.repo.DeleteOneSession(sessionID)

	if err != nil {
		return err
	}
	return nil
}
