package domain

import "time"

// =====================
// Quote (1:N Dispatcher)
// =====================
type Quote struct {
	ID          uint         `gorm:"primaryKey" json:"-"`
	Dispatchers []Dispatcher `gorm:"foreignKey:QuoteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"dispatchers"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// =====================
// Dispatcher (1:N Offer)
// =====================
type Dispatcher struct {
	ID                         uint   `gorm:"primaryKey" json:"-"`
	QuoteID                    uint   `gorm:"index" json:"-"`
	RequestID                  string `json:"request_id"`
	RegisteredNumberShipper    string `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int    `json:"zipcode_origin"`

	Offers []Offer `gorm:"foreignKey:DispatcherID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"offers"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// =====================
// Offer
// =====================
type Offer struct {
	ID                            uint             `gorm:"primaryKey" json:"id"`
	DispatcherID                  uint             `gorm:"index" json:"-"`
	Offer                         int              `json:"offer"`
	SimulationType                int              `json:"simulation_type"`
	CarrierID                     uint             `json:"-"`
	Carrier                       CarrierInfo      `gorm:"foreignKey:CarrierID" json:"carrier"`
	Service                       string           `json:"service"`
	ServiceCode                   string           `json:"service_code"`
	DeliveryTimeID                uint             `json:"-"`
	DeliveryTime                  DeliveryTimeInfo `gorm:"foreignKey:DeliveryTimeID" json:"delivery_time"`
	OriginalDeliveryTimeID        uint             `json:"-"`
	OriginalDeliveryTime          DeliveryTimeInfo `gorm:"foreignKey:OriginalDeliveryTimeID" json:"original_delivery_time"`
	CarrierOriginalDeliveryTimeID uint             `json:"-"`
	CarrierOriginalDeliveryTime   DeliveryTimeInfo `gorm:"foreignKey:CarrierOriginalDeliveryTimeID" json:"carrier_original_delivery_time"`
	Expiration                    time.Time        `json:"expiration"`
	CostPrice                     float64          `json:"cost_price"`
	FinalPrice                    float64          `json:"final_price"`
	Identifier                    string           `json:"identifier"`
	HomeDelivery                  bool             `json:"home_delivery"`
	Modal                         string           `json:"modal"`
	WeightID                      uint             `json:"-"`
	Weights                       WeightInfo       `gorm:"foreignKey:WeightID" json:"weights"`
	CreatedAt                     time.Time        `json:"-" swaggerignore:"true"`
	UpdatedAt                     time.Time        `json:"-" swaggerignore:"true"`
	DeletedAt                     *time.Time       `json:"-" gorm:"index" swaggerignore:"true"`
}

// =====================
// CarrierInfo
// =====================
type CarrierInfo struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Name             string     `json:"name"`
	RegisteredNumber string     `json:"registered_number"`
	StateInscription string     `json:"state_inscription"`
	Logo             string     `json:"logo"`
	Reference        int        `json:"reference"`
	CompanyName      string     `json:"company_name"`
	CreatedAt        time.Time  `json:"-" swaggerignore:"true"`
	UpdatedAt        time.Time  `json:"-" swaggerignore:"true"`
	DeletedAt        *time.Time `json:"-" gorm:"index" swaggerignore:"true"`
}

// =====================
// DeliveryTimeInfo
// =====================
type DeliveryTimeInfo struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Days          int    `json:"days"`
	EstimatedDate string `json:"estimated_date"`
	CreatedAt time.Time  `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time  `json:"-" swaggerignore:"true"`
	DeletedAt *time.Time `json:"-" gorm:"index" swaggerignore:"true"`
}

// =====================
// WeightInfo
// =====================
type WeightInfo struct {
	ID   uint `gorm:"primaryKey" json:"id"`
	Real int  `json:"real"`
	CreatedAt time.Time  `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time  `json:"-" swaggerignore:"true"`
	DeletedAt *time.Time `json:"-" gorm:"index" swaggerignore:"true"`
}

// =====================
// FreteResumo (Projection)
// =====================
type ResumeQuotes struct {
	Service                  string     `json:"transportadora" gorm:"column:transportadora"`
	QtdQuotes                int64      `json:"quantidade_cotacoes" gorm:"column:quantidade_cotacoes"`
	QtdResults               int64      `json:"quantidade_resultados" gorm:"column:quantidade_resultados"`
	TotalPrice               float64    `json:"total_preco_frete" gorm:"column:total_preco_frete"`
	AveragePrice             float64    `json:"media_preco_frete" gorm:"column:media_preco_frete"`
	CheapestShippingOverall  float64    `json:"frete_mais_barato_geral" gorm:"column:frete_mais_barato_geral"`
	CostliestShippingOverall float64    `json:"frete_mais_caro_geral" gorm:"column:frete_mais_caro_geral"`
	CreatedAt                time.Time  `json:"-" swaggerignore:"true"`
	UpdatedAt                time.Time  `json:"-" swaggerignore:"true"`
	DeletedAt                *time.Time `json:"-" gorm:"index" swaggerignore:"true"`
}
