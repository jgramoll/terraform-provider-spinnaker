package client

import (
	"fmt"
)

// SpinnakerError Error response from spinnaker
type SpinnakerError struct {
	ErrorMsg  string `json:"error"`
	Exception string `json:"exception"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
	Timestamp string `json:"timestamp"`
	Body      string `json:"body"`
}

// For error interface
func (r *SpinnakerError) Error() string {
	return fmt.Sprintf("%d %v: %v%v\n%v", r.Status, r.ErrorMsg, r.Message,
		r.Body, r.Exception)
}
