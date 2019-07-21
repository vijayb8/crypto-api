package pricing

import (
	"net/http"

	"github.com/crypto-api/pkg/platform/errors"
	"github.com/crypto-api/pkg/platform/web"
	log "github.com/sirupsen/logrus"
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
			l.Error(err)
			web.WriteErrorResponse(w, err)
			return
		}
		web.WriteSuccessResponse(w, http.StatusOK, priceList)
	})
}
