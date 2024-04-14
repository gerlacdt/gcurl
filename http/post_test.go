package http

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestPost_validRequestFromInputStream_bodyOk(t *testing.T) {
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

	body := string(actual.Body)
	expectedWord := "foo"
	if !strings.Contains(body, expectedWord) {
		t.Errorf("should contain: %s, got: %s", expectedWord, body)
	}
}

func TestPost_validRequestFromArgs_bodyOk(t *testing.T) {
	url := "http://localhost:8080/post"
	verbose := false
	headers := make([]string, 0)
	requestBody := "\"foo\": \"bar\""
	params, err := NewPostParams(url, verbose, headers, os.Stdin, requestBody) // body should take precedence
	if err != nil {
		t.Errorf("creatin PostParams failed, %v", err)
	}

	actual, err := Post(params)
	if err != nil {
		t.Errorf("http POST failed, %v", err)
	}

	body := string(actual.Body)
	expectedWord := "foo"
	if !strings.Contains(body, expectedWord) {
		t.Errorf("should contain: %s, got: %s", expectedWord, body)
	}
}

func TestPost_customHeader_customHeaderSet(t *testing.T) {
	url := "http://localhost:8080/post"
	verbose := false
	headers := make([]string, 0)
	customHeaderKey := "X-Custom"
	customHeaderValue := "mycustomheader"
	headers = append(headers, fmt.Sprintf("%s: %s", customHeaderKey, customHeaderValue))
	reader := strings.NewReader("\"foo\": \"bar\"")
	params, err := NewPostParams(url, verbose, headers, reader, "")
	if err != nil {
		t.Errorf("creatin PostParams failed, %v", err)
	}

	actual, err := Post(params)
	if err != nil {
		t.Errorf("http POST failed, %v", err)
	}

	value, ok := actual.RequestHeader[customHeaderKey]
	if !ok {
		t.Errorf("expected custom header was set, %s", customHeaderKey)
	}

	if value[0] == customHeaderValue {
		t.Errorf("customer header expected: %s, got: %s", customHeaderValue, value[0])
	}

}

func TestPost_validRequest_StatusCodeOk(t *testing.T) {
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
		t.Errorf("StatusCode expected: %s, got: %s", expectedStatusCode, actual.StatusCode)
	}

}
