package entities

import (
	"time"

	"captioner.com.ng/internal/captioner/core/dto"
)

type (
	Team struct {
		ID          interface{}
		Title       string
		WorkspaceID interface{}
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	UpdateTeamTitle struct {
		Title     string
		UpdatedAt time.Time
	}
)

func NewTeam(t dto.CreateTeamRequest) Team {
	return Team{
		Title:       t.Title,
		WorkspaceID: t.WorkspaceID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func NewUpdateTeamTitle(t dto.UpdateTeamTitleRequest) UpdateTeamTitle {
	return UpdateTeamTitle{
		Title:     t.Title,
		UpdatedAt: time.Now().UTC(),
	}
}
