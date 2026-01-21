package v1

import (
	"context"
	"fmt"
	"log/slog"

	"cotacao-fretes/internal/infra/jwtauth"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	slog *slog.Logger
	aut  jwtauth.Auth
}

func NewAuthController(lg *slog.Logger, aut *jwtauth.Auth) *AuthController {
	return &AuthController{
		lg,
		*aut,
	}
}

type AuthRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Claims struct {
	ClientID string `json:"client_id"`
	jwt.RegisteredClaims
}

func (ac *AuthController) Authenticate(c *fiber.Ctx) error {
	var req AuthRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Erro ao processar a requisição",
		})
	}

	if req.ClientID != ac.aut.CotacaoClientID || req.ClientSecret != ac.aut.CotacaoClientSecret {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Erro: credenciais inválidas",
		})
	}

	token, err := ac.aut.GenerateJWT(req.ClientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Erro ao gerar o token: %v", err),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (ac *AuthController) JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	var ctx context.Context

	if authHeader == "" {
		ac.slog.ErrorContext(ctx, "Acesso negado: token de autorização não fornecido")

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "token de autorização não fornecido",
		})
	}

	claims, err := ac.aut.ValidateJWT(authHeader)
	if err != nil {
		msgFull := fmt.Sprintf("Acesso negado: %s", err.Error())
		ac.slog.ErrorContext(ctx, msgFull)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fmt.Sprintf("Acesso negado: %v", err),
		})
	}
	c.Locals("claims", claims)

	return c.Next()
}
