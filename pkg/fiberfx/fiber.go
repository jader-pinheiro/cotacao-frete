package fiberfx

import "github.com/gofiber/fiber/v2"

type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"3000"`
}

func New() (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		Prefork:               false,
		Immutable:             true,
		DisableStartupMessage: true,
	})

	return app, nil
}
