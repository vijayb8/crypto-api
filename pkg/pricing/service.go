package pricing

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/vijayb8/crypto-api/pkg/platform/errors"
	bhttp "github.com/vijayb8/crypto-api/pkg/platform/http"
)

// Service defines the values required to connect to coinmarket
type Service struct {
	URL    string
	APIKey string
}

// NewService returns new service
func NewService(url string, api_key string) (*Service, error) {
	return &Service{
		URL:    url,
		APIKey: api_key,
	}, nil
}

func (s *Service) GetPricing(queryVal *PricingReq) (*PricingResp, error) {
	httpClient := &bhttp.Client{}
	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	q.Add("start", queryVal.Start)
	q.Add("limit", queryVal.Limit)
	q.Add("convert", queryVal.Convert)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", s.APIKey)
	req.URL.RawQuery = q.Encode()

	apiResp, err := httpClient.Do(req, "CoinMarket API")

	var resp PricingResp
	if err := json.Unmarshal(apiResp, &resp); err != nil {
		return nil, errors.New(errors.EINTERNAL, "", "unmarshal_pricing_resp", err)
	}
	return &resp, nil
}

func (s *Service) ListPrices(queryVal *PricingReq) (*ListPrices, error) {

	priceList, err := s.GetPricing(queryVal)
	if err != nil {
		return nil, err
	}

	var listPrice []PricingData
	for _, v := range priceList.Data {
		listPrice = append(listPrice, PricingData{
			Currency: v.Symbol,
			Price:    v.Quote.Usd.Price,
		})
	}

	return &ListPrices{
		ListPrice: listPrice,
	}, err
}
