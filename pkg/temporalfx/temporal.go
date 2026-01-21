package temporalfx

import (
	"fmt"
	"log/slog"

	"go.opentelemetry.io/otel"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/contrib/datadog/tracing"
	"go.temporal.io/sdk/contrib/opentelemetry"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/worker"
)

type Config struct {
	Namespace string `env:"TEMPORAL_NAMESPACE" envDefault:"default"`
	URI       string `env:"TEMPORAL_URI" envDefault:"localhost:7233"`
	TaskQueue string `env:"TEMPORAL_TASK_QUEUE" envDefault:"task-queue"`
}

func NewClient(cfg Config) (client.Client, error) {
	tracingInterceptor := tracing.NewTracingInterceptor(tracing.TracerOptions{})

	return client.Dial(client.Options{
		Namespace:    cfg.Namespace,
		HostPort:     cfg.URI,
		Interceptors: []interceptor.ClientInterceptor{tracingInterceptor},
	})
}

func NewWorker(c client.Client, cfg Config) worker.Worker {
	tracer := otel.Tracer("temporal-worker-tracer")

	tracingInterceptor, err := opentelemetry.NewTracingInterceptor(opentelemetry.TracerOptions{Tracer: tracer})
	if err != nil {
		slog.Error(fmt.Sprintf("erro ao criar TracingInterceptor: %v", err))
	}

	return worker.New(c, cfg.TaskQueue, worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{tracingInterceptor},
	})
}
