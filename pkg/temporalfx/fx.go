package temporalfx

import (
	"context"
	"cotacao-fretes/pkg/configfx"
	"cotacao-fretes/pkg/probes"
	"fmt"
	"log/slog"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
)

func ModuleClient() fx.Option {
	return fx.Options(
		configfx.Module[Config](),
		fx.Decorate(fx.Annotate(decorateLogger)),
		fx.Provide(NewClient),
		fx.Invoke(clientLifecycle),
	)
}

type pLog struct {
	fx.In
	C client.Client
	L *slog.Logger `optional:"true"`
}

func decorateLogger(p pLog) (client.Client, error) {
	if p.L == nil {
		return p.C, nil
	}
	return client.NewClientFromExisting(p.C, client.Options{
		Logger: log.NewStructuredLogger(p.L),
	})
}

func clientLifecycle(lifecycle fx.Lifecycle, c client.Client) {
	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			return nil
		},
		OnStop: func(_ context.Context) error {
			c.Close()
			return nil
		},
	})
}

func ClientReady(c client.Client) probes.Check {
	return func(ctx context.Context) error {
		_, err := c.CheckHealth(ctx, &client.CheckHealthRequest{})
		if err != nil {
			return fmt.Errorf("temporal: %w", err)
		}
		return nil
	}
}

func ModuleWorker() fx.Option {
	return fx.Options(
		ModuleClient(),
		probes.WithReady(ClientReady),
		fx.Provide(NewWorker),
		fx.Invoke(lifecycle),
	)
}

type params struct {
	fx.In
	Worker     worker.Worker
	Activities []any `group:"activities"`
	Workflows  []any `group:"workflows"`
}

func lifecycle(lc fx.Lifecycle, p params) {
	for _, a := range p.Activities {
		p.Worker.RegisterActivity(a)
	}

	for _, wf := range p.Workflows {
		p.Worker.RegisterWorkflow(wf)
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			return p.Worker.Start()
		},
		OnStop: func(_ context.Context) error {
			p.Worker.Stop()
			return nil
		},
	})
}

func Activity(a any) fx.Option {
	return fx.Provide(fx.Annotate(a, fx.ResultTags(`group:"activities"`)))
}

func Workflow(a any) fx.Option {
	return fx.Provide(fx.Annotate(a, fx.ResultTags(`group:"workflows"`)))
}
