package user

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	client   *mongo.Client
	User     User
	validate *validator.Validate
}

func NewUserService(client *mongo.Client) *UserService {
	return &UserService{
		client:   client,
		validate: validator.New(),
	}
}

func (u *UserService) GetUser(id string) (*User, error) {
	userID, _ := primitive.ObjectIDFromHex(id)
	user, err := NewUserRepository(u.client).GetOneUser(userID)

	return user, err
}

func (u *UserService) GetUsers() ([]User, error) {
	user, err := NewUserRepository(u.client).GetAllUsers()

	return user, err
}

func (u *UserService) CreateUser(user CreateUser) (*User, error) {
	var createdUser *User
	fail := u.validate.Struct(user)
	if fail != nil {
		return nil, fail
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	newUser := CreateUser{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := NewUserRepository(u.client).CreateUser(newUser)
	return createdUser, err
}

func (u *UserService) UpdateUser(id string, user UpdateUser) (*User, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	userID, _ := primitive.ObjectIDFromHex(id)
	newUser := UpdateUser{
		Password: string(hashedPassword),
	}
	returnUser, err := NewUserRepository(u.client).UpdateUser(userID, newUser)

	return returnUser, err
}

func (u *UserService) DeleteUser(id string) error {
	userID, _ := primitive.ObjectIDFromHex(id)
	err := NewUserRepository(u.client).DeleteUser(userID)

	return err
}
