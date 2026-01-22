package requests

type RequestQuote struct {
	Recipient struct {
		Address struct {
			Zipcode string `json:"zipcode" validate:"required"`
		} `json:"address" validate:"required,dive"`
	} `json:"recipient" validate:"required,dive"`
	Volumes []struct {
		Category      int     `json:"category" validate:"required"`
		Amount        int     `json:"amount" validate:"required"`
		UnitaryWeight int     `json:"unitary_weight" validate:"required"`
		Price         int     `json:"price" validate:"required"`
		Sku           string  `json:"sku" validate:"required"`
		Height        float64 `json:"height" validate:"required"`
		Width         float64 `json:"width" validate:"required"`
		Length        float64 `json:"length" validate:"required"`
	} `json:"volumes" validate:"required,dive"`
}
