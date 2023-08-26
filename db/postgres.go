package db

import (
	"context"

	"github.com/mohidex/identity-service/models"
	"gorm.io/gorm"
)

type PgDB struct {
	db *gorm.DB
}

func NewPgDB(db *gorm.DB) *PgDB {
	return &PgDB{db: db}
}

func (d *PgDB) SaveUser(ctx context.Context, user *models.User) (*models.User, error) {
	db := d.db.WithContext(ctx)

	// Save the user using GORM's Create method
	result := db.Create(user)
	if result.Error != nil {
		return &models.User{}, result.Error
	}

	return user, nil
}

func (d *PgDB) GetAllUsers(ctx context.Context) ([]models.User, error) {
	db := d.db.WithContext(ctx)

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (d *PgDB) DeleteUser(ctx context.Context, id uint) error {
	db := d.db.WithContext(ctx)

	if result := db.Delete(&models.User{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *PgDB) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	db := d.db.WithContext(ctx)

	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return &models.User{}, err
	}
	return &user, nil
}

func (d *PgDB) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	db := d.db.WithContext(ctx)

	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
