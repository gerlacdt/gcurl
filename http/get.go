package http

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) error {

	resp, err := http.Get(url)
	if err != err {
		return err
	}
	defer resp.Body.Close()
	// output to stdout
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s", body)
	return nil
}
