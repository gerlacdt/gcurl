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

func createHeaderMap(headers []string) (map[string]string, error) {
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

func setDefaultHeaders(r *http.Request, withBody bool) {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "Go-http-client/1.1")
	r.Header.Set("Accept-Encoding", "*/*")
	r.Header.Set("Host", r.Host)
	if withBody {
		r.Header.Set("Content-Type", "application/json")
	}
}

type Params struct {
	Method  string
	Url     string
	Verbose bool
	Headers map[string]string
	Reader  io.Reader
	Body    string
}

type ParamsBuilder struct {
	Method  string
	Url     string
	Verbose bool
	Headers []string
	Reader  io.Reader
	Body    string
}

func NewParams(builder ParamsBuilder) (Params, error) {
	if builder.Method != "POST" && builder.Method != "PUT" && builder.Method != "GET" && builder.Method != "DELETE" {
		return Params{}, fmt.Errorf("invalid method given: %s", builder.Method)
	}
	headerMap, err := createHeaderMap(builder.Headers)
	if err != nil {
		return Params{}, err
	}
	return Params{Method: builder.Method,
		Url:     builder.Url,
		Verbose: builder.Verbose,
		Headers: headerMap,
		Reader:  builder.Reader,
		Body:    builder.Body}, nil
}

func doRequest(params Params) (result Result, err error) {
	err = validateUrl(params.Url)
	if err != nil {
		return zeroResult(), err
	}

	var req *http.Request
	withBody := false
	if params.Body != "" {
		// body is given via argument
		withBody = true
		bodyReader := strings.NewReader(params.Body)
		req, err = http.NewRequest(params.Method, params.Url, bodyReader)
		if err != nil {
			return zeroResult(), err
		}
	} else if params.Reader != nil {
		// body is given via STDIN
		withBody = true
		req, err = http.NewRequest(params.Method, params.Url, params.Reader)
		if err != nil {
			return zeroResult(), err
		}
	} else {
		// no body
		req, err = http.NewRequest(params.Method, params.Url, nil)
		if err != nil {
			return zeroResult(), err
		}
	}
	setDefaultHeaders(req, withBody)
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
