package http

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
)

func testGetEndpoint(req *http.Request) (*http.Response, error) {
	resp, err := httpmock.NewJsonResponse(200, map[string]interface{}{
		"foo": "bar",
	})
	resp.Header.Add("key", "value")
	// We check we have received the correct headers
	if req.Header.Get("foo") == "bar" {
		return resp, err
	}

	return nil, fmt.Errorf("Wrong input parameters")
}

func Test_ResponseWithDefinedObject(t *testing.T) {
	// Define http mock results
	httpmock.Activate()
	defer httpmock.Deactivate()
	url := "http://test:8080/test"
	httpmock.RegisterResponder(
		"GET",
		url,
		testGetEndpoint,
	)

	// GIVEN we want to have an specific object filled as an output
	out := struct {
		Foo string `json:"foo"`
	}{}

	// AND the request is going to have specific headers
	headers := map[string]string{
		"foo": "bar",
	}

	// WHEN we make the HTTP request
	resp, err := Get(Request{
		URL:     url,
		Headers: headers,
	}, &out)

	// THEN we shouldn't get any error
	if err != nil {
		t.Errorf("Got an error making the request: %v", err)
		return
	}

	// AND the response code should be 200
	if resp.Code != http.StatusOK {
		t.Errorf("Unexpected status code, got %v", resp.Code)
	}

	// AND the outcome object should be filled
	if out.Foo != "bar" {
		t.Errorf("Unexpected status code, got %v", resp.Body)
	}

	// AND we should be able to get the response headers
	if resp.Headers.Get("key") != "value" {
		t.Errorf("Unexpected header value: %v", resp.Headers["key"][0])
	}
}

func Test_ResponseWithUndefinedObject(t *testing.T) {
	// Define http mock results
	httpmock.Activate()
	defer httpmock.Deactivate()
	url := "http://test:8080/test"
	httpmock.RegisterResponder(
		"GET",
		url,
		testGetEndpoint,
	)

	// GIVEN the request is going to have specific headers
	headers := map[string]string{
		"foo": "bar",
	}

	// WHEN we make the HTTP request without defining the output
	resp, err := Get(Request{
		URL:     url,
		Headers: headers,
	}, nil)

	// THEN we shouldn't get any error
	if err != nil {
		t.Errorf("Got an error making the request: %v", err)
		return
	}

	// AND the response code should be 200
	if resp.Code != http.StatusOK {
		t.Errorf("Unexpected status code, got %v", resp.Code)
	}

	// AND the outcome object should be filled
	if string(resp.Body.([]byte)) != `{"foo":"bar"}` {
		t.Errorf("Unexpected body value: %v", string(resp.Body.([]byte)))
	}

	// AND we should be able to get the response headers
	if resp.Headers.Get("key") != "value" {
		t.Errorf("Unexpected header value: %v", resp.Headers["key"][0])
	}
}
