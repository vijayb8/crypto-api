package pricing

// PricingReq is req json to coinmarket
type PricingReq struct {
	Start   string `json:"start"`
	Limit   string `json:"limit"`
	Convert string `json:"convert"`
}

// PricingResp is resp from coinmarket
type PricingResp struct {
	Data []Coindata `json:"data"`
}

type ListPrices struct {
	ListPrice []PricingData
}

type PricingData struct {
	Currency string
	Price    float64
}

type Coindata struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  Quote1 `json:"quote"`
}

type Quote1 struct {
	Usd USD `json:"USD"`
}

type USD struct {
	Price float64 `json:"price"`
}
