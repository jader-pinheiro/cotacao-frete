package requesthttp

import (
	"go.uber.org/fx"

	"cotacao-fretes/pkg/configfx"
	"cotacao-fretes/pkg/httpfx"
)

func Module() fx.Option {
	return fx.Options(
		httpfx.Module(),
		configfx.Module[Config](),
		fx.Provide(New),
	)
}
