package puri

import (
	"errors"
	"net/url"
	"strings"
)

const Version string = "v0.1.5"

func parseURL(uri string) (*url.URL, error) {
	parsed, err := url.Parse(uri)
	if err != nil || len(uri) == 0 {
		return nil, errors.New("invalid uri")
	}
	return parsed, nil
}

func ExtractParam(uri string, param string) (*string, error) {
	parsed, err := parseURL(uri)
	if err != nil {
		return nil, err
	}
	v := parsed.Query()
	r := v.Get(param)
	return &r, nil
}

func ExtractScheme(uri string) (*string, error) {
	parsed, err := parseURL(uri)
	if err != nil {
		return nil, err
	}

	if parsed != nil && len(parsed.Scheme) == 0 {
		return nil, errors.New("no scheme")
	}

	return &parsed.Scheme, nil
}

func ExtractHost(uri string) (*string, error) {
	parsed, err := parseURL(uri)
	if err != nil {
		return nil, err
	}

	if parsed != nil && len(parsed.Host) == 0 {
		if len(parsed.Path) > 0 {
			return &parsed.Path, nil
		}
		return nil, errors.New("no host")
	}

	hp := strings.Split(parsed.Host, ":")
	return &hp[0], nil
}

func ExtractPort(uri string) (*string, error) {
	parsed, err := parseURL(uri)
	if err != nil {
		return nil, err
	}

	if parsed != nil && len(parsed.Host) == 0 {
		if len(parsed.Path) > 0 {
			parsed.Host = parsed.Path
		}
		if len(parsed.Scheme) > 0 {
			parsed.Host = parsed.Scheme
		}
	}

	hp := strings.Split(parsed.Host, ":")
	if len(hp) != 2 {
		if !strings.Contains(uri, "://") {
			hp = strings.Split(uri, ":")
		}
		if len(hp) != 2 {
			return nil, errors.New("no port")
		}
	}

	return &hp[1], nil
}
