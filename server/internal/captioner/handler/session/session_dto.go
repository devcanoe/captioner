package session

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`
	RefreshToken string             `json:"refresh_token" bson:"refresh_token"`
	SessionToken string             `json:"session_token" bson:"session_token"`
	ExpiresAt    time.Time          `json:"expires_at" bson:"expires_at"`
	IsActive     bool               `json:"is_active" bson:"is_active"`
	IP           string             `json:"ip" bson:"ip"`
	Device       string             `json:"device" bson:"device"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type CreateSession struct {
	UserID       primitive.ObjectID `json:"user_id" verified:"required"`
	RefreshToken string             `verified:"required"`
	SessionToken string             `verified:"required"`
	IP           string             `verified:"required"`
	Device       string             `verified:"required"`
}

type UpdateSession struct {
	SessionToken string `verified:"required"`
	IsActive     bool
	UpdatedAt    time.Time
}
