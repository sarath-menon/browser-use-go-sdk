package browseruse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIError represents an error returned by the Browser Use API
type APIError struct {
	StatusCode int
	Message    string
	Body       string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Body)
}

// IsNotFound returns true if the error is a 404 Not Found error
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusNotFound
	}
	return false
}

// IsBadRequest returns true if the error is a 400 Bad Request error
func IsBadRequest(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusBadRequest
	}
	return false
}

// IsValidationError returns true if the error is a 422 Validation error
func IsValidationError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusUnprocessableEntity
	}
	return false
}

// IsRateLimitError returns true if the error is a 429 Rate Limit error
func IsRateLimitError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusTooManyRequests
	}
	return false
}

// parseError attempts to parse an error response from the API
func parseError(statusCode int, body []byte) error {
	apiErr := &APIError{
		StatusCode: statusCode,
		Body:       string(body),
	}

	// Try to extract a message from common error formats
	var errorResponse map[string]interface{}
	if err := json.Unmarshal(body, &errorResponse); err == nil {
		// Try common error message fields
		if msg, ok := errorResponse["message"].(string); ok {
			apiErr.Message = msg
		} else if msg, ok := errorResponse["error"].(string); ok {
			apiErr.Message = msg
		} else if detail, ok := errorResponse["detail"].(string); ok {
			apiErr.Message = detail
		}
	}

	// Add human-readable messages for specific status codes
	if apiErr.Message == "" {
		switch statusCode {
		case http.StatusBadRequest:
			apiErr.Message = "Session is stopped or has running task"
		case http.StatusNotFound:
			apiErr.Message = "Resource not found"
		case http.StatusUnprocessableEntity:
			apiErr.Message = "Request validation failed"
		case http.StatusTooManyRequests:
			apiErr.Message = "Too many concurrent active sessions"
		case http.StatusInternalServerError:
			apiErr.Message = "Internal server error"
		}
	}

	return apiErr
}
