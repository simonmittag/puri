package puri

import (
	"errors"
	"net/url"
	"strings"
)

const Version string = "v0.1.8"

const colon = ":"
const slash = "/"
const anchor = "#"
const schemeSeparator = "://"

// ExtractQuery returns a pointer to a string containing the value of the specified parameter. Note,
// returns empty strings if value not found so can be used for optional params
func ExtractQuery(uri url.URL, param string) (*string, error) {
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

func ExtractPath(uri url.URL) (*string, error) {
	p := uri.String()

	if strings.Contains(p, colon) {
		if strings.Contains(p, schemeSeparator) {
			p = p[strings.Index(p, schemeSeparator)+3:]
		}
		if strings.Contains(p, colon) {
			p = p[strings.LastIndex(p, colon):]
		}
		if strings.Contains(p, slash) {
			p = p[strings.Index(p, slash):]
		} else {
			if strings.Contains(p, anchor) {
				p = p[strings.Index(p, anchor):]
			} else {
				p = ""
			}
		}
		p = trimQuery(p)
		return &p, nil
	}

	tld, err := IANA.hasTLDInString(p)
	if err == nil {
		lp := strings.ToLower(p)
		ltld := strings.ToLower(*tld)
		i1 := strings.Index(lp, ltld)
		p1 := p[i1+len(ltld):]
		p1 = trimQuery(p1)
		if strings.Contains(p1, slash) {
			p1 = p1[strings.Index(p1, slash):]
		}
		return &p1, nil
	}

	return &p, nil
}

func trimQuery(p string) string {
	if strings.Contains(p, "?") {
		p = p[:strings.Index(p, "?")]
	}
	return p
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

// NoPath returns an error indicating there is path specified
func NoPath() (*string, error) {
	return nil, errors.New("no path")
}
