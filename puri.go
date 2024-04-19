package puri

import (
	"errors"
	"net/url"
	"strings"
)

const Version string = "v0.1.5"

const colon = ":"
const schemeSeparator = "://"

// ExtractParam returns a pointer to a string containing the value of the specified parameter. Note,
// returns empty strings if value not found so can be used for optional params
func ExtractParam(uri url.URL, param string) (*string, error) {
	v := uri.Query()
	r := v.Get(param)
	return &r, nil
}

// ExtractScheme extracts the scheme from a given URL.
func ExtractScheme(uri url.URL) (*string, error) {
	if len(uri.Scheme) == 0 {
		return NoScheme()
	}

	return &uri.Scheme, nil
}

// ExtractHost extracts the host from the given URL. If the host is in the format "host:port", it splits the host
// and returns only the host part.
func ExtractHost(uri url.URL) (*string, error) {
	if len(uri.Host) == 0 {
		if len(uri.Path) > 0 {
			return &uri.Path, nil
		}
		return NoHost()
	}

	hp := strings.Split(uri.Host, colon)
	return &hp[0], nil
}

// ExtractPort extracts the port from a given URL if present, otherwise returns an error.
func ExtractPort(uri url.URL) (*string, error) {
	if len(uri.Host) == 0 {
		if len(uri.Path) > 0 {
			uri.Host = uri.Path
		}
		if len(uri.Scheme) > 0 {
			uri.Host = uri.Scheme
		}
	}

	hp := strings.Split(uri.Host, colon)
	if len(hp) != 2 {
		if !strings.Contains(uri.String(), schemeSeparator) {
			hp = strings.Split(uri.String(), colon)
		}
		if len(hp) != 2 {
			return NoPort()
		}
	}

	return &hp[1], nil
}

// NoPort returns an error indicating that there is no port associated with the URL. It is used
// in the ExtractPort function to handle cases where the port cannot be extracted from the URL.
func NoPort() (*string, error) {
	p := ""
	return &p, nil
}

// NoScheme returns an error indicating that no scheme was provided.
// It returns a nil string pointer and an error with the message "no scheme".
func NoScheme() (*string, error) {
	return nil, errors.New("no scheme")
}

// NoHost returns an error indicating that there is no host available.
func NoHost() (*string, error) {
	return nil, errors.New("no host")
}
