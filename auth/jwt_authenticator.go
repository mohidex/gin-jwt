package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mohidex/identity-service/models"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type JWTAuthenticator struct {
	secretKey  []byte
	ttlMinutes int
	clock      func() time.Time
}

func NewJWTAuthenticator(secretKey string, ttlMinutes int) *JWTAuthenticator {
	return &JWTAuthenticator{
		secretKey:  []byte(secretKey),
		ttlMinutes: ttlMinutes,
		clock:      time.Now,
	}
}

func (jm *JWTAuthenticator) GenerateToken(ctx context.Context, user *models.User) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	expirationTime := jm.clock().Add(time.Duration(jm.ttlMinutes) * time.Minute)

	claims := jwt.MapClaims{
		"sub":      user.ID,
		"iat":      jm.clock().Unix(),
		"exp":      expirationTime.Unix(),
		"username": user.Username,
		"email":    user.Email,
		"is_admin": user.Admin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (jm *JWTAuthenticator) VerifyToken(ctx context.Context, token string) (*models.RequestUser, error) {
	if ctx.Err() != nil {
		return &models.RequestUser{}, ctx.Err()
	}

	claims, err := jm.verifyToken(token)
	if err != nil {
		return &models.RequestUser{}, err
	}
	userID, ok := claims["sub"].(float64)
	if !ok {
		return &models.RequestUser{}, ErrInvalidToken
	}

	username, _ := claims["username"].(string)
	isAdmin, _ := claims["is_admin"].(bool)
	email, _ := claims["email"].(string)

	requestUser := &models.RequestUser{
		ID:       uint(userID),
		Username: username,
		Email:    email,
		Admin:    isAdmin,
	}

	return requestUser, nil
}

func (jm *JWTAuthenticator) verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jm.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
