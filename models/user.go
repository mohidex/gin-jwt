package models

import (
	"html"
	"strings"

	"github.com/mohidex/identity-service/settings"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null;" json:"-"`
	PhotoURL string `gorm:"size:255" json:"photo_url"`
	Active   bool   `gorm:"not null;default:true" json:"is_active"`
	Admin    bool   `gorm:"not null;default:false" json:"is_admin"`
}

// type UserManager struct {
// 	db   *gorm.DB
// 	user *User
// }

func (user *User) Save() (*User, error) {
	db := settings.GetDB()
	if result := db.Create(&user); result.Error != nil {
		return &User{}, result.Error
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User

	if result := settings.DB.Where("username=?", username).Find(&user); result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	if err := settings.DB.Preload("Blogs").Where("ID=?", id).Find(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
