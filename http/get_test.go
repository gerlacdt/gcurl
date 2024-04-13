package http

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	url := "http://localhost:8080/get"
	verbose := false
	headers := make([]string, 0)

	params, err := NewGetParams(url, verbose, headers)
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := Get(params)
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
