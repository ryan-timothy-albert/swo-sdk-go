// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"fmt"
	"net/http"
)

type APIError struct {
	Message     string
	StatusCode  int
	Body        string
	RawResponse *http.Response
}

var _ error = &APIError{}

func NewAPIError(message string, statusCode int, body string, httpRes *http.Response) *APIError {
	return &APIError{
		Message:     message,
		StatusCode:  statusCode,
		Body:        body,
		RawResponse: httpRes,
	}
}

func (e *APIError) Error() string {
	body := ""
	if len(e.Body) > 0 {
		body = fmt.Sprintf("\n%s", e.Body)
	}

	return fmt.Sprintf("%s: Status %d%s", e.Message, e.StatusCode, body)
}
