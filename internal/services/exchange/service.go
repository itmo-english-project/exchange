package exchange

import "context"

type storage interface {
	InsertTest(ctx context.Context, comment string) error
	SelectTest(ctx context.Context) ([]string, error)
}

type Service struct {
	storage storage
}

func NewService(storage storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Insert(ctx context.Context, comment string) error {
	return s.storage.InsertTest(ctx, comment)
}

func (s *Service) Get(ctx context.Context) ([]string, error) {
	tests, err := s.storage.SelectTest(ctx)
	if err != nil {
		return nil, err
	}
	return tests, nil
}
