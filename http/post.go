package http

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getHeaderMap(headers []string) (map[string]string, error) {
	headerMap := make(map[string]string)
	var err error

	for _, h := range headers {
		splitArr := strings.Split(h, ":")
		if len(splitArr) != 2 {
			return nil, fmt.Errorf("given invalid header, got: %s", h)
		}
		headerKey := splitArr[0]
		headerValue := splitArr[1]
		headerMap[headerKey] = headerValue
	}

	return headerMap, err
}

func Post(url string, headers []string) (string, error) {
	err := validateUrl(url)
	if err != nil {
		return "", err
	}

	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return "", err
	}

	contentType, ok := headerMap["Content-Type"]
	if !ok {
		contentType = "application/json" // default to json
	}

	resp, err := http.Post(url, contentType, os.Stdin)
	if err != nil {
		return "", nil
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
