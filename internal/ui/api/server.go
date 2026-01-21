package api

import (
	"cotacao-fretes/internal/infra/jwtauth"
	"cotacao-fretes/internal/infra/requesthttp"
	v1routes "cotacao-fretes/internal/ui/api/v1"
	"cotacao-fretes/pkg/fiberfx"
	"cotacao-fretes/pkg/probes"
	"cotacao-fretes/pkg/slogfx"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		slogfx.Module(),
		jwtauth.Module("QUOTE_SERVICE"),
		requesthttp.Module(),

		fiberfx.Module(),
		probes.HTTP(),

		v1routes.Module(),
	)
}
