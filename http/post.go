package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type PostParams struct {
	Url     string
	Verbose bool
	Headers []string
	Reader  io.Reader
}

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

func setDefaultHeadersPost(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "Go-http-client/1.1")
	r.Header.Set("Accept-Encoding", "gzip")
	r.Header.Set("Host", r.Host)
	r.Header.Set("Content-Type", "application/json")
}

func Post(params PostParams) (result Result, err error) {
	err = validateUrl(params.Url)
	if err != nil {
		return zeroResult(), err
	}

	headerMap, err := getHeaderMap(params.Headers)
	if err != nil {
		return zeroResult(), err
	}

	req, err := http.NewRequest("POST", params.Url, params.Reader)
	if err != nil {
		return zeroResult(), err
	}
	setDefaultHeadersPost(req)
	for headerKey, headerValue := range headerMap {
		req.Header.Set(headerKey, headerValue)
	}
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
