package dto

import "time"

type (
	CreateWorkspaceMemberRequest struct {
		UserID      interface{}
		Permission  []string
		WorkspaceID interface{}
	}
	CreateTeamMemberRequest struct {
		UserID     interface{}
		Permission []string
		TeamID     interface{}
	}
	UpdateWorkspaceMemberRequest struct {
		Permission  []string
		WorkspaceID interface{}
	}
	UpdateTeamMemberRequest struct {
		Permission []string
		TeamID     interface{}
	}

	MemberResponse struct {
		ID          interface{}
		UserID      interface{}
		Permission  []string
		WorkspaceID interface{}
		TeamID      interface{}
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
