package ordering

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/vijayb8/crypto-api/pkg/platform/errors"
	bhttp "github.com/vijayb8/crypto-api/pkg/platform/http"
	pricing "github.com/vijayb8/crypto-api/pkg/pricing"
)

type pricingGetter interface {
	ListPrices(req *pricing.PricingReq) (*pricing.ListPrices, error)
}

// Client defines the values required to connect to coinmarket
type Service struct {
	URL            string
	APIKey         string
	PricingService pricingGetter
}

type PriceInfo struct {
	Symbol string
	Id     int
}

// NewService returns new service
func NewService(url string, api_key string, pricingService *pricing.Service) (*Service, error) {
	return &Service{
		URL:            url,
		APIKey:         api_key,
		PricingService: pricingService,
	}, nil
}

func (s *Service) GetTopList(queryVal *OrderReq, p2Service *pricing.Service) (*Response, error) {

	httpClient := &bhttp.Client{}
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("limit", queryVal.Limit)
	q.Add("convert", queryVal.Tysm)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("authorization", "ApiKey "+s.APIKey)
	req.URL.RawQuery = q.Encode()

	apiResp, err := httpClient.Do(req, "CryptoCompare API")

	pricingResp, err := s.PricingService.ListPrices(&pricing.PricingReq{
		Limit:   queryVal.Limit,
		Start:   "1",
		Convert: "USD",
	})
	if err != nil {
		return nil, err
	}

	var priceData []*PriceInfo
	for i, v := range pricingResp.ListPrice {
		priceData = append(priceData, &PriceInfo{
			Symbol: v.Currency,
			Id:     i,
		})

	}

	var resp OrderResp
	if err := json.Unmarshal(apiResp, &resp); err != nil {
		return nil, errors.New(errors.EINTERNAL, "", "unmarshal_order_resp", err)
	}

	var coins []CoinData
	for i, v := range resp.Data {
		for _, k := range pricingResp.ListPrice {
			if k.Currency == v.Coin.ID {
				coins = append(coins, CoinData{
					ID:    v.Coin.ID,
					Name:  v.Coin.Name,
					Rank:  i,
					Price: k.Price,
				})
			}
		}
	}

	return &Response{
		CoinData: coins,
	}, nil
}
