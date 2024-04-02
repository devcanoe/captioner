package dto

import "time"

type (
	CreateUserRequest struct {
		Email    string `validate:"email,required"`
		Password string `validate:"required"`
	}
	SigninUserRequest struct {
		Email    string `validate:"email,required"`
		Password string `validate:"required"`
	}
	UpdateUserPasswordRequest struct {
		ID       string `json:"id" bson:"_id"`
		Password string `validate:"required"`
	}
	UpdateUserStatusRequest struct {
		ID            string `json:"id" bson:"_id"`
		AccountStatus string `json:"account_status" validate:"required"`
	}
	VerifyEmailRequest struct {
		Email string `json:"email" validate:"email,required"`
	}

	UserResponse struct {
		ID            string    `json:"id" bson:"_id"`
		Email         string    `json:"email,omitempty" bson:"email"`
		Role          string    `json:"role,omitempty" bson:"role"`
		AccountStatus string    `json:"account_status,omitempty" bson:"account_status"`
		CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	}
)
