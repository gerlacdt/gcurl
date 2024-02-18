package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func validateUrl(givenUrl string) error {
	_, err := url.ParseRequestURI(givenUrl)
	return err
}

func Get(url string, verbose bool) (body string, err error) {
	err = validateUrl(url)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		err = resp.Body.Close()
	}()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if verbose {
		// print request headers
		for reqHeader, reqHeaderValue := range req.Header {
			fmt.Fprintf(os.Stderr, "%s : %s\n", reqHeader, strings.Join(reqHeaderValue, ","))
		}
		fmt.Fprintf(os.Stderr, "\n")
		// print response headers
		for respHeader, respHeaderValue := range resp.Header {
			fmt.Fprintf(os.Stderr, "%s : %s\n", respHeader, strings.Join(respHeaderValue, ","))
		}
	}
	body = string(bodyBytes)
	return
}
