package gormfx

import (
	"cotacao-fretes/pkg/configfx"
	"cotacao-fretes/pkg/probes"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		configfx.Module[Config](),
		probes.WithReady(Check),
		fx.Provide(New),
	)
}
