package main

import (
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// HTTP helper methods for sending authenticated requests to Brizo
// ---------------------------------------------------------------

// HTTPGet sends a GET request
func HTTPGet(path string) (string, error) {
	request := gorequest.New().Get(buildURL(path))
	requestHeaders(request)
	response, body, errs := request.End()

	return handleResponse(response, body, errs)
}

// HTTPPost sends a POST request
func HTTPPost(path string) (string, error) {
	request := gorequest.New().Post(buildURL(path))
	requestHeaders(request)
	response, body, errs := request.End()

	return handleResponse(response, body, errs)
}

func handleResponse(response gorequest.Response, body string, errs []error) (string, error) {
	if response.StatusCode == 401 {
		return "", errors.New("Unauthorized")
	}

	if len(errs) != 0 {
		fmt.Println("Error from API")
		return "", errs[0]
	}

	return body, nil
}

// requestHeaders configures default auth and content headers for a request
func requestHeaders(request *gorequest.SuperAgent) *gorequest.SuperAgent {
	return request.
		Set("Authorization", "Bearer "+Config.Token).
		Set("Content-Type", "application/json")
}

// buildURL appends the provided path to the configured endpoint
func buildURL(path string) string {
	return Config.Endpoint + path
}
