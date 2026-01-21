package slogfx

import (
	"cotacao-fretes/pkg/configfx"
	"log/slog"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func Module() fx.Option {
	return fx.Options(
		configfx.Module[Config](),
		fx.Provide(New),
		fx.WithLogger(func(l *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: l}
		}),
	)
}
