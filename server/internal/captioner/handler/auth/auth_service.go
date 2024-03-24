package auth

import (
	"time"

	"captioner.com.ng/api/server/utils"
	"captioner.com.ng/api/types"
	"captioner.com.ng/internal/captioner/handler/session"
	"captioner.com.ng/internal/captioner/handler/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	REFRESH_TOKEN = "refresh_token"
	SESSION_TOKEN = "session_token"
)

type AuthService struct {
	userRepo    *user.UserRepository
	sessionRepo *session.SessionRepository
}

func NewAuthService(client *mongo.Client) *AuthService {
	return &AuthService{
		userRepo:    user.NewUserRepository(client),
		sessionRepo: session.NewSessionRepository(client),
	}
}

type IAuthService interface {
	LoginUser(params Signin) (*user.User, *session.Session, error)
	RegisterUser(params Signup) (*user.User, error)
	VerifyUserEmail(token Token) error
}

func (a *AuthService) Login(params Signin) (*user.User, *session.Session, error) {
	var payload utils.JWTPayload
	var signinFilter = bson.M{
		"email": params.Email,
	}
	var sessionParams session.CreateSession

	user, err := a.userRepo.GetOneUser(signinFilter)

	if err != nil {
		return nil, nil, err
	}

	fail := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if fail != nil {
		return nil, nil, fail
	}

	payload = utils.JWTPayload{
		UserID: user.ID,
		Type:   REFRESH_TOKEN,
	}

	refresh_token, err := utils.CreateToken(payload, time.Now().Add(types.REFRESH_TOKEN_EXPIRE*time.Second))
	if err != nil {
		return nil, nil, err
	}
	payload.Type = SESSION_TOKEN
	session_token, err := utils.CreateToken(payload, time.Now().Add(types.SESSION_TOKEN_EXPIRE*time.Second))
	if err != nil {
		return nil, nil, err
	}
	sessionParams = session.CreateSession{
		UserID:       user.ID,
		RefreshToken: refresh_token,
		SessionToken: session_token,
		IP:           params.IP,
		Device:       params.Device,
	}

	session, invalid := a.sessionRepo.CreateOneSession(sessionParams)

	if invalid != nil {
		return nil, nil, invalid
	}

	return user, session, nil
}

// func (undefined) Register(params Signup) (*user.User, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (undefined) VerifyEmail(token Token) error {
// 	panic("not implemented") // TODO: Implement
// }
