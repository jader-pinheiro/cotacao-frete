package clientcredentialsfx

import (
	"context"
	"cotacao-fretes/pkg/httpfx"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/clientcredentials"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type Config struct {
	ClientID     string `env:"CLIENT_ID"`
	ClientSecret string `env:"CLIENT_SECRET"`
	TokenURL     string `env:"CLIENT_TOKEN_URL"`
}

type ClientCredentialsWrapper struct {
	clientcredentials.Config
}

type Clienter interface {
	Client(ctx context.Context) httpfx.HTTPClient
}

func (c *ClientCredentialsWrapper) Client(ctx context.Context) httpfx.HTTPClient {
	return httptrace.WrapClient(c.Config.Client(ctx), httptrace.RTWithResourceNamer(func(h *http.Request) string {
		return fmt.Sprintf("%s %s://%s%s", h.Method, h.URL.Scheme, h.URL.Host, h.URL.Path)
	}))
}

func New(cfg Config) *ClientCredentialsWrapper {
	return &ClientCredentialsWrapper{
		Config: clientcredentials.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			TokenURL:     cfg.TokenURL,
		},
	}
}
