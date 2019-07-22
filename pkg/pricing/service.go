package pricing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vijayb8/crypto-api/pkg/platform/errors"
	bhttp "github.com/vijayb8/crypto-api/pkg/platform/http"
)

// Client defines the values required to connect to coinmarket
type Client struct {
	URL    string
	APIKey string
}

// NewClient returns new service
func NewClient(url string, api_key string) (*Client, error) {
	return &Client{
		URL:    url,
		APIKey: api_key,
	}, nil
}

func (client *Client) ListPrices(queryVal *PricingReq) (*PricingResp, error) {
	httpClient := &bhttp.Client{}
	req, err := http.NewRequest("GET", client.URL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	q := url.Values{}
	q.Add("start", queryVal.Start)
	q.Add("limit", queryVal.Limit)
	q.Add("convert", queryVal.Convert)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", client.APIKey)
	req.URL.RawQuery = q.Encode()

	apiResp, err := httpClient.Do(req, "CoinMarket API")

	var resp PricingResp
	if err := json.Unmarshal(apiResp, &resp); err != nil {
		return nil, errors.New(errors.EINTERNAL, "", "unmarshal_pricing_resp", err)
	}
	return &resp, nil
}
