package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header  string
		want    string
		wantErr bool
	}{
		"valid":     {header: "ApiKey abc123", want: "abc123"},
		"no header": {header: "", wantErr: true},
		"malformed": {header: "Bearer abc123", wantErr: true},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h := http.Header{}
			if tc.header != "" {
				h.Set("Authorization", tc.header)
			}
			got, err := GetAPIKey(h)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if got != tc.want {
				t.Fatalf("got %q, want %q", got, tc.want)
			}
		})
	}
}
