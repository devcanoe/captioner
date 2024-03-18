package types

import "time"

var Addr string

type User struct {
	ID        uint32    `json:"_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IUser interface {
	GetOneUser(id string) (*User, error)
	GetAllUsers() (*[]User, error)
	CreateUser(params User) error
	UpdateUser(id string, params User) (*User, error)
	DeleteUser(id string) error
}
