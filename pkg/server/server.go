package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	msBeer "msBeer/pkg"
	"msBeer/pkg/beers"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type server struct {
	serverID string

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
	r.HandleFunc("/beers/{beerID:[0-9]+}/boxprice", s.BeerPrice).Methods(http.MethodGet)

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
	allBeers, err := s.beer.FetchBeers(r.Context())
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msBeer.FetchAnswer{
		Beers: allBeers,
	})
}

func (s *server) AddBeers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	var newBeer msBeer.Beer

	err = json.Unmarshal(body, &newBeer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
			Description:  "Request invalida",
		})
		return
	}

	exist, err := s.beer.CreateBeer(r.Context(), &newBeer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	if exist {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(ResponseService{
			Description: "El ID de la cerveza ya existe",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ResponseService{
		Description: "Cerveza creada",
	})
}

func (s *server) BeerDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["beerID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	answ, err := s.beer.FetchBeerByID(r.Context(), int(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	if answ == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ResponseService{
			Description: "El Id de la cerveza no existe",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msBeer.FetchByIdAnswer{
		BeerByID: *answ,
	})

}

func (s *server) BeerPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["beerID"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	answ, err := s.beer.FetchBeerByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
	}

	if answ == nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ResponseService{
			Description: "El Id de la cerveza no existe",
		})
	}

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	decoder := schema.NewDecoder()
	var options BoxPriceOpts
	if err := decoder.Decode(&options, r.URL.Query()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	apiUrl := fmt.Sprintf("%s?access_key=%s&currencies=%s,%s&format=1", os.Getenv("CURRENCY_API_URL"), os.Getenv("CURRENCY_API_KEY"), options.Currency, answ.Currency)
	fmt.Println(apiUrl)

	resp, err := http.Get(apiUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
	}

	var currency CurrencyAnswer
	err = json.Unmarshal(body, &currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ResponseService{
			ErrorContain: err.Error(),
		})
		return
	}

	currencyQty := 6

	if options.Quantity != 0 {
		currencyQty = options.Quantity
	}

	total := (currency.Quotes[fmt.Sprintf("USD%s", answ.Country)] / currency.Quotes[fmt.Sprintf("USD%s", options.Currency)]) * float32(currencyQty) * answ.Price
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msBeer.PriceByBox{
		PriceBox: msBeer.BeerBox{
			TotalPrice: total,
		},
	})

}
