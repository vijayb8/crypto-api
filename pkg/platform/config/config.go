package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents configs for application
type Config struct {
	Port         int    `envconfig:"PORT" required:"true"`
	LogLevel     string `envconfig:"LOG_LEVEL" default:"info"`
	HTTPTimeouts *HTTPTimeouts
	CoinMarket   *CoinMarket
}

// HTTPTimeouts represents timeouts for http server and client in seconds
type HTTPTimeouts struct {
	Client      time.Duration `envconfig:"HTTP_TIMEOUT_CLIENT" default:"5s"`
	ServerRead  time.Duration `envconfig:"HTTP_TIMEOUT_SERVER_READ" default:"5s"`
	ServerWrite time.Duration `envconfig:"HTTP_TIMEOUT_SERVER_WRITE" default:"10s"`
}

// CoinMarket represents the parameter for CoinMarket
type CoinMarket struct {
	ApiKey string `envconfig:"CoinMarket_API_Key" required:"true"`
	URL    string `envconfig:"CoinMarket_URL" required:"true"`
}

// Get returns config based on environment vars
func Get() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)

	return &cfg, err
}
