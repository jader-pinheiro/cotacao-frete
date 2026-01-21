package jwtauth

import (
	"cotacao-fretes/pkg/configfx"

	"go.uber.org/fx"
)

func Module(name string) fx.Option {
	return fx.Options(
		configfx.Module[Config](configfx.Prefix(name)),
		fx.Provide(New),
	)
}
