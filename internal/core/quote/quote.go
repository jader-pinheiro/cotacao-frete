package quote

import (
	"context"

	"cotacao-fretes/internal/domain"
)

func New(db DBPort) *Service {
	return &Service{
		db,
	}
}

type DBPort interface {
	Insert(ctx context.Context, settlement domain.Quote) (domain.Quote, error)
	Get(ctx context.Context, key int) (*domain.Quote, error)
	GetResumeQuote(ctx context.Context, limit *int) (*[]domain.ResumeQuotes, error)
}

type Service struct {
	db DBPort
}

func (s *Service) Get(ctx context.Context, lastQuotes int) (*domain.Quote, error) {
	return s.db.Get(ctx, lastQuotes)
}

func (s *Service) InsertQuote(ctx context.Context, quotes domain.Quote) (domain.Quote, error) {
	return s.db.Insert(ctx, quotes)
}

func (s *Service) GetResumeQuote(ctx context.Context, limit *int) (*[]domain.ResumeQuotes, error) {
	return s.db.GetResumeQuote(ctx, limit)
}
