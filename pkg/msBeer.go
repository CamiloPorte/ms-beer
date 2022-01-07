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

type FetchAnswer struct {
	Beers []Beer `json:"beers"`
}

type FetchByIdAnswer struct {
	BeerByID Beer `json:"beer"`
}

type PriceByBox struct {
	PriceBox BeerBox `json:"BeerBox"`
}

//Se añade interfaces para simular el funcionamiento del almacenamiento del sistema.
type Repository interface {
	FetchBeers(ctx context.Context) ([]Beer, error)
	CreateBeer(ctx context.Context, b *Beer) (bool, error)
	FetchBeerByID(ctx context.Context, ID int) (*Beer, error)
}
