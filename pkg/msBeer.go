package msBeer

import (
	"context"
)

type Beer struct {
	ID       int     `json:"ID"`
	Name     string  `json:"name"`
	Brewery  string  `json:"Brewery"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}

type BeerBox struct {
	TotalPrice float32 `json:"price_total"`
}

type Repository interface {
	FetchBeers(ctx context.Context) ([]Beer, error)
	CreateBeer(ctx context.Context, b *Beer) error
	FetchBeerByID(ctx context.Context, ID string) (*Beer, error)
	FetchBoxPriceByID(ctx context.Context, ID string) (int, error)
}
