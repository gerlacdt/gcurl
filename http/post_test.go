package http

import (
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	url := "http://localhost:8080/post"
	verbose := false
	headers := make([]string, 0)
	reader := strings.NewReader("\"foo\": \"bar\"")
	params, err := NewPostParams(url, verbose, headers, reader, "")
	if err != nil {
		t.Errorf("creatin PostParams failed, %v", err)
	}

	actual, err := Post(params)
	if err != nil {
		t.Errorf("http POST failed, %v", err)
	}

	expectedStatusCode := "200"
	if !strings.Contains(actual.StatusCode, expectedStatusCode) {
		t.Errorf("expected StatusCode: %s, got: %s", expectedStatusCode, actual.StatusCode)
	}

	body := string(actual.Body)
	expectedWord := "foo"
	if !strings.Contains(body, expectedWord) {
		t.Errorf("should contain: %s, got: %s", expectedWord, body)
	}
}
