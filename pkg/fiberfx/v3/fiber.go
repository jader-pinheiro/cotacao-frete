package fiberfx

import "github.com/gofiber/fiber/v3"

type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"3000"`
}

func New() (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	return app, nil
}
