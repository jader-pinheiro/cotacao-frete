package jwtauth

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func New(cfg Config) *Auth {
	return &Auth{
		CotacaoClientID:     cfg.CotacaoClientID,
		CotacaoClientSecret: cfg.CotacaoClientSecret,
	}
}

type Config struct {
	CotacaoClientID     string `env:"CLIENT_ID"`
	CotacaoClientSecret string `env:"CLIENT_SECRET"`
}

type Auth struct {
	CotacaoClientID     string
	CotacaoClientSecret string
}

// nolint
func (Auth) GenerateJWT(clientID string) (string, error) {
	claims := jwt.MapClaims{
		"client_id": clientID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("cotacao-frete"))
}

func (Auth) ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	var ctx context.Context

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.ErrorContext(ctx, "token com método de assinatura inválido")
			return nil, errors.New("token com método de assinatura inválido")
		}
		return []byte("cotacao-frete"), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	slog.ErrorContext(ctx, "Error Claims invalid")
	return nil, errors.New("error Claims invalid")
}
