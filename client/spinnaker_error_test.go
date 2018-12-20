package client

import (
	"testing"
)

func TestSpinnakerError(t *testing.T) {
	err := error(&SpinnakerError{})
	if err.Error() == "" {
		t.Fatal("Should be an error")
	}
}
