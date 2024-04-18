package http

import (
	"io"
	"net/http"
	"strings"
)

func Put(params PostParams) (result Result, err error) {
	err = validateUrl(params.Url)
	if err != nil {
		return zeroResult(), err
	}

	var req *http.Request
	if params.Body == "" {
		req, err = http.NewRequest("PUT", params.Url, params.Reader)
		if err != nil {
			return zeroResult(), err
		}
	} else {
		bodyReader := strings.NewReader(params.Body)
		req, err = http.NewRequest("PUT", params.Url, bodyReader)
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
