package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func (jm *JWTAuthenticator) GenerateToken(ctx context.Context, userID uint) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	expirationTime := jm.clock().Add(time.Duration(jm.ttlMinutes) * time.Minute)

	claims := jwt.MapClaims{
		"sub": userID,
		"iat": jm.clock().Unix(),
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jm.secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (jm *JWTAuthenticator) VerifyToken(ctx context.Context, token string) (uint, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}

	claims, err := jm.verifyToken(token)
	if err != nil {
		return 0, err
	}
	userID, ok := claims["sub"].(float64)
	if !ok {
		return 0, ErrInvalidToken
	}

	return uint(userID), nil
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
