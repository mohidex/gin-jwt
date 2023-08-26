package auth

import "context"

type Authenticator interface {
	GenerateToken(ctx context.Context, userID uint) (string, error)
	VerifyToken(ctx context.Context, token string) (uint, error)
}
