package dto

import "time"

type (
	CreateWorkspaceRequest struct {
		Name string `json:"name" bson:"name"`
	}
	UpdateWorkspaceRequest struct {
		Name string `json:"name"`
	}

	WorkspaceResponse struct {
		ID        interface{} `json:"id"`
		Name      string      `json:"name"`
		CreatedAt time.Time   `json:"created_at"`
		UpdatedAt time.Time   `json:"updated_at"`
	}
)
