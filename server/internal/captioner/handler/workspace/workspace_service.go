package workspace

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkspaceService struct {
	client    *mongo.Client
	Workspace Workspace
	validate  *validator.Validate
}

func NewWorkspaceService(client *mongo.Client) *WorkspaceService {
	return &WorkspaceService{
		client:   client,
		validate: validator.New(),
	}
}

func (w *WorkspaceService) GetOne(id string) (*Workspace, error) {
	workspaceID, _ := primitive.ObjectIDFromHex(id)

	workspace, err := NewWorkspaceRepsository(w.client).GetOneWorkpace(workspaceID)

	if err != nil {
		return nil, err
	}
	return workspace, nil
}

func (w *WorkspaceService) GetAll() (*[]Workspace, error) {
	workspaces, err := NewWorkspaceRepsository(w.client).GetAllWorkpaces()
	if err != nil {
		return nil, err
	}
	return workspaces, nil
}

func (w *WorkspaceService) CreateOne(params CreateWorkspace) (*Workspace, error) {
	v := w.validate
	fail := v.Struct(params)
	if fail != nil {
		return nil, fail
	}

	result, err := NewWorkspaceRepsository(w.client).CreateWorkspace(params)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (w *WorkspaceService) UpdateOne(id string, params UpdateWorkspace) (*Workspace, error) {
	v := w.validate
	fail := v.Struct(params)
	workspaceID, _ := primitive.ObjectIDFromHex(id)
	if fail != nil {
		return nil, fail
	}
	result, err := NewWorkspaceRepsository(w.client).UpdateWorkspace(workspaceID, params)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (w *WorkspaceService) DeleteOne(id string) error {
	workspaceID, _ := primitive.ObjectIDFromHex(id)

	err := NewWorkspaceRepsository(w.client).DeleteWorkspace(workspaceID)

	if err != nil {
		return err
	}
	return nil
}
