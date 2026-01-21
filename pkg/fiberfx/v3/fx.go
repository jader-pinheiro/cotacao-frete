package fiberfx

import (
	"context"
	"cotacao-fretes/pkg/configfx"
	"net"

	"github.com/gofiber/fiber/v2"
	fibertrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gofiber/fiber.v2"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		configfx.Module[Config](),
		fx.Provide(New),
		fx.Invoke(lifecycle),
	)
}

type params struct {
	fx.In
	LC      fx.Lifecycle
	App     *fiber.App
	Cfg     Config
	SubApps []*fiber.App `group:"subapp"`
}

func lifecycle(p params) {
	p.App.Use(fibertrace.Middleware())

	for _, subApp := range p.SubApps {
		p.App.Use("/", subApp)
	}

	p.LC.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			done := make(chan error)
			p.App.Hooks().OnListen(func(_ fiber.ListenData) error {
				done <- nil
				return nil
			})

			go func() {
				done <- p.App.Listen(net.JoinHostPort(p.Cfg.Host, p.Cfg.Port))
			}()
			return <-done
		},
		OnStop: func(ctx context.Context) error {
			return p.App.ShutdownWithContext(ctx)
		},
	})
}

func Mount(subapp any) fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotate(subapp, fx.ResultTags(`group:"subapp"`))),
	)
}
