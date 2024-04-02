package entities

import (
	"time"

	"captioner.com.ng/internal/captioner/core/dto"
)

type (
	Session struct {
		ID           interface{} `json:"id" bson:"_id"`
		Email        string      `json:"email" bson:"email"`
		RefreshToken string      `json:"refresh_token" bson:"refresh_token"`
		ExpiresAt    time.Time   `json:"expires_at" bson:"expires_at"`
		IP_Address   string      `json:"ip_address" bson:"ip_address"`
		User_Agent   string      `json:"user_agent" bson:"user_agent"`
		CreatedAt    time.Time   `json:"created_at" bson:"created_at"`
		UpdatedAt    time.Time   `json:"updated_at" bson:"updated_at"`
	}
)

func NewSession(s dto.CreateSessionRequest) Session {
	return Session{
		Email:        s.Email,
		RefreshToken: s.RefreshToken,
		ExpiresAt:    time.Now().AddDate(0, 2, 0),
		IP_Address:   s.IP_Address,
		User_Agent:   s.User_Agent,
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}
}
