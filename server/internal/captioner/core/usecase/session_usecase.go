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
	ISession interface {
		GetSessions(ctx context.Context, request constants.Identifier) (*[]entities.Session, int, error)
		CreateSession(ctx context.Context, request dto.CreateSessionRequest) (int, error)
		DeleteSession(ctx context.Context, request constants.Identifier) (int, error)
	}

	SessionUsecase struct {
	}
)

var _ ISession = (*SessionUsecase)(nil)

func InitSessionUsecase(db *mongo.Database, v *validator.Validate) *SessionUsecase {
	return &SessionUsecase{}
}

func (s *SessionUsecase) GetSessions(ctx context.Context, request constants.Identifier) (*[]entities.Session, int, error) {
	panic("not implemented") // TODO: Implement

}

func (s *SessionUsecase) CreateSession(ctx context.Context, request dto.CreateSessionRequest) (int, error) {

	panic("not implemented") // TODO: Implement
}

func (s *SessionUsecase) DeleteSession(ctx context.Context, request constants.Identifier) (int, error) {

	panic("not implemented") // TODO: Implement
}
