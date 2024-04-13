package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type GetParams struct {
	Url     string
	Verbose bool
	Headers map[string]string
}

func NewGetParams(url string, verbose bool, headers []string) (GetParams, error) {
	headerMap, err := getHeaderMap(headers)
	if err != nil {
		return GetParams{}, err
	}
	return GetParams{Url: url, Verbose: verbose, Headers: headerMap}, nil
}

type Result struct {
	Body          []byte
	StatusCode    string
	Proto         string
	Header        map[string][]string
	RequestHeader map[string][]string
}

func (r *Result) Print(verbose bool) {
	if verbose {
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
	return Result{make([]byte, 0), "", "", m, m}
}

func validateUrl(givenUrl string) error {
	_, err := url.ParseRequestURI(givenUrl)
	return err
}

func setDefaultHeaders(r *http.Request) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "Go-http-client/1.1")
	r.Header.Set("Accept-Encoding", "gzip")
	r.Header.Set("Host", r.Host)
}

func Get(params GetParams) (response Result, err error) {
	err = validateUrl(params.Url)
	if err != nil {
		return zeroResult(), err
	}
	req, err := http.NewRequest("GET", params.Url, nil)
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

	response = Result{bodyBytes, resp.Status, resp.Proto, resp.Header, req.Header}
	return response, nil
}
