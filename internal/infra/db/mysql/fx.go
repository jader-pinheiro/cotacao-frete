package mysql

import (
	"cotacao-fretes/internal/core/quote"
	"cotacao-fretes/internal/domain"
	"cotacao-fretes/pkg/gormfx"
	"cotacao-fretes/pkg/gormfx/mysqlfx"
	"cotacao-fretes/pkg/xfx"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		gormfx.Module(),
		mysqlfx.Module(),
		gormfx.Migrate(
			&domain.Quote{},
			&domain.Dispatcher{},
			&domain.Offer{},
			&domain.CarrierInfo{},
			&domain.DeliveryTimeInfo{},
			&domain.WeightInfo{},
		),
		xfx.ProvideAs[quote.DBPort](New),
	)
}
