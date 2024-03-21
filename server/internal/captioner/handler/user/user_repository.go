package user

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DATABASE   = "captioner"
	COLLECTION = "users"
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

type IUser interface {
	GetOneUser(id primitive.ObjectID) (*User, error)
	GetAllUsers() (*[]User, error)
	CreateUser(user CreateUser) (*User, error)
	UpdateUser(id primitive.ObjectID) (*User, error)
	DeleteUser(id primitive.ObjectID) error
}

type UserRepository struct {
	client *mongo.Collection
	User   User
}

func NewUserRepository(client *mongo.Client) *UserRepository {
	return &UserRepository{
		client: client.Database(DATABASE).Collection(COLLECTION),
	}
}

func (u *UserRepository) GetOneUser(id primitive.ObjectID) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := u.client
	var user *User
	defer cancel()

	err := m.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u *UserRepository) GetAllUsers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := u.client
	var users []User

	defer cancel()

	result, err := m.Find(ctx, bson.M{})
	if err != nil {
		return users, err
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var singleUser User
		err := result.Decode(&singleUser)
		if err != nil {
			return nil, err
		}
		users = append(users, singleUser)
	}

	return users, nil
}

func (u *UserRepository) CreateUser(user CreateUser) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := u.client
	var newUser *User
	defer cancel()

	newUser = &User{
		ID:        primitive.NewObjectID(),
		Email:     user.Email,
		Username:  fmt.Sprintf("%s%d", strings.Split(string(user.Email), "@")[0], rand.Intn(10000)),
		Password:  user.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err := m.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u *UserRepository) UpdateUser(id primitive.ObjectID, user UpdateUser) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := u.client
	var returnUser *User
	defer cancel()

	userParam := bson.M{
		"password":   user.Password,
		"updated_at": time.Now().UTC(),
	}
	result := m.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": userParam}, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err := result.Decode(&returnUser)
	if err != nil {
		return nil, err
	}

	return returnUser, nil
}

func (u *UserRepository) DeleteUser(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	m := u.client
	defer cancel()

	result, err := m.DeleteOne(ctx, bson.M{"_id": id})
	if result.DeletedCount < 1 {
		return errors.New("UserID not found")
	}
	if err != nil {
		return err

	}
	return nil
}
