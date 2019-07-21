package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents configs for application
type Config struct {
	Port          int    `envconfig:"PORT" required:"true"`
	LogLevel      string `envconfig:"LOG_LEVEL" default:"info"`
	HTTPTimeouts  *HTTPTimeouts
	CoinMarket    *CoinMarket
	CryptoCompare *CryptoCompare
}

// HTTPTimeouts represents timeouts for http server and client in seconds
type HTTPTimeouts struct {
	Client      time.Duration `envconfig:"HTTP_TIMEOUT_CLIENT" default:"5s"`
	ServerRead  time.Duration `envconfig:"HTTP_TIMEOUT_SERVER_READ" default:"5s"`
	ServerWrite time.Duration `envconfig:"HTTP_TIMEOUT_SERVER_WRITE" default:"10s"`
}

// CoinMarket represents the parameter for CoinMarket
type CoinMarket struct {
	ApiKey string `envconfig:"CoinMarket_API_Key" required:"true" default:"119407f2-930c-47a4-8705-9edaab401be6"`
	URL    string `envconfig:"CoinMarket_URL" required:"true" default:"https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"`
}

// CryptoCompare represents the parameter for CryptoCompare
type CryptoCompare struct {
	ApiKey string `envconfig:"CryptoCompare_API_Key" required:"true" default:"9a2717a8351f22b96e6dc29d57ce533d1c1c9cf3b74cb4129774414c113b6756"`
	URL    string `envconfig:"CryptoCompare_URL" required:"true" default:"https://min-api.cryptocompare.com/data/top/totalvolfull"`
}

// Get returns config based on environment vars
func Get() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	return &cfg, err
}
