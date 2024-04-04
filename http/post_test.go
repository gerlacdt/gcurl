package http

import (
	"fmt"
	"strings"
	"testing"
)

func TestPost(t *testing.T) {
	url := "http://localhost:8080/post"
	headers := make([]string, 0)

	reader := strings.NewReader("\"foo\": \"bar\"")
	actual, err := Post(url, headers, reader)
	if err != nil {
		t.Fatalf("http POST failed, %v", err)
	}

	// TODO add assertions
	fmt.Printf("%s", actual)
}
