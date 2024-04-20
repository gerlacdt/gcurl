package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ParamsWithBody struct {
	Params
	Reader io.Reader
	Body   string
}

func NewParamsWithBody(method string, url string, verbose bool, headers []string, reader io.Reader, body string) (ParamsWithBody, error) {
	if method != "POST" && method != "PUT" {
		return ParamsWithBody{}, fmt.Errorf("invalid method given: %s", method)
	}
	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return ParamsWithBody{}, err
	}
	return ParamsWithBody{Params: Params{Method: method,
		Url:     url,
		Verbose: verbose,
		Headers: headerMap,
	},
		Reader: reader,
		Body:   body}, nil
}

func Post(params ParamsWithBody) (result Result, err error) {
	return requestWithBody(params)
}

func requestWithBody(params ParamsWithBody) (result Result, err error) {
	err = validateUrl(params.Url)
	if err != nil {
		return zeroResult(), err
	}

	var req *http.Request
	if params.Body == "" {
		req, err = http.NewRequest(params.Method, params.Url, params.Reader)
		if err != nil {
			return zeroResult(), err
		}
	} else {
		bodyReader := strings.NewReader(params.Body)
		req, err = http.NewRequest(params.Method, params.Url, bodyReader)
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
	result = Result{bodyBytes, resp.Status, resp.Proto, resp.Header, req.Header, req.Method, req.URL.EscapedPath()}
	return result, nil
}
