package beers

import (
	"context"
	msBeer "msBeer/pkg"
	"msBeer/pkg/log"
)

type Service interface {
	FetchBeers(ctx context.Context) ([]msBeer.Beer, error)
	CreateBeer(ctx context.Context, b *msBeer.Beer) (bool, error)
	FetchBeerByID(ctx context.Context, ID int) (*msBeer.Beer, error)
}

type service struct {
	repository msBeer.Repository
	logger     log.Logger
}

func NewService(repository msBeer.Repository, logger log.Logger) Service {
	return &service{repository, logger}
}

func (s *service) CreateBeer(ctx context.Context, newBeer *msBeer.Beer) (bool, error) {
	exist, err := s.repository.CreateBeer(ctx, newBeer)
	if err != nil {
		s.logger.UnexpectedError(ctx, err)
		return false, err
	}
	return exist, nil
}

func (s *service) FetchBeers(ctx context.Context) ([]msBeer.Beer, error) {
	answ, err := s.repository.FetchBeers(ctx)
	if err != nil {
		s.logger.UnexpectedError(ctx, err)
		return nil, err
	}
	return answ, nil
}

func (s *service) FetchBeerByID(ctx context.Context, ID int) (*msBeer.Beer, error) {
	answ, err := s.repository.FetchBeerByID(ctx, ID)
	if err != nil {
		s.logger.UnexpectedError(ctx, err)
		return nil, err
	}
	return answ, nil
}
