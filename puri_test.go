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
		uri        string
		name       string
		wantScheme string
		hasError   bool
	}{
		{"ftp://example.com", "ftp scheme", "ftp", false},
		{"http://example.com", "http scheme", "http", false},
		{"https://example.com", "https scheme", "https", false},
		{"", "empty uri", "", true},
		{"bad_string", "invalid uri", "", true},
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
