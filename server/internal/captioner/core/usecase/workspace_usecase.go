package usecase

import (
	"context"

	"captioner.com.ng/internal/captioner/core/dto"
	"captioner.com.ng/internal/captioner/core/entities"
	"captioner.com.ng/pkg/constants"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	IWorkspace interface {
		GetWorkspace(ctx context.Context, request constants.Identifier) (*entities.Workspace, error)
		GetWorkspaces(ctx context.Context, request constants.Identifier) (*[]entities.Workspace, error)
		CreateWorkspace(ctx context.Context, request dto.CreateWorkspaceRequest) (*entities.Workspace, error)
		UpdateWorkspaceName(ctx context.Context, request dto.UpdateWorkspaceRequest) (*entities.Workspace, error)
		DeleteWorkSpace(ctx context.Context, request constants.Identifier) error
	}

	WorkspaceUsecase struct {
	}
)

var _ IWorkspace = (*WorkspaceUsecase)(nil)

func InitWorkspaceUsecase(db *mongo.Database, v *validator.Validate) *WorkspaceUsecase {
	return &WorkspaceUsecase{}
}

func (w *WorkspaceUsecase) GetWorkspace(ctx context.Context, request constants.Identifier) (*entities.Workspace, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WorkspaceUsecase) GetWorkspaces(ctx context.Context, request constants.Identifier) (*[]entities.Workspace, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WorkspaceUsecase) CreateWorkspace(ctx context.Context, request dto.CreateWorkspaceRequest) (*entities.Workspace, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WorkspaceUsecase) UpdateWorkspaceName(ctx context.Context, request dto.UpdateWorkspaceRequest) (*entities.Workspace, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WorkspaceUsecase) DeleteWorkSpace(ctx context.Context, request constants.Identifier) error {
	panic("not implemented") // TODO: Implement
}
