package server

import (
	"encoding/json"
	msBeer "msBeer/pkg"
	"msBeer/pkg/beers"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	serverID string
	httpAddr string

	router http.Handler
	beer   beers.Service
}

type Server interface {
	Router() http.Handler
	FetchBeers(w http.ResponseWriter, r *http.Request)
	AddBeers(w http.ResponseWriter, r *http.Request)
	BeerDetail(w http.ResponseWriter, r *http.Request)
	BeerPrice(w http.ResponseWriter, r *http.Request)
}

func router(s *server) {
	r := mux.NewRouter()

	r.Use(newServerMiddleware(s.serverID))

	r.HandleFunc("/Beers", s.FetchBeers).Methods(http.MethodGet)
	r.HandleFunc("/Beers", s.AddBeers).Methods(http.MethodPost)
	r.HandleFunc("/beers/{beerID:[0-9]+}", s.BeerDetail).Methods(http.MethodGet)
	r.HandleFunc("/beers/{beerID:[0-9]+}", s.BeerPrice).Methods(http.MethodGet)

	s.router = r
}

func (s *server) Router() http.Handler {
	return s.router
}

func New(
	serverID string,
	Br beers.Service) Server {
	a := &server{
		serverID: serverID,
		beer:     Br}
	router(a)

	return a
}

func (s *server) FetchBeers(w http.ResponseWriter, r *http.Request) {
	gophers, _ := s.beer.FetchBeers(r.Context())
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(gophers)

}

func (s *server) AddBeers(w http.ResponseWriter, r *http.Request) {
	var ms msBeer.Beer
	_ = s.beer.CreateBeer(r.Context(), &ms)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(nil)

}

func (s *server) BeerDetail(w http.ResponseWriter, r *http.Request) {
	id := 0
	gophers, _ := s.beer.FetchBeerByID(r.Context(), id)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(gophers)

}

func (s *server) BeerPrice(w http.ResponseWriter, r *http.Request) {
	id := 0
	gophers, _ := s.beer.FetchBoxPriceByID(r.Context(), id)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(gophers)

}
