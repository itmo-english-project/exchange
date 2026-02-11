package exchanges

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"

	"github.com/itmo-english-project/common/pkg/storage/postgres"
)

type Storage struct {
	p *pgxpool.Pool
}

func NewStorage(db postgres.Validator) (*Storage, error) {
	if err := db.ValidateQueries([]string{
		createExchange,
	}); err != nil {
		return nil, fmt.Errorf("invalid queries: %w", err)
	}
	return &Storage{
		p: db.Pool(),
	}, nil
}

type ExchangeStatus string

const (
	ExchangeStatusPending   ExchangeStatus = "pending"
	ExchangeStatusAccepted  ExchangeStatus = "accepted"
	ExchangeStatusRejected  ExchangeStatus = "rejected"
	ExchangeStatusCancelled ExchangeStatus = "cancelled"
)

func (s *Storage) InsertExchange(ctx context.Context, fromID string, toID string, status ExchangeStatus, comment string) error {
	_, err := s.p.Exec(ctx, createExchange, fromID, toID, status, comment)
	return errors.Wrap(err, "test query")
}

func (s *Storage) GetExchanges(ctx context.Context) ([]string, error) {
	var testSelect string
	var tests []string
	rows, err := s.p.Query(ctx, testSelect)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select")
	}
	defer rows.Close()

	for rows.Next() {
		var test string
		if err = rows.Scan(&test); err != nil {
			return nil, errors.Wrap(err, "test query")
		}
		tests = append(tests, test)
	}

	return tests, nil
}
