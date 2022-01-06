package beers

import (
	"context"
	msBeer "msBeer/pkg"
)

type Service interface {
	FetchBeers(ctx context.Context) ([]msBeer.Beer, error)
	CreateBeer(ctx context.Context, b *msBeer.Beer) error
	FetchBeerByID(ctx context.Context, ID int) (*msBeer.Beer, error)
	FetchBoxPriceByID(ctx context.Context, ID int) (int, error)
}

type service struct {
	repository msBeer.Repository
}

func NewService(repository msBeer.Repository) Service {
	return &service{repository}
}

func (s *service) CreateBeer(ctx context.Context, newBeer *msBeer.Beer) error {
	return s.repository.CreateBeer(ctx, newBeer)
}

func (s *service) FetchBeers(ctx context.Context) ([]msBeer.Beer, error) {
	answ, err := s.repository.FetchBeers(ctx)
	if err != nil {
		return nil, err
	}
	return answ, nil
}

func (s *service) FetchBeerByID(ctx context.Context, ID int) (*msBeer.Beer, error) {
	answ, err := s.repository.FetchBeerByID(ctx, ID)
	if err != nil {
		return nil, err
	}
	return answ, nil
}

func (s *service) FetchBoxPriceByID(ctx context.Context, ID int) (int, error) {
	answ, err := s.repository.FetchBoxPriceByID(ctx, ID)
	if err != nil {
		return 0, err
	}
	return answ, nil
}
