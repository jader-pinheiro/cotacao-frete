package xfx

import "go.uber.org/fx"

func ProvideAs[T any](inst any) fx.Option {
	return fx.Provide(fx.Annotate(inst, fx.As(new(T))))
}
