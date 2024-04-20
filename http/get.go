package http

import (
	"io"
	"net/http"
)

type Params struct {
	Method  string
	Url     string
	Verbose bool
	Headers map[string]string
}

func NewParams(method string, url string, verbose bool, headers []string) (Params, error) {
	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return Params{}, err
	}
	return Params{Method: method, Url: url, Verbose: verbose, Headers: headerMap}, nil
}

func Get(params Params) (response Result, err error) {
	return request(params)
}

func request(params Params) (response Result, err error) {
	err = validateUrl(params.Url)
	if err != nil {
		return zeroResult(), err
	}
	req, err := http.NewRequest(params.Method, params.Url, nil)
	if err != nil {
		return zeroResult(), err
	}
	setDefaultHeaders(req)
	for headerKey, headerValue := range params.Headers {
		req.Header.Set(headerKey, headerValue)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return zeroResult(), err
	}
	defer func() {
		err = resp.Body.Close()
	}()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return zeroResult(), err
	}

	response = Result{bodyBytes, resp.Status, resp.Proto, resp.Header, req.Header, req.Method, req.URL.EscapedPath()}
	return response, nil
}
