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
		testSelect,
		testInsert,
	}); err != nil {
		return nil, fmt.Errorf("invalid queries: %w", err)
	}
	return &Storage{
		p: db.Pool(),
	}, nil
}

var tests []string

func (s *Storage) InsertTest(ctx context.Context, comment string) error {
	_, err := s.p.Exec(ctx, testInsert, comment)
	return errors.Wrap(err, "test query")
}

func (s *Storage) SelectTest(ctx context.Context) ([]string, error) {
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
