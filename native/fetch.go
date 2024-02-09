package native

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const USER_AGENT = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0"

type FetchOpts struct {
	Headers map[string]string
	Method  string
	Data    string
	Verbose bool
}

type FetchResponse struct {
	body []byte
}

func (f *FetchResponse) ToString() string {
	return string(f.body)
}

func (f *FetchResponse) JSON() (map[string]any, error) {
	var jsonRes map[string]any

	err := json.Unmarshal(f.body, &jsonRes)

	return jsonRes, err
}

func fetchBody(url string, opts *FetchOpts) (io.ReadCloser, error) {
	client := &http.Client{}
	method := "GET"
	var body io.Reader = nil

	if opts != nil {
		method = opts.Method

		if method == "POST" {
			body = strings.NewReader(opts.Data)
		}
	}

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("User-Agent", USER_AGENT)

	if opts != nil {
		if opts.Headers != nil {
			for k, v := range opts.Headers {
				req.Header.Set(k, v)
			}
		}
	}

	if opts.Verbose {
		fmt.Println(method, url, opts.Headers, opts)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

// https://zetcode.com/golang/getpostrequest/
func Fetch(url string, opts *FetchOpts) (*FetchResponse, error) {
	resBody, err := fetchBody(url, opts)

	if err != nil {
		return nil, err
	}

	defer resBody.Close()

	body, err := io.ReadAll(resBody)

	if err != nil {
		return nil, err
	}

	response := &FetchResponse{body}

	return response, nil
}
