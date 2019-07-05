package sfhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Response is the type that all the HTTP requests are going to return.
type Response struct {
	Code    int
	Body    interface{}
	Headers http.Header
}

//Request is the struct which defines am http request
type Request struct {
	URL     string
	Headers map[string]string
	Body    []byte
}

// Get executes a GET request against the specified URL.
func Get(r Request, out interface{}) (*Response, error) {
	return do("GET", r, out)
}

// Post executes a POST request against the specified URL.
func Post(r Request, out interface{}) (*Response, error) {
	return do("POST", r, out)
}

// Patch executes a PATCH request against the specified URL.
func Patch(r Request, out interface{}) (*Response, error) {
	return do("PATCH", r, out)
}

// Put executes a PUT request against the specified URL.
func Put(r Request, out interface{}) (*Response, error) {
	return do("PUT", r, out)
}

func do(method string, request Request, out interface{}) (*Response, error) {

	// Build the request
	req, err := http.NewRequest(method, request.URL, bytes.NewBuffer(request.Body))
	if err != nil {
		return nil, fmt.Errorf("error building the %s request for %s: %v", method, request.URL, err)
	}

	// Fill the request with the headers
	for k, v := range request.Headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout: time.Second * time.Duration(10),
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error doing %s to %s: %v", method, request.URL, err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading the response when doiing %s to %s: %v", method, request.URL, err)
	}

	// Initialising the response
	var response Response
	if out != nil {
		if json.Unmarshal(respBytes, out) != nil {
			return nil, fmt.Errorf("error decoding the response when doiing %s to %s: %v", method, request.URL, err)
		}
		response.Body = out
	} else {
		// If out hasn't been defined, we return the response body as bytes
		response.Body = respBytes
	}

	response.Code = resp.StatusCode
	response.Headers = resp.Header

	return &response, nil
}
