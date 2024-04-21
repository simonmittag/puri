package puri

import (
	"net/url"
	"testing"
)

func TestHasTLDInUrl(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		wantTLD *string
		wantErr bool
	}{
		{
			name:    "Has .COM TLD",
			uri:     "http://example.com",
			wantTLD: p(".COM"),
			wantErr: false,
		},
		{
			name:    "Has .DE TLD",
			uri:     "http://example.de/blahblah?k=v",
			wantTLD: p(".DE"),
			wantErr: false,
		},
		{
			name:    "Has no TLD",
			uri:     "http://example",
			wantTLD: nil,
			wantErr: true,
		},
		{
			name:    "Has .ORG TLD in path",
			uri:     "http://example.org/blah",
			wantTLD: p(".ORG"),
			wantErr: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uri, _ := url.Parse(tt.uri)

			got, err := NewTLDS().hasTLDInUrl(*uri)
			if tt.wantErr {
				if err == nil {
					t.Error("want error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("don't want error but got: %v", err)
				}
				if *tt.wantTLD != *got {
					t.Errorf("tld doesn't match, want: %s, got: %s", *tt.wantTLD, *got)
				}
			}
		})
	}
}

func TestHasTLDInString(t *testing.T) {
	tests := []struct {
		name    string
		uri     string
		wantTLD *string
		wantErr bool
	}{
		{
			name:    "Has .COM TLD",
			uri:     "http://example.com",
			wantTLD: p(".COM"),
			wantErr: false,
		},
		{
			name:    "Has .DE TLD",
			uri:     "http://example.de/blahblah?k=v",
			wantTLD: p(".DE"),
			wantErr: false,
		},
		{
			name:    "Has no TLD",
			uri:     "http://example",
			wantTLD: nil,
			wantErr: true,
		},
		{
			name:    "Has .ORG TLD in path",
			uri:     "http://example.org/blah",
			wantTLD: p(".ORG"),
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IANA.hasTLDInString(tt.uri)
			if tt.wantErr {
				if err == nil {
					t.Error("want error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("don't want error but got: %v", err)
				}
				if *tt.wantTLD != *got {
					t.Errorf("tld doesn't match, want: %s, got: %s", *tt.wantTLD, *got)
				}
			}
		})
	}
}
