package client

import (
	"fmt"
)

type SpinnakerError struct {
	ErrorMsg  string `json:"error"`
	Exception string
	Message   string
	Status    int
	Timestamp int64
	Body      string
}

func (r *SpinnakerError) Error() string {
	return fmt.Sprintf("%d %v: %v%v\n%v", r.Status, r.ErrorMsg, r.Message,
		r.Body, r.Exception)
}
