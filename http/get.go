package http

import (
	"io"
	"net/http"
	"net/url"
)

func validateUrl(givenUrl string) error {
	_, err := url.ParseRequestURI(givenUrl)
	return err
}

func Get(url string) (string, error) {
	err := validateUrl(url)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(url)
	if err != err {
		return "", err
	}
	defer resp.Body.Close()
	// output to stdout
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
