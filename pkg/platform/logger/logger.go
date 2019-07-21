package logger

import (
	"bytes"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Get returns a new log instance
func Get(logLevel string) (*log.Logger, error) {
	l := log.New()
	l.Formatter = new(log.JSONFormatter)

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}
	l.Infof("log level: %s", level.String())
	l.Level = level

	return l, nil
}

// GetRequestFields returns a log.Fields object filled containing request information
func GetRequestFields(req *http.Request) (map[string]string, error) {
	requestFields := map[string]string{
		"method":      req.Method,
		"auth_header": req.Header.Get("Authorization"),
		"host":        req.URL.Hostname(),
		"path":        req.URL.Path,
		"query":       req.URL.Query().Encode(),
	}

	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		requestFields["body"] = string(body)
	}

	return requestFields, nil
}
