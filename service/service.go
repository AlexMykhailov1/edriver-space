package service

import (
	"bytes"
	"context"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/model"
	"github.com/ASeegull/edriver-space/pkg/auth"
	"github.com/ASeegull/edriver-space/pkg/hash"
	"github.com/ASeegull/edriver-space/repository"
)

type Auth interface {
	SignIn(ctx context.Context, user UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, sessionId string) (Tokens, error)
	GetUserById(ctx context.Context, userId string) (*model.User, error)
	DeleteSession(ctx context.Context, sessionId string) error
}

type Uploader interface {
	XMLFinesService(ctx context.Context, data model.Data) error
	ReadFinesExcel(ctx context.Context, r *bytes.Reader) error
}

type Services struct {
	Auth     Auth
	Uploader Uploader
}

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager, hasher hash.PasswordHasher, cfg *config.Config) *Services {
	return &Services{
		Auth:     NewAuthService(repos, tokenManager, hasher, cfg),
		Uploader: NewUploadService(repos, cfg),
	}
}
