package mysql

import (
	"context"
	"errors"
	"fmt"

	"cotacao-fretes/internal/domain"

	"gorm.io/gorm"
)

func New(db *gorm.DB) *Adapter {
	return &Adapter{db}
}

type Adapter struct {
	db *gorm.DB
}

const (
	MsgNotFoundRecord string = "not found record for key: %s"
	MsgErrGetQuotes   string = "failed to get quotes in database error: %v"
)

func (a *Adapter) Insert(ctx context.Context, quote domain.Quote) (domain.Quote, error) {

	err := a.db.
		WithContext(ctx).
		Session(&gorm.Session{FullSaveAssociations: true}).
		Create(&quote).Error

	if err != nil {
		return domain.Quote{}, fmt.Errorf("failed to insert quote in database: %w", err)
	}

	return quote, nil
}

func (a *Adapter) Get(ctx context.Context, key int) (*domain.Quote, error) {
	var quote domain.Quote

	err := a.db.
		WithContext(ctx).
		Preload("Dispatchers").
		Preload("Dispatchers.Offers").
		Preload("Dispatchers.Offers.Carrier").
		Preload("Dispatchers.Offers.DeliveryTime").
		Preload("Dispatchers.Offers.OriginalDeliveryTime").
		Preload("Dispatchers.Offers.CarrierOriginalDeliveryTime").
		Preload("Dispatchers.Offers.Weights").
		Where("id = ?", key).
		First(&quote).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf(MsgNotFoundRecord, key)
		}
		return nil, fmt.Errorf(MsgErrGetQuotes, err)
	}

	return &quote, nil
}

func (a *Adapter) GetResumeQuote(ctx context.Context, limit *int) (*[]domain.ResumeQuotes, error) {

	var result []domain.ResumeQuotes

	sql := `
        SELECT
            o.service                                  AS transportadora,
            COUNT(DISTINCT q.id)                       AS quantidade_cotacoes,
            COUNT(*)                                   AS quantidade_resultados,
            ROUND(SUM(o.final_price), 2)               AS total_preco_frete,
            ROUND(AVG(o.final_price), 2)               AS media_preco_frete,
            ROUND(g.frete_mais_barato_geral, 2)        AS frete_mais_barato_geral,
            ROUND(g.frete_mais_caro_geral, 2)          AS frete_mais_caro_geral
        FROM quotes q
        INNER JOIN dispatchers d ON q.id = d.quote_id
        INNER JOIN offers o ON o.dispatcher_id = d.id
        CROSS JOIN (
            SELECT
                MIN(final_price) AS frete_mais_barato_geral,
                MAX(final_price) AS frete_mais_caro_geral
            FROM offers
        ) g
        GROUP BY o.service
        ORDER BY quantidade_resultados DESC
    `

	var args []interface{}

	if limit != nil {
		sql += " LIMIT ?"
		args = append(args, *limit)
	}

	err := a.db.
		WithContext(ctx).
		Raw(sql, args...).Scan(&result).Error
	return &result, err
}
