package exchange

import (
	"context"

	"github.com/itmo-english-project/exchange/internal/adapters/out/postgres/exchanges"
)

type storage interface {
	InsertExchange(ctx context.Context, fromID string, toID string, status exchanges.ExchangeStatus, comment string) error
	GetExchanges(ctx context.Context) ([]string, error)
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) InsertExchange(ctx context.Context, fromID string, toID string, status exchanges.ExchangeStatus, comment string) error {
	return s.storage.InsertExchange(ctx, fromID, toID, status, comment)
}

func (s *Service) GetExchanges(ctx context.Context) ([]string, error) {
	tests, err := s.storage.GetExchanges(ctx)
	if err != nil {
		return nil, err
	}
	return tests, nil
}
