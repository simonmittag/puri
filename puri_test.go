package puri

import (
	"testing"
)

func p[T any](t T) *T {
	return &t
}

func TestExtractParam(t *testing.T) {
	tests := []struct {
		name  string
		uri   string
		param string
		want  *string
	}{
		{
			name:  "case 1: valid url and param",
			uri:   "http://example.com/?key=value",
			param: "key",
			want:  p("value"),
		},
		{
			name:  "case 2: valid url, param is missing",
			uri:   "http://example.com/?key=value",
			param: "missing",
		},
		{
			name:  "case 3: invalid url, valid param",
			uri:   "http:/example.com?param1=value1",
			param: "param1",
			want:  p("value1"),
		},
		{
			name:  "case 4: empty url and param",
			uri:   "",
			param: "",
		},
		// Add more cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractParam(tt.uri, tt.param)
			if err == nil {
				if tt.want != nil && (*got != *tt.want) {
					t.Errorf("ExtractParam() = %v, want %v", got, tt.want)
				}
			} else {
				if tt.want != nil {
					t.Errorf("wanted a result: %s, but got error: %v", *tt.want, err)
				}
			}
		})
	}
}

func TestExtractScheme(t *testing.T) {
	tests := []struct {
		name       string
		uri        string
		wantScheme *string
		wantError  bool
	}{
		{"ftp", "ftp://example.com", p("ftp"), false},
		{"http", "http://example.com", p("http"), false},
		{"https", "https://example.com", p("https"), false},
		{"none", "empty uri", nil, true},
		{"bad", "invalid uri", nil, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotScheme, err := ExtractScheme(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("ExtractScheme() error = %v, hasError %v", err, tc.wantError)
				return
			}

			if tc.wantScheme != nil && (*gotScheme != *tc.wantScheme) {
				t.Errorf("ExtractScheme() = %v, want %v", gotScheme, tc.wantScheme)
			}
		})
	}
}

func TestExtractHost(t *testing.T) {
	tests := []struct {
		name      string
		uri       string
		wantHost  *string
		wantError bool
	}{
		{"ftp host", "ftp://example.com", p("example.com"), false},
		{"http host with port", "http://example.com:8080", p("example.com"), false},
		{"http host with port and path", "http://example.com:8080/blah/blah?k=v", p("example.com"), false},
		{"simple", "example.com", p("example.com"), false},
		{"simpler", "host", p("host"), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotHost, err := ExtractHost(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("ExtractHost() error = %v, hasError %v", err, tc.wantError)
				return
			}
			if *gotHost != *tc.wantHost {
				t.Errorf("ExtractScheme() = %v, want %v", gotHost, tc.wantHost)
			}
		})
	}
}

func TestExtractPort(t *testing.T) {
	tests := []struct {
		name      string
		uri       string
		wantPort  *string
		wantError bool
	}{
		{"ftp host no port", "ftp://example.com", nil, true},
		{"http host with port", "http://example.com:8080", p("8080"), false},
		{"http host with port and path", "http://example.com:8080/blah/blah?k=v", p("8080"), false},
		{"simple", "example.com:80", p("80"), false},
		{"simpler", "host", nil, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPort, err := ExtractPort(tc.uri)
			if (err != nil) != tc.wantError {
				t.Errorf("ExtractPort() error = %v, hasError %v", err, tc.wantError)
				return
			}
			if tc.wantPort != nil && (*gotPort != *tc.wantPort) {
				t.Errorf("ExtractPort() = %v, want %v", gotPort, tc.wantPort)
			}
		})
	}
}
