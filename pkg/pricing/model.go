package pricing

// PricingReq is req json to coinmarket
type PricingReq struct {
	Start   string `json:"start"`
	Limit   string `json:"limit"`
	Convert string `json:"convert"`
}

// PricingResp is resp from coinmarket
type PricingResp struct {
	data []coindata `json:"data"`
}

type coindata struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Quote  quote  `json:"quote"`
}

type quote struct {
	Usd USD `json:"USD"`
}

type USD struct {
	Price float64 `json:"price"`
}
