package usecase

import (
	"context"

	"captioner.com.ng/internal/captioner/adapters/store"
	"captioner.com.ng/internal/captioner/core/dto"
	"captioner.com.ng/internal/captioner/core/entities"
	"captioner.com.ng/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IUser interface {
		GetUser(ctx context.Context, request []constants.Identifier) (*dto.UserResponse, int, error)
		GetUsers(ctx context.Context, request []constants.Identifier) (*[]dto.UserResponse, int, error)
		CreateUser(ctx context.Context, request dto.CreateUserRequest) (*dto.UserResponse, int, error)
		UpdateUserPassword(ctx context.Context, request dto.UpdateUserPasswordRequest) (*dto.UserResponse, int, error)
		UpdateUserStatus(ctx context.Context, request dto.UpdateUserStatusRequest) (*dto.UserResponse, int, error)
		DeleteUser(ctx context.Context, request constants.Identifier) (int, error)
	}

	UserUsecase struct {
		store *store.Store[dto.UserResponse]
	}
)

var _ IUser = (*UserUsecase)(nil)

func InitUserUsecase(client *mongo.Database) *UserUsecase {
	return &UserUsecase{
		store: store.InitStore[dto.UserResponse](client.Collection("users")),
	}
}

func (u *UserUsecase) GetUser(ctx context.Context, request []constants.Identifier) (*dto.UserResponse, int, error) {

	result, err := u.store.Get(ctx, request)
	if err != nil {
		return nil, 400, err
	}

	return result, 200, nil
}

func (u *UserUsecase) GetUsers(ctx context.Context, request []constants.Identifier) (*[]dto.UserResponse, int, error) {
	result, err := u.store.GetAll(ctx, request)
	if err != nil {
		return nil, 400, err
	}

	return result, 200, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, request dto.CreateUserRequest) (*dto.UserResponse, int, error) {
	user := entities.NewUser(request)

	result, err := u.store.Create(ctx, user)
	if err != nil {
		return nil, 400, err
	}

	response := dto.UserResponse{
		ID:        result.(primitive.ObjectID).Hex(),
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &response, 201, nil
}

func (u *UserUsecase) UpdateUserPassword(ctx context.Context, request dto.UpdateUserPasswordRequest) (*dto.UserResponse, int, error) {
	newPassword := entities.NewUserPassword(request)

	user_id, _ := primitive.ObjectIDFromHex(request.ID)

	filter := []constants.Identifier{{
		Key:   "_id",
		Value: user_id,
	}}

	body := []constants.Identifier{{
		Key:   "password",
		Value: newPassword.Password,
	}, {
		Key:   "updated_at",
		Value: newPassword.UpdatedAt,
	}}

	result, err := u.store.Update(ctx, filter, body)

	if err != nil {
		return nil, 400, err
	}

	return result, 200, nil
}

func (u *UserUsecase) UpdateUserStatus(ctx context.Context, request dto.UpdateUserStatusRequest) (*dto.UserResponse, int, error) {
	newStatus := entities.NewUserStatus(request)

	user_id, _ := primitive.ObjectIDFromHex(request.ID)

	filter := []constants.Identifier{{
		Key:   "_id",
		Value: user_id,
	}}

	body := []constants.Identifier{{
		Key:   "account_status",
		Value: newStatus.AccountStatus,
	}, {
		Key:   "updated_at",
		Value: newStatus.UpdatedAt,
	}}

	result, err := u.store.Update(ctx, filter, body)

	if err != nil {
		return nil, 400, err
	}

	return result, 200, nil
}

func (u *UserUsecase) DeleteUser(ctx context.Context, request constants.Identifier) (int, error) {

	user_id, _ := primitive.ObjectIDFromHex(request.Value.(string))

	filter := []constants.Identifier{{
		Key:   "_id",
		Value: user_id,
	}}

	err := u.store.Delete(ctx, filter)

	if err != nil {
		return 400, err
	}

	return 204, nil
}
