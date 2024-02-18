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

func Get(url string, verbose bool) (string, error) {
	err := validateUrl(url)
	if err != nil {
		return "", err
	}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
