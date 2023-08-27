package models

import (
	"html"
	"strings"

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

func (u *User) ToUserResponse() interface{} {
	minimalUser := struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		PhotoURL string `json:"photo_url"`
	}{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		PhotoURL: u.PhotoURL,
	}
	return minimalUser
}
