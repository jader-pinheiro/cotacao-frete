package v1

import (
	_ "cotacao-fretes/docs" // Importa a documentação gerada
	"cotacao-fretes/internal/pkg/scalar"
	"cotacao-fretes/internal/pkg/scalar/confs"
	"fmt"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
)

// nolint
func NewV1(
	quote *QuoteController,
	auth *AuthController,
) *fiber.App {
	app := fiber.New()

	app.Use(fiberzap.New())

	v1 := app.Group("/v1")

	v1.Post("/auth/token", auth.Authenticate)

	v1.Get("/quote/metrics", quote.GetResumeQuotes)
	v1.Post("/quote", quote.Insert)

	app.Get("/docs/*", func(c *fiber.Ctx) error {
		htmlContent, err := scalar.ApiReferenceHTML(&confs.Options{
			SpecURL: "./docs/swagger.json",
			CustomOptions: confs.CustomOptions{
				PageTitle: "Cotação de Fretes - Documentação API",
				// LogoURL:   "https://cdn.scalar.com/photos/mars.jpg", // (Opcional) Adicione seu logo
			},
			DarkMode: true,
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Erro ao renderizar documentação Scalar: %v", err))
		}

		c.Set("Content-Type", "text/html")
		return c.SendString(htmlContent)
	})

	return app
}
