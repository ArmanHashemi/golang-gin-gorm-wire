package usecase

import (
	"application/src/model"
	"context"
	"encoding/json"
	"time"

	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	jwt "github.com/dgrijalva/jwt-go"
)

type key int

const (
	ClientAddress key = iota
	RequestID     key = iota
)

func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

type ContextKey struct {
	ctx context.Context
}

func ContextKeyLogger(ctx context.Context) *ContextKey {
	return &ContextKey{ctx: ctx}
}

func (c *ContextKey) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	client := c.ctx.Value(ClientAddress)
	if client == nil {
		client = ""
	}
	requestID := c.ctx.Value(RequestID)
	if requestID == nil {
		requestID = ""
	}

	enc.AddString("Client", client.(string))
	enc.AddString("Request-ID", requestID.(string))
	return nil
}

func CtxLogger(ctx context.Context, logger *zap.Logger) *zap.Logger {
	return logger.With(zap.Object("ctx", ContextKeyLogger(ctx)))
}

var jwtSecret = []byte("your-secret-key")

func GenerateToken(user model.AuthUser) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &model.AuthUser{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Id,
			ExpiresAt: expirationTime.Unix(),
		},
		Username: user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*model.AuthUser, error) {
	claims := &model.AuthUser{}

	jwtCode := tokenString[7:]

	token, err := jwt.ParseWithClaims(jwtCode, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
