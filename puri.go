package puri

import (
	"errors"
	"net/url"
)

const Version string = "v0.1.4"

func ExtractParam(uri string, param string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil || len(uri) == 0 {
		return "", errors.New("invalid uri")
	}
	v := parsed.Query()
	return v.Get(param), nil
}

func ExtractScheme(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil || len(uri) == 0 {
		return "", errors.New("invalid uri")
	}
	if parsed != nil && len(parsed.Scheme) == 0 {
		return "", errors.New("no scheme")
	}
	return parsed.Scheme, nil
}
