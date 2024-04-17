package puri

import (
	"testing"
)

func TestExtractParam(t *testing.T) {
	tests := []struct {
		name  string
		uri   string
		param string
		want  string
	}{
		{
			name:  "case 1: valid url and param",
			uri:   "http://example.com/?key=value",
			param: "key",
			want:  "value",
		},
		{
			name:  "case 2: valid url, param is missing",
			uri:   "http://example.com/?key=value",
			param: "missing",
			want:  "",
		},
		{
			name:  "case 3: invalid url, valid param",
			uri:   "http:/example.com?param1=value1",
			param: "param1",
			want:  "value1",
		},
		{
			name:  "case 4: empty url and param",
			uri:   "",
			param: "",
			want:  "",
		},
		// Add more cases as needed.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractParam(tt.uri, tt.param)
			if err == nil {
				if got != tt.want {
					t.Errorf("ExtractParam() = %v, want %v", got, tt.want)
				}
			} else {
				if tt.want != "" {
					t.Errorf("wanted a result: %s, but got error: %v", tt.want, err)
				}
			}
		})
	}
}

func TestExtractScheme(t *testing.T) {
	samples := []struct {
		name       string
		uri        string
		wantScheme string
		hasError   bool
	}{
		{"ftp", "ftp://example.com", "ftp", false},
		{"http", "http://example.com", "http", false},
		{"https", "https://example.com", "https", false},
		{"none", "empty uri", "", true},
		{"bad", "invalid uri", "", true},
	}

	for _, sample := range samples {
		t.Run(sample.name, func(t *testing.T) {
			gotScheme, err := ExtractScheme(sample.uri)
			if (err != nil) != sample.hasError {
				t.Errorf("ExtractScheme() error = %v, hasError %v", err, sample.hasError)
				return
			}
			if gotScheme != sample.wantScheme {
				t.Errorf("ExtractScheme() = %v, want %v", gotScheme, sample.wantScheme)
			}
		})
	}
}

func TestExtractHost(t *testing.T) {
	samples := []struct {
		name     string
		uri      string
		wantHost string
		hasError bool
	}{
		{"ftp host", "ftp://example.com", "example.com", false},
		{"http host with port", "http://example.com:8080", "example.com", false},
		{"http host with port and path", "http://example.com:8080/blah/blah?k=v", "example.com", false},
		{"simple", "example.com", "example.com", false},
		{"simpler", "host", "host", false},
	}

	for _, sample := range samples {
		t.Run(sample.name, func(t *testing.T) {
			gotHost, err := ExtractHost(sample.uri)
			if (err != nil) != sample.hasError {
				t.Errorf("ExtractHost() error = %v, hasError %v", err, sample.hasError)
				return
			}
			if gotHost != sample.wantHost {
				t.Errorf("ExtractScheme() = %v, want %v", gotHost, sample.wantHost)
			}
		})
	}
}
