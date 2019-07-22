package ordering

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/vijayb8/crypto-api/pkg/platform/errors"
	"github.com/vijayb8/crypto-api/pkg/platform/web"
	pricing "github.com/vijayb8/crypto-api/pkg/pricing"
)

func GetTopList(client *Client, p2Service *pricing.Service, l *log.Logger) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limit := r.URL.Query().Get("limit")
		topList, err := client.GetTopList(&OrderReq{
			Limit: limit,
			Tysm:  "USD",
		}, p2Service)
		if err != nil {
			err = errors.Wrap(err, "can't get top list")
			l.Error(err)
			web.WriteErrorResponse(w, err)
			return
		}
		web.WriteSuccessResponse(w, http.StatusOK, topList)
	})
}
