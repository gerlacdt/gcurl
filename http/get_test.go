package http

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	url := "http://localhost:8080/get"
	verbose := false

	actual, err := Get(url, verbose)
	if err != nil {
		t.Fatalf("http GET failed, %v", err)
	}

	if !strings.Contains(actual, "User-Agent") {
		t.Fatalf("expected contains User-Agend, got: %s", actual)
	}
}
