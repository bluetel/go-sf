package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Response is the type that all the HTTP requests are going to return.
type Response struct {
	Code    int
	Body    interface{}
	Headers map[string][]string
}

// Get executes a GET request against the specified URL.
func Get(url string, headers map[string]string, out interface{}) (*Response, error) {
	return request("GET", url, headers, out, nil)
}

// Post executes a POST request against the specified URL.
func Post(url string, headers map[string]string, out, payload interface{}) (*Response, error) {
	return request("POST", url, headers, out, payload)
}

// Patch executes a PATCH request against the specified URL.
func Patch(url string, headers map[string]string, out, payload interface{}) (*Response, error) {
	return request("PATCH", url, headers, out, payload)
}

// Put executes a PUT request against the specified URL.
func Put(url string, headers map[string]string, out, payload interface{}) (*Response, error) {
	return request("PUT", url, headers, out, payload)
}

func request(method, url string, headers map[string]string, out, payload interface{}) (*Response, error) {
	var body io.Reader
	if payload != nil {

	}

	// Build the request
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Error building the %s request for %s: %v", method, url, err)
	}

	// Fill the request with the headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error doing %s to %s: %v", method, url, err)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading the response when doiing %s to %s: %v", method, url, err)
	}

	// Initialising the response
	var response Response
	if out != nil {
		errJSON := json.Unmarshal(respBytes, out)
		if errJSON != nil {
			return nil, fmt.Errorf("Error decoding the response when doiing %s to %s: %v", method, url, err)
		}
		response.Body = out
	} else {
		// If out hasn't been defined, we return the response body as an string
		response.Body = string(respBytes)
	}

	response.Code = resp.StatusCode
	response.Headers = resp.Header

	return &response, nil
}
