package validation

import (
	"cotacao-fretes/internal/pkg/dto/requests"
	"testing"
)

func TestValidateQuoteRequest_Valid(t *testing.T) {
	payload := requests.RequestQuote{}
	payload.Recipient.Address.Zipcode = "12345-678"
	payload.Volumes = []struct {
		Category      int     `json:"category" validate:"required"`
		Amount        int     `json:"amount" validate:"required"`
		UnitaryWeight int     `json:"unitary_weight" validate:"required"`
		Price         int     `json:"price" validate:"required"`
		Sku           string  `json:"sku" validate:"required"`
		Height        float64 `json:"height" validate:"required"`
		Width         float64 `json:"width" validate:"required"`
		Length        float64 `json:"length" validate:"required"`
	}{
		{
			Category:      1,
			Amount:        1,
			UnitaryWeight: 5,
			Price:         100,
			Sku:           "PROD123",
			Height:        10.5,
			Width:         15.0,
			Length:        20.0,
		},
	}

	message := ValidateQuoteRequest(payload)

	if message != "" {
		t.Errorf("esperado mensagem vazia, obtido %q", message)
	}
}

func TestValidateQuoteRequest_Invalid(t *testing.T) {
	payload := requests.RequestQuote{}
	payload.Recipient.Address.Zipcode = "12345-678"
	payload.Volumes = []struct {
		Category      int     `json:"category" validate:"required"`
		Amount        int     `json:"amount" validate:"required"`
		UnitaryWeight int     `json:"unitary_weight" validate:"required"`
		Price         int     `json:"price" validate:"required"`
		Sku           string  `json:"sku" validate:"required"`
		Height        float64 `json:"height" validate:"required"`
		Width         float64 `json:"width" validate:"required"`
		Length        float64 `json:"length" validate:"required"`
	}{
		{
			Category:      1,
			Amount:        1,
			UnitaryWeight: 5,
			Price:         100,
			Sku:           "",
			Height:        10.5,
			Width:         15.0,
			Length:        20.0,
		},
	}

	message := ValidateQuoteRequest(payload)

	if message != "O campo Sku é obrigatório." {
		t.Errorf("esperado mensagem de erro, obtido %q", message)
	}
}
