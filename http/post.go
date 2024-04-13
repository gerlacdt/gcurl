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
	Headers map[string]string
	Reader  io.Reader
	Body    string
}

func NewPostParams(url string, verbose bool, headers []string, reader io.Reader, body string) (PostParams, error) {
	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return PostParams{}, err
	}
	return PostParams{Url: url, Verbose: verbose, Headers: headerMap, Reader: reader, Body: body}, nil
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

	var req *http.Request
	if params.Body == "" {
		req, err = http.NewRequest("POST", params.Url, params.Reader)
		if err != nil {
			return zeroResult(), err
		}
	} else {
		bodyReader := strings.NewReader(params.Body)
		req, err = http.NewRequest("POST", params.Url, bodyReader)
		if err != nil {
			return zeroResult(), err
		}
	}
	setDefaultHeadersPost(req)
	for headerKey, headerValue := range params.Headers {
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
