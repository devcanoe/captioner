package entities

import (
	"time"

	"captioner.com.ng/internal/captioner/core/dto"
)

type (
	Workspace struct {
		ID        interface{} `json:"id" bson:"_id"`
		Name      string      `json:"name" bson:"name"`
		CreatedAt time.Time   `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time   `json:"updated_at" bson:"updated_at"`
	}

	UpdateWorkspaceName struct {
		Name      string `json:"name" validate:"required"`
		UpdatedAt time.Time
	}
)

func NewWorkspace(w *dto.CreateWorkspaceRequest) Workspace {
	return Workspace{
		Name:      w.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func NewUpdateName(w *dto.UpdateWorkspaceRequest) UpdateWorkspaceName {
	return UpdateWorkspaceName{
		Name:      w.Name,
		UpdatedAt: time.Now().UTC(),
	}
}
