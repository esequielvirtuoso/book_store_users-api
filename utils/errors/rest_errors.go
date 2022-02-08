// Package errors must be used while handling errors in REST applications.
package errors

import (
	"errors"
	"net/http"
)

// RestErr is a standard struct to be used while handling errors in REST applications.
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewError returns a new error with an input message
func NewError(msg string) error {
	return errors.New(msg)
}

/* NewBadRequestError returns a standardized struct with the correct status,
*  and error tag for bad request situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

/* NewNotFoundError returns a standardized struct with the correct status,
*  and error tag for not found situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

/* NewInternalServerError returns a standardized struct with the correct status,
*  and error tag for internal error situations
*  Args:
*  message (string): The message to be assigned to the struct's Message field
 */
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
