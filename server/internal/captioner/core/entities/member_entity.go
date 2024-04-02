package entities

import (
	"time"

	"captioner.com.ng/internal/captioner/core/dto"
)

const ()

type (
	Member struct {
		ID          interface{}
		UserID      interface{}
		Permissions []string
		WorkspaceID interface{}
		TeamID      interface{}
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	UpdateWorkspaceMember struct {
		WorkspaceID  interface{}
		PermissionID interface{}
		UpdatedAt    time.Time
	}

	UpdateTeamMember struct {
		TeamID       interface{}
		PermissionID interface{}
		UpdatedAt    time.Time
	}
)

func NewWorkspaceMember(w dto.CreateWorkspaceMemberRequest) Member {
	return Member{
		UserID:      w.UserID,
		WorkspaceID: w.WorkspaceID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func NewTeamMember(w dto.CreateTeamMemberRequest) Member {
	return Member{
		UserID:    w.UserID,
		TeamID:    w.TeamID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func NewUpdateWorkspaceMember(w dto.UpdateWorkspaceMemberRequest) Member {
	return Member{
		WorkspaceID: w.WorkspaceID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func NewUpdateTeamMember(w dto.UpdateTeamMemberRequest) Member {
	return Member{
		TeamID:    w.TeamID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
