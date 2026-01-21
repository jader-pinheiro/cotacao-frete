package requesthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"cotacao-fretes/internal/domain"
	"cotacao-fretes/internal/pkg/dto/requests"
	"cotacao-fretes/pkg/httpfx"
)

type Config struct {
	BaseURL      string `env:"BASE_URL_QUOTE"`
	QuoteURI     string `env:"QUOTE_URI"`
	ClientID     string `env:"UNIQUE_CLIENT_ID"`
	ClientSecret string `env:"UNIQUE_CLIENT_SECRET"`
	TokenURL     string `env:"UNIQUE_CLIENT_TOKEN_URL"`
}

func New(cfg Config) *Client {
	clienter := httpfx.New()
	baseURL := strings.TrimSpace(cfg.BaseURL)
	quoteURI := strings.TrimSpace(cfg.QuoteURI)

	return &Client{
		client:   clienter,
		baseURL:  baseURL,
		quoteURI: quoteURI,
	}
}

var (
	ErrCreate           = errors.New("error when requesting: to get quote")
	ErrNotFound         = errors.New("error when requesting: url not found")
	ErrMethodNotAllowed = errors.New("error when requesting: method not allowed")
	ErrGatewayTimeout   = errors.New("error when requesting: gateway timeout")
)

type Client struct {
	client   httpfx.HTTPClient
	baseURL  string
	quoteURI string
}

func (sp *Client) Request(ctx context.Context, method string, path string, body any) (*http.Request, error) {
	var reader io.Reader

	if body != nil {
		payload, _ := json.Marshal(body)
		reader = bytes.NewReader(payload)
	}

	url := fmt.Sprintf("%s/%s", sp.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, url, reader)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (r *Client) GetQuoteWithPayload(params requests.RequestQuote) (domain.Quote, error) {

	body := mapPayloadToRequestGetQuote(params)
	req, err := r.Request(context.Background(), "POST", r.quoteURI, body)
	if err != nil {
		return domain.Quote{}, err
	}

	res, err := r.client.Do(req)

	if err != nil {

		return domain.Quote{}, err
	}

	if res.StatusCode == http.StatusNotFound {
		return domain.Quote{}, ErrNotFound
	}

	if res.StatusCode == http.StatusMethodNotAllowed {
		return domain.Quote{}, ErrMethodNotAllowed
	}

	if res.StatusCode == http.StatusGatewayTimeout {
		return domain.Quote{}, ErrGatewayTimeout
	}

	if res.StatusCode >= http.StatusBadRequest {
		return domain.Quote{}, errors.New("error when requesting: bad request (status code: " + res.Status + ")")

	}

	if res.StatusCode == http.StatusForbidden {
		return domain.Quote{}, errors.New("error when requesting: forbidden (status code: " + res.Status + ")")

	}

	var result domain.Quote
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return domain.Quote{}, err
	}

	defer res.Body.Close()
	return result, nil

}

func mapPayloadToRequestGetQuote(payload requests.RequestQuote) requests.RequestGetQuote {
	zipcode, _ := strconv.Atoi(payload.Recipient.Address.Zipcode)

	volumes := make([]requests.Volume, len(payload.Volumes))
	for i, v := range payload.Volumes {
		volumes[i] = requests.Volume{
			Category:      strconv.Itoa(v.Category),
			Amount:        v.Amount,
			UnitaryWeight: v.UnitaryWeight,
			UnitaryPrice:  float64(v.Price),
			Price:         v.Price,
			Sku:           v.Sku,
			Height:        v.Height,
			Width:         v.Width,
			Length:        v.Length,
		}
	}

	return requests.RequestGetQuote{
		Shipper: requests.Shipper{
			RegisteredNumber: "25438296000158",
			Token:            "1d52a9b6b78cf07b08586152459a5c90",
			PlatformCode:     "5AKVkHqCn",
		},
		Recipient: requests.Recipient{
			Zipcode: zipcode,
			Country: "BRA",
			Type:    1,
		},
		Dispatchers: []requests.Dispatcher{
			{
				RegisteredNumber: "25438296000158",
				Zipcode:          29161376,
				TotalPrice:       0.0,
				Volumes:          volumes,
			},
		},
		SimulationType: []int{0},
	}
}
