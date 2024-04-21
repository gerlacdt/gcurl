package http

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
)

var update = flag.Bool("update", false, "update golden files")

func TestGet_validRequest_statusCodeOk(t *testing.T) {
	client := NewClient()
	method := "GET"
	url := "http://localhost:8080/get"
	verbose := false
	headers := make([]string, 0)

	params, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := client.Get(params)
	if err != nil {
		t.Errorf("http GET failed, %v", err)
	}

	expectedStatusCode := "200"
	if !strings.Contains(actual.StatusCode, expectedStatusCode) {
		t.Errorf("expected statusCode: %s, got: %s", expectedStatusCode, actual.StatusCode)
	}
}

func TestGet_validRequest_customHeaderSet(t *testing.T) {
	client := NewClient()
	method := "GET"
	url := "http://localhost:8080/get"
	verbose := false
	headers := make([]string, 0)
	customHeaderKey := "X-Custom"
	customHeaderValue := "foobar"
	headers = append(headers, fmt.Sprintf("%s: %s", customHeaderKey, customHeaderValue))

	params, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := client.Get(params)
	if err != nil {
		t.Errorf("http GET failed, %v", err)
	}

	value, ok := actual.RequestHeader[customHeaderKey]
	if !ok {
		t.Errorf("expected customHeader was set, %s", customHeaderKey)
	}

	if value[0] == customHeaderValue {
		t.Errorf("customHeaderValue, expected %s, got %s", customHeaderValue, value)
	}
}

func TestGet_validRequest_bodyOk(t *testing.T) {
	client := NewClient()
	method := "GET"
	url := "http://localhost:8080/get"
	verbose := false
	headers := make([]string, 0)

	params, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := client.Get(params)
	if err != nil {
		t.Errorf("http GET failed, %v", err)
	}

	body := string(actual.Body)
	if !strings.Contains(body, "origin") {
		t.Errorf("expected origin, got: %v", actual)
	}
	if !strings.Contains(body, "User-Agent") {
		t.Errorf("expected User-Agent, got: %v", actual)
	}
}

func TestGet_nonExistingUrl_statusCodeNotFound(t *testing.T) {
	client := NewClient()
	method := "GET"
	url := "http://localhost:8080/notexist"
	verbose := false
	headers := make([]string, 0)

	params, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	actual, err := client.Get(params)
	if err != nil {
		t.Errorf("http GET failed, %v", err)
	}

	expectedStatusCode := "404"
	if !strings.Contains(actual.StatusCode, expectedStatusCode) {
		t.Errorf("expected statusCode: %s, got: %s", expectedStatusCode, actual.StatusCode)
	}
}

func TestGet_invalidHeader_fail(t *testing.T) {
	method := "GET"
	url := "http://localhost:8080/get"
	verbose := false
	headers := make([]string, 1)
	headers = append(headers, "X-Custom foobar") // missing semicolon delimiter

	_, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err == nil {
		t.Errorf("GetParams should fail but error was nil, %v", err)
	}
}

func TestGet_httpmock_ok(t *testing.T) {
	// arrange
	client := NewClient()
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles",
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))

	// act
	method := "GET"
	url := "https://api.mybiz.com/articles"
	verbose := false
	headers := make([]string, 0)
	params, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err != nil {
		t.Errorf("GETParams creation failed, %v", err)
	}
	response, err := client.Get(params)
	if err != nil {
		t.Errorf("http GET request failed: %v", err)
	}

	// assert
	callCount := httpmock.GetTotalCallCount()
	expected := 1
	if callCount != expected {
		t.Errorf("number of calls expected: %d, got: %d", expected, callCount)
	}

	expectedBody := "Great Article"
	body := string(response.Body)
	if !strings.Contains(body, expectedBody) {
		t.Errorf("body contains expected: %s, got: %s", expectedBody, body)
	}
}

// golden files reference Mitchell Hashimoto slides
// https://speakerdeck.com/mitchellh/advanced-testing-with-go?slide=20
func TestGet_goldenFile_ok(t *testing.T) {
	// arrange
	client := NewClient()
	method := "GET"
	url := "http://localhost:8080/get"
	verbose := false
	headers := make([]string, 0)
	params, err := NewParamsWithBody(method, url, verbose, headers, nil, "")
	if err != nil {
		t.Errorf("GetParams creation failed, %v", err)
	}

	// act
	actual, err := client.Get(params)
	if err != nil {
		t.Errorf("http GET failed, %v", err)
	}

	// assert
	golden := filepath.Join("test-fixtures", "get_body.golden")
	if *update {
		err = os.WriteFile(golden, actual.Body, 0644)
		if err != nil {
			t.Errorf("Failed updating golden file: %v", err)
		}
	}
	expected, err := os.ReadFile(golden)
	if err != nil {
		t.Errorf("Failed reading golden file, %v", err)
	}
	if !bytes.Equal(actual.Body, expected) {
		t.Errorf("body expected golden: %s, got: %s", expected, actual.Body)
	}
}
