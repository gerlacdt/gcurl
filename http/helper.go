package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Result struct {
	Body          []byte
	StatusCode    string
	Proto         string
	Header        map[string][]string
	RequestHeader map[string][]string
	RequestMethod string
	RequestUri    string
}

func (r *Result) Print(verbose bool) {
	if verbose {
		fmt.Fprintf(os.Stderr, "> %s %s\n", r.RequestMethod, r.RequestUri)
		for reqHeader, reqHeaderValue := range r.RequestHeader {
			fmt.Fprintf(os.Stderr, "> %s : %s\n", reqHeader, strings.Join(reqHeaderValue, ","))
		}
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "< %s %s\n", r.Proto, r.StatusCode)
		for respHeader, respHeaderValue := range r.Header {
			fmt.Fprintf(os.Stderr, "< %s : %s\n", respHeader, strings.Join(respHeaderValue, ","))
		}
	}
	fmt.Printf("%s", r.Body)
}

func zeroResult() Result {
	m := make(map[string][]string)
	return Result{make([]byte, 0), "", "", m, m, "", ""}
}

func validateUrl(givenUrl string) error {
	_, err := url.ParseRequestURI(givenUrl)
	return err
}

func setDefaultHeaders(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "Go-http-client/1.1")
	r.Header.Set("Accept-Encoding", "*/*")
	r.Header.Set("Host", r.Host)
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

func setDefaultHeadersWithBody(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "Go-http-client/1.1")
	r.Header.Set("Accept-Encoding", "*/*")
	r.Header.Set("Host", r.Host)
	r.Header.Set("Content-Type", "application/json")
}

type Params struct {
	Method  string
	Url     string
	Verbose bool
	Headers map[string]string
}

func NewParams(method string, url string, verbose bool, headers []string) (Params, error) {
	if method != "GET" && method != "DELETE" {
		return Params{}, fmt.Errorf("invalid method given: %s", method)
	}
	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return Params{}, err
	}
	return Params{Method: method, Url: url, Verbose: verbose, Headers: headerMap}, nil
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
