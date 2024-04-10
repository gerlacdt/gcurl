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
		t.Errorf("http GET failed, %v", err)
	}

	expectedStatusCode := "200"
	if !strings.Contains(actual.StatusCode, expectedStatusCode) {
		t.Errorf("expected statusCode: %s, got: %s", expectedStatusCode, actual.StatusCode)
	}

	body := string(actual.Body)
	if !strings.Contains(body, "User-Agent") {
		t.Errorf("expected User-Agent, got: %v", actual)
	}

	if !strings.Contains(body, "origin") {
		t.Errorf("expected origin, got: %v", actual)
	}
}
