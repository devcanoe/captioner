package user

import "captioner.com.ng/internal/captioner/types"

type UserHandler struct {
	User *types.User
}

func (u *UserHandler) GetOneUser(id string) (*types.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandler) GetAllUsers() (*[]types.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandler) CreateUser(params types.User) error {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandler) UpdateUser(id string, params types.User) (*types.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *UserHandler) DeleteUser(id string) error {
	panic("not implemented") // TODO: Implement
}
