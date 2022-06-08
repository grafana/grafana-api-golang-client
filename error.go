package gapi

import (
	"errors"
	"fmt"
)

type GApiError struct {
	statusCode int
	message    string
}

// Error implements the error interface.
func (e *GApiError) Error() string {
	return fmt.Sprintf("status: %d, body: %v", e.statusCode, e.message)
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *GApiError
	if errors.As(err, &e) && e.statusCode == 404 {
		return true
	}
	return false
}

func GetGApiError(err error) *GApiError {
	if err == nil {
		return nil
	}
	var e *GApiError
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func (e *GApiError) StatusCode() int {
	return e.statusCode
}

func (e *GApiError) Message() string {
	return e.message
}
