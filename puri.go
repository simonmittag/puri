package puri

import "net/url"

const Version string = "v0.1"

func ExtractParam(uri string, param string) string {
	parsed, _ := url.Parse(uri)
	v := parsed.Query()
	return v.Get(param)
}
