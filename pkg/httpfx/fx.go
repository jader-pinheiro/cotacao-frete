package httpfx

import (
	"net/http"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(New),
	)
}

func New() *http.Client {
	return http.DefaultClient
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
