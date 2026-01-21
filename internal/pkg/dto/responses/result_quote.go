package responses

import "time"

// ResultQuote é a estrutura de nível superior para o resultado da cotação.
type ResultQuote struct {
	Dispatchers []Dispatcher `json:"dispatchers"`
}

// Dispatcher representa um despachante no resultado da cotação.
type Dispatcher struct {
	ID                         string  `json:"id"`
	RequestID                  string  `json:"request_id"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int     `json:"zipcode_origin"`
	Offers                     []Offer `json:"offers"`
}

// Offer representa uma oferta de frete de uma transportadora.
type Offer struct {
	Offer                       int              `json:"offer"`
	SimulationType              int              `json:"simulation_type"`
	Carrier                     CarrierInfo      `json:"carrier"`
	Service                     string           `json:"service"`
	ServiceCode                 string           `json:"service_code"`
	DeliveryTime                DeliveryTimeInfo `json:"delivery_time"`
	Expiration                  time.Time        `json:"expiration"`
	CostPrice                   float64          `json:"cost_price"`
	FinalPrice                  float64          `json:"final_price"`
	Weights                     WeightInfo       `json:"weights"`
	OriginalDeliveryTime        DeliveryTimeInfo `json:"original_delivery_time"`
	Identifier                  string           `json:"identifier"`
	HomeDelivery                bool             `json:"home_delivery"`
	CarrierOriginalDeliveryTime DeliveryTimeInfo `json:"carrier_original_delivery_time"`
	Modal                       string           `json:"modal"`
}

// CarrierInfo contém detalhes sobre a transportadora.
type CarrierInfo struct {
	Name             string `json:"name"`
	RegisteredNumber string `json:"registered_number"`
	StateInscription string `json:"state_inscription"`
	Logo             string `json:"logo"`
	Reference        int    `json:"reference"`
	CompanyName      string `json:"company_name"`
}

// DeliveryTimeInfo contém detalhes sobre o tempo de entrega.
type DeliveryTimeInfo struct {
	Days          int    `json:"days"`
	EstimatedDate string `json:"estimated_date"`
}

// WeightInfo contém detalhes sobre o peso.
type WeightInfo struct {
	Real int `json:"real"`
}
