package dto

import "time"

type (
	CreateTeamRequest struct {
		Title       string
		WorkspaceID interface{}
	}
	UpdateTeamTitleRequest struct {
		Title string
	}

	TeamResponse struct {
		ID          interface{}
		Title       string
		WorkspaceID string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
