package probes

import (
	"go.uber.org/fx"
)

func HTTP() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotate(
				New,
				fx.ResultTags(`group:"subapp"`),
			),
		),
	)
}

func WithReady(r any) fx.Option {
	return fx.Provide(fx.Annotate(r, fx.ResultTags(`group:"readiness"`)))
}
