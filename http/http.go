package http

import "net/http"

// Get gets
func Get(url string) (*http.Response, error) {
	return request("GET", url)
}

func request(method, url string) (*http.Response, error) {
	return nil, nil
}
