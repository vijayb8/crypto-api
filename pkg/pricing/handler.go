package pricing

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vijayb8/crypto-api/pkg/platform/errors"
	"github.com/vijayb8/crypto-api/pkg/platform/web"
)

func GetPricing(client *Client, l *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := r.URL.Query().Get("start")
		limit := r.URL.Query().Get("limit")
		priceList, err := client.ListPrices(&PricingReq{
			Start:   start,
			Limit:   limit,
			Convert: "USD",
		})
		if err != nil {
			err = errors.Wrap(err, "can't get price list")
			web.WriteErrorResponse(w, err)
			return
		}
		web.WriteSuccessResponse(w, http.StatusOK, priceList)
	})
}
