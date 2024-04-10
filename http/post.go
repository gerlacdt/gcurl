package http

import (
	"fmt"
	"io"
	"net/http"
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

func Post(url string, headers []string, reader io.Reader) (result Result, err error) {
	err = validateUrl(url)
	if err != nil {
		return zeroResult(), err
	}

	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return zeroResult(), err
	}

	contentType, ok := headerMap["Content-Type"]
	if !ok {
		contentType = "application/json" // default to json
	}

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return zeroResult(), err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", contentType)
	client := &http.Client{}
	resp, err := client.Do(req)

	// resp, err := http.Post(url, contentType, os.Stdin)
	if err != nil {
		return zeroResult(), nil
	}
	defer func() {
		err = resp.Body.Close()
	}()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return zeroResult(), err
	}
	result = Result{bodyBytes, resp.Status, resp.Proto, resp.Header, req.Header}
	return result, nil
}
