package server

type ResponseService struct {
	ErrorContain string `json:"error,omitempty"`
	Description  string `json:"description,omitempty"`
}

type BoxPriceOpts struct {
	Quantity int    `schema:"quantity"`
	Currency string `schema:"currency,requiered"`
}

type CurrencyAnswer struct {
	Quotes map[string]float32 `json:"quotes"`
}
