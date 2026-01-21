package probes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type Check func(ctx context.Context) error

type Params struct {
	fx.In
	Readiness []Check `group:"readiness"`
}

func New(p Params) *fiber.App {
	app := fiber.New()

	hc := app.Group("/health")

	hc.Get("/alive", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	hc.Get("/ready", func(ctx *fiber.Ctx) error {
		for _, check := range p.Readiness {
			if err := check(ctx.UserContext()); err != nil {
				return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
		}

		return ctx.SendStatus(fiber.StatusOK)
	})

	return app
}
