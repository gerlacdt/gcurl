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
	Header        map[string][]string
	RequestHeader map[string][]string
}

func (r *Result) Print(verbose bool) {
	if verbose {
		// print request headers
		for reqHeader, reqHeaderValue := range r.RequestHeader {
			fmt.Fprintf(os.Stderr, "%s : %s\n", reqHeader, strings.Join(reqHeaderValue, ","))
		}
		fmt.Fprintf(os.Stderr, "\n")
		// print response headers
		for respHeader, respHeaderValue := range r.Header {
			fmt.Fprintf(os.Stderr, "%s : %s\n", respHeader, strings.Join(respHeaderValue, ","))
		}
	}
	fmt.Printf("%s", r.Body)
}

func zeroResult() Result {
	m := make(map[string][]string)
	return Result{make([]byte, 0), m, m}
}

func validateUrl(givenUrl string) error {
	_, err := url.ParseRequestURI(givenUrl)
	return err
}

func Get(url string, verbose bool) (response Result, err error) {
	err = validateUrl(url)
	if err != nil {
		return zeroResult(), err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return zeroResult(), err
	}
	req.Header.Add("Accept", "application/json")
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

	response = Result{bodyBytes, resp.Header, req.Header}
	return response, nil
}
