package dto

type (
	CreateSessionRequest struct {
		Email        string `json:"email" verified:"required"`
		RefreshToken string `verified:"required"`
		IP_Address   string `verified:"required"`
		User_Agent   string `verified:"required"`
	}
)
