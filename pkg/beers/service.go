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
