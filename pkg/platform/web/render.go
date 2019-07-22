package web

import (
	"encoding/json"
	"net/http"

	"github.com/vijayb8/crypto-api/pkg/platform/errors"
)

// successResponse represents success response structure
type successResponse struct {
	Data Data        `json:"data"`
	Meta interface{} `json:"meta"`
}

// errorResponse represents error response structure
// swagger:model
type errorResponse struct {
	// array of errors
	Errors []string `json:"errors"`
	// meta information
	Meta interface{} `json:"meta"`
}

// Data represents data from jsonAPI
// swagger:model
type Data struct {
	Attributes interface{} `json:"attributes"`
}

// WriteSuccessResponse writes successful response
func WriteSuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	payload := &successResponse{
		Data: Data{
			Attributes: data,
		},
	}
	writeJSON(w, code, payload)
}

// WriteErrorResponse writes error response
func WriteErrorResponse(w http.ResponseWriter, errs ...error) {
	var code int
	var errorsList []string
	for _, v := range errs {
		if v != nil {
			if e, ok := v.(*errors.Error); ok {
				switch e.Code {
				case errors.EINTERNAL:
					code = http.StatusInternalServerError
				case errors.ENOTFOUND:
					code = http.StatusNotFound
				case errors.EINVALID:
					code = http.StatusBadRequest
				case errors.EUNAUTHORIZED:
					code = http.StatusUnauthorized
				}
			}
		}
		errorsList = append(errorsList, errors.Message(v))
	}

	if code == 0 {
		code = http.StatusInternalServerError
	}

	payload := &errorResponse{
		Errors: errorsList,
	}

	writeJSON(w, code, payload)
}

// writeJSON writes json data to response
func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	if v == nil || code == http.StatusNoContent {
		return
	}

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, "cannot encode data", http.StatusInternalServerError)
	}
}
