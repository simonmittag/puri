package puri

import (
	"net/url"
	"testing"
)

func p[T any](t T) *T {
	return &t
}

func toTestURL(uri string) url.URL {
	u, e := url.Parse(uri)
	if e != nil {
		panic("test setup failure")
	}
	return *u
}

func TestExtractParam(t *testing.T) {
	tests := []struct {
		name  string
		uri   url.URL
		param string
		want  *string
	}{
		{
			name:  "case 1: valid url and param",
			uri:   toTestURL("http://example.com/?key=value"),
			param: "key",
			want:  p("value"),
		},
		{
			name:  "case 2: valid url, param is missing",
			uri:   toTestURL("http://example.com/?key=value"),
			param: "missing",
		},
		{
			name:  "case 3: invalid url, valid param",
			uri:   toTestURL("http:/example.com?param1=value1"),
			param: "param1",
			want:  p("value1"),
		},
		{
			name:  "case 4: empty url and param",
			uri:   toTestURL(""),
			param: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractQuery(tt.uri, tt.param)
			if err == nil {
				if tt.want != nil && (*got != *tt.want) {
					t.Errorf("got param: %v, want param: %v", got, tt.want)
				}
			} else {
				if tt.want != nil {
					t.Errorf("wanted param: %s, but got error: %v", *tt.want, err)
				}
			}
		})
	}
}

func TestExtractScheme(t *testing.T) {
	tests := []struct {
		name       string
		uri        url.URL
		wantScheme *string
		wantError  bool
	}{
		{"ftp", toTestURL("ftp://example.com"), p("ftp"), false},
		{"http", toTestURL("http://example.com"), p("http"), false},
		{"https", toTestURL("https://example.com"), p("https"), false},
		{"none", toTestURL("empty uri"), nil, true},
		{"bad", toTestURL("invalid uri"), nil, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotScheme, err := ExtractScheme(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("got error: %v, want error: %v", err, tc.wantError)
				return
			}

			if tc.wantScheme != nil && (*gotScheme != *tc.wantScheme) {
				t.Errorf("got scheme: %v, want scheme: %v", gotScheme, tc.wantScheme)
			}
		})
	}
}

func TestExtractHost(t *testing.T) {
	tests := []struct {
		name      string
		uri       url.URL
		wantHost  *string
		wantError bool
	}{
		{"ftp host", toTestURL("ftp://example.com"), p("example.com"), false},
		{"http host with port", toTestURL("http://example.com:8080"), p("example.com"), false},
		{"http host with port and path", toTestURL("http://example.com:8080/blah/blah?k=v"), p("example.com"), false},
		{"simple", toTestURL("example.com"), p("example.com"), false},
		{"simpler", toTestURL("host"), p("host"), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotHost, err := ExtractHost(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("got error: %v, want error: %v", err, tc.wantError)
				return
			}
			if *gotHost != *tc.wantHost {
				t.Errorf("got host: %v, want host: %v", gotHost, tc.wantHost)
			}
		})
	}
}

func TestExtractPort(t *testing.T) {
	tests := []struct {
		name      string
		uri       url.URL
		wantPort  *string
		wantError bool
	}{
		{"ftp host no port", toTestURL("ftp://example.com"), p(""), false},
		{"http host with port", toTestURL("http://example.com:8080"), p("8080"), false},
		{"http host with port and path", toTestURL("http://example.com:8080/blah/blah?k=v"), p("8080"), false},
		{"simple", toTestURL("example.com:80"), p("80"), false},
		{"simpler", toTestURL("host"), p(""), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPort, err := ExtractPort(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("got error: %v, want error: %v", err, tc.wantError)
				return
			}
			if tc.wantPort != nil && (*gotPort != *tc.wantPort) {
				t.Errorf("got port: %v, want port: %v", gotPort, tc.wantPort)
			}
		})
	}
}

func TestExtractPath(t *testing.T) {
	tests := []struct {
		name      string
		uri       url.URL
		wantPath  *string
		wantError bool
	}{
		{"Path with ftp", toTestURL("ftp://example.com/path"), p("/path"), false},
		{"Path with http and port", toTestURL("http://example.com:8080/path"), p("/path"), false},
		{"Path with http and anchor", toTestURL("http://example.com/path#a"), p("/path#a"), false},
		{"Path with http, port and query", toTestURL("http://example.com:8080/path/blah?k=v"), p("/path/blah"), false},
		{"Path with http, port and query and anchor", toTestURL("http://example.com:8080/path/blah#a?k=v"), p("/path/blah#a"), false},
		{"Path with tld host", toTestURL("example.com/path"), p("/path"), false},
		{"Path with host and query", toTestURL("example.com/path?k=v"), p("/path"), false},
		{"Path with host and query and anchor", toTestURL("example.com/path#x?k=v"), p("/path#x"), false},
		{"Path with host and port", toTestURL("example.com:8080/path/a"), p("/path/a"), false},
		{"Path with host and port and query", toTestURL("example.com:8080/path/a?k=v"), p("/path/a"), false},
		{"Path with host and port and anchor", toTestURL("example.com:8080/path/a#x?k=v"), p("/path/a#x"), false},
		{"Path only", toTestURL("some/path"), p("some/path"), false},
		{"Absolute Path only", toTestURL("/some/path/more"), p("/some/path/more"), false},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPath, err := ExtractPath(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("got error = %v, want error %v", err, tc.wantError)
				return
			}
			if tc.wantPath != nil && (*gotPath != *tc.wantPath) {
				t.Errorf("got path: %v, want path: %v", *gotPath, *tc.wantPath)
			}
		})
	}
}
