package session

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DATABASE   = "captioner"
	COLLECTION = "sessions"
)

type SessionRepository struct {
	client  *mongo.Collection
	Session Session
}

func NewSessionRepository(s *mongo.Client) *SessionRepository {
	return &SessionRepository{
		client: s.Database(DATABASE).Collection(COLLECTION),
	}
}

type ISession interface {
	GetAllSession() (*[]Session, error)
	GetOneSession(id primitive.ObjectID) (*Session, error)
	CreateOneSession(params CreateSession) (*Session, error)
	UpdateOneSession(id primitive.ObjectID, params UpdateSession) (*Session, error)
	DeleteOneSession(id primitive.ObjectID) error
}

func (s *SessionRepository) GetAllSession() (*[]Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := s.client
	var sessions []Session
	defer cancel()

	result, err := m.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var singleSession Session
		result.Decode(&singleSession)
		sessions = append(sessions, singleSession)
	}
	return &sessions, nil
}

func (s *SessionRepository) GetOneSession(id primitive.ObjectID) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := s.client
	var session Session
	defer cancel()

	result := m.FindOne(ctx, bson.M{"_id": id})

	err := result.Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionRepository) CreateOneSession(params CreateSession) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := s.client
	var newSession Session
	defer cancel()

	newSession = Session{
		ID:           primitive.NewObjectID(),
		UserID:       params.UserID,
		RefreshToken: params.RefreshToken,
		SessionToken: params.SessionToken,
		IsActive:     true,
		ExpiresAt:    time.Now().UTC(),
		IP:           params.IP,
		Device:       params.Device,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
	_, err := m.InsertOne(ctx, newSession)

	if err != nil {
		return nil, err
	}

	return &newSession, nil
}

func (s *SessionRepository) UpdateOneSession(id primitive.ObjectID, params UpdateSession) (*Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := s.client
	var session Session
	defer cancel()
	var newSession = bson.M{
		"session_token": params.SessionToken,
		"is_active":     params.IsActive,
		"updated_at":    time.Now().UTC(),
	}
	result := m.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": newSession})
	err := result.Decode(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (s *SessionRepository) DeleteOneSession(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := s.client
	defer cancel()

	result, err := m.DeleteOne(ctx, bson.M{"_id": id})

	if result.DeletedCount > 1 {
		return errors.New("user not found")
	}

	if err != nil {
		return err
	}
	return nil
}
