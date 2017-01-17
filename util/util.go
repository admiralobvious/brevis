package util

import (
	"net/url"
)

// Error holds an error message
type ErrorResponse struct {
	Message string `json:"error"`
}

// IsValidUri checks whether a URI is valid or not
func IsValidUri(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	return true
}