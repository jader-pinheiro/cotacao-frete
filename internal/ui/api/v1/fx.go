package v1

import (
	"cotacao-fretes/internal/core/quote"
	"cotacao-fretes/internal/infra/db/mysql"
	"cotacao-fretes/pkg/fiberfx"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(

		mysql.Module(),
		quote.Module(),

		fx.Provide(NewQuoteController),
		fx.Provide(NewAuthController),

		fiberfx.Mount(NewV1),
	)
}
