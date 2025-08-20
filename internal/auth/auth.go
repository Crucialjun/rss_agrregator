package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("authorization header is missing")
	}

	vals := strings.Split(val," ")

	if len(vals) != 2 {
		return "", errors.New("invalid authorization header format")
	}

	if(vals[0] != "Bearer") {
		return "", errors.New("authorization header must start with 'Bearer'")
	}

	return vals[1], nil
}
