package configfx

import (
	"github.com/caarlos0/env/v11"
	"go.uber.org/fx"
)

type Opt func(o *env.Options)

func Prefix(prefix string) Opt {
	return func(o *env.Options) {
		o.Prefix = prefix + "_"
	}
}

func Module[T any](opts ...Opt) fx.Option {
	opt := env.Options{}
	for _, o := range opts {
		o(&opt)
	}

	return fx.Options(
		fx.Provide(func() (T, error) {
			return env.ParseAsWithOptions[T](opt)
		}),
	)
}
