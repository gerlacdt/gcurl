package http

import (
	"fmt"
	"strings"
	"testing"
)

func TestDelete_validRequest_statusCodeOk(t *testing.T) {
	client := NewClient()
	method := "DELETE"
	url := "http://localhost:8080/delete"
	verbose := false
	headers := make([]string, 0)

	params, err := NewParams(ParamsBuilder{
		Method: method, Url: url, Verbose: verbose, Headers: headers, Reader: nil, Body: ""})
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := client.Delete(params)
	if err != nil {
		t.Errorf("http DELETE failed, %v", err)
	}

	expectedStatusCode := "200"
	if !strings.Contains(actual.StatusCode, expectedStatusCode) {
		t.Errorf("expected statusCode: %s, got: %s", expectedStatusCode, actual.StatusCode)
	}
}

func TestDelete_validRequest_customHeaderSet(t *testing.T) {
	client := NewClient()
	method := "DELETE"
	url := "http://localhost:8080/delete"
	verbose := false
	headers := make([]string, 0)
	customHeaderKey := "X-Custom"
	customHeaderValue := "foobar"
	headers = append(headers, fmt.Sprintf("%s: %s", customHeaderKey, customHeaderValue))

	params, err := NewParams(ParamsBuilder{
		Method: method, Url: url, Verbose: verbose, Headers: headers, Reader: nil, Body: ""})
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := client.Delete(params)
	if err != nil {
		t.Errorf("http DELETE failed, %v", err)
	}

	value, ok := actual.RequestHeader[customHeaderKey]
	if !ok {
		t.Errorf("expected customHeader was set, %s", customHeaderKey)
	}

	if value[0] == customHeaderValue {
		t.Errorf("customHeaderValue, expected %s, got %s", customHeaderValue, value)
	}
}
