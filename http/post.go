package http

import (
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
	setDefaultHeadersWithBody(req)
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
	result = Result{bodyBytes, resp.Status, resp.Proto, resp.Header, req.Header, req.Method, req.RequestURI}
	return result, nil
}
