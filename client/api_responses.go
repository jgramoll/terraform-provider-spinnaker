package client

import (
	"fmt"
)

type SpinnakerError struct {
	StatusCode int    `json:"statuscode"`
	StatusDesc string `json:"statusdesc"`
	Message    string `json:"errormessage"`
}

func (r *SpinnakerError) Error() string {
	return fmt.Sprintf("%d %v: %v", r.StatusCode, r.StatusDesc, r.Message)
}

type errorJsonResponse struct {
	Error *SpinnakerError `json:"error"`
}
