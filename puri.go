package puri

import (
	"errors"
	"net/url"
	"strings"
)

const Version string = "v0.1.5"

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

func ExtractHost(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil || len(uri) == 0 {
		return "", errors.New("invalid uri")
	}
	if parsed != nil && len(parsed.Host) == 0 {
		if len(parsed.Path) > 0 {
			return parsed.Path, nil
		}
		return "", errors.New("no host")
	}
	hp := strings.Split(parsed.Host, ":")
	return hp[0], nil
}
