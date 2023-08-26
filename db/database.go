package db

import (
	"context"

	"github.com/mohidex/identity-service/models"
)

type Database interface {
	SaveUser(ctx context.Context, user *models.User) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	DeleteUser(ctx context.Context, id uint) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
}
