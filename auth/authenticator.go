package auth

import (
	"context"

	"github.com/mohidex/identity-service/models"
)

type Authenticator interface {
	GenerateToken(ctx context.Context, user *models.User) (string, error)
	VerifyToken(ctx context.Context, token string) (*models.RequestUser, error)
}
