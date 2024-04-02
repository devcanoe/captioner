package entities

import (
	"time"

	"captioner.com.ng/internal/captioner/core/dto"
	"golang.org/x/crypto/bcrypt"
)

const (
	ACTIVE  = "ACTIVE"
	BLOCKED = "BLOCKED"
	BANNED  = "BANNED"
	DELETED = "DELETED"

	//ROLE
	DEFAULT_ROLE = "USER"
	ADMIN        = "ADMIN"
)

type (
	User struct {
		Email         string    `json:"email,omitempty" bson:"email"`
		Password      string    `json:"password,omitempty" bson:"password"`
		Role          string    `json:"role,omitempty" bson:"role"`
		AccountStatus string    `json:"account_status,omitempty" bson:"account_status"`
		CreatedAt     time.Time `json:"created_at,omitempty" bson:"created_at"`
		UpdatedAt     time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	}

	UpdatePassword struct {
		Password  string `validate:"required"`
		UpdatedAt time.Time
	}

	UpdateStatus struct {
		AccountStatus string `json:"account_status" validate:"required"`
		UpdatedAt     time.Time
	}
)

func NewUserPassword(u dto.UpdateUserPasswordRequest) UpdatePassword {
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	return UpdatePassword{
		Password:  string(hasedPassword),
		UpdatedAt: time.Now().UTC(),
	}
}

func NewUserStatus(u dto.UpdateUserStatusRequest) UpdateStatus {
	return UpdateStatus{
		AccountStatus: u.AccountStatus,
		UpdatedAt:     time.Now().UTC(),
	}
}

func NewUser(u dto.CreateUserRequest) User {
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	return User{
		Email:         u.Email,
		Password:      string(hasedPassword),
		Role:          DEFAULT_ROLE,
		AccountStatus: ACTIVE,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
}

func NewAdmin(u dto.CreateUserRequest) User {
	hasedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	return User{
		Email:         u.Email,
		Password:      string(hasedPassword),
		Role:          ADMIN,
		AccountStatus: ACTIVE,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
}
