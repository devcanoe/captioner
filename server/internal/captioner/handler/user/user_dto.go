package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Password  string             `json:"password,omitempty" bson:"password"`
	Username  string             `json:"username,omitempty" bson:"username"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type CreateUser struct {
	Email    string `validate:"email,required"`
	Password string `validate:"required"`
}

type UpdateUser struct {
	Password  string `validate:"required"`
	UpdatedAt time.Time
}
