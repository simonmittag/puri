package puri

import "net/url"

const Version string = "v0.1.2"

func ExtractParam(uri string, param string) (string, error) {
	parsed, err := url.Parse(uri)
	if err != nil || len(uri) == 0 {
		return "", err
	}
	v := parsed.Query()
	return v.Get(param), nil
}
