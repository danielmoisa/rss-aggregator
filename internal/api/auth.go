package api

import (
	"errors"
	"net/http"
	"strings"
)

const API_KEY = "ApiKey"

// Get Api key from the http headers
// Example: Authorization: ApiKey {apikey}
func getApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != API_KEY {
		return "", errors.New("invalid auth header format")
	}

	return vals[1], nil

}
