package ordering

// OrderReq is req json to cryptocompare
type OrderReq struct {
	Limit string `json:"limit"`
	Tysm  string `json:"tysm"`
}

// OrderResp is resp json to cryptocompare
type OrderResp struct {
	Data []DataInfo `json:"Data"`
}

type CoinInfo struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type USD struct {
	PRICE float64 `json:"PRICE"`
}

type RAW struct {
	USDCur USD `json:"USD"`
}

type DataInfo struct {
	Coin CoinInfo `json:"CoinInfo"`
	Raw  RAW      `json:"RAW"`
}
