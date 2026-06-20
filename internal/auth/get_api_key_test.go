package auth

import (
	"net/http"
	"testing"
)

func TestFirst(t *testing.T) {
	if 2+2 != 4 {
		t.Fatalf("this should never fail")
	}
}

func TestGetApiKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	req.Header.Set("Authorization", "ApiKey 12345abc")
	key, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("couldn't get an api key")
	}
	if key != "12345abc" {
		t.Fatalf("couldn't get proper api key")
	}
}
