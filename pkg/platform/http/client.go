package http

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/vijayb8/crypto-api/pkg/platform/errors"
)

// Client represents http client
type Client struct {
	client http.Client
	logger *log.Logger
}

// NewClient creates a new client with circuit breaker and retry mechanism
func NewClient(timeout time.Duration, logger *log.Logger) *Client {
	return &Client{
		client: http.Client{
			Timeout: timeout,
		},
		logger: logger,
	}
}

// ValidateResponse validates the response from the http client
func (c *Client) ValidateResponse(res *http.Response, err error, op string) ([]byte, error) {
	if err != nil {
		return nil, errors.New(errors.EINTERNAL, "", op, err)
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New(errors.ENOTFOUND, "", op, err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			c.logger.Error(err)
		}
	}()
	return ioutil.ReadAll(res.Body)
}

// Get sends an HTTP GET request and returns an HTTP response
func (c *Client) Get(url, op string) ([]byte, error) {
	res, err := c.client.Get(url)
	return c.ValidateResponse(res, err, op)
}

// Do sends an HTTP GET request and returns an HTTP response
func (c *Client) Do(req *http.Request, op string) ([]byte, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return c.ValidateResponse(res, err, op)
}
