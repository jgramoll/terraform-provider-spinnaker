package client

import (
	"encoding/json"
	"fmt"
	"reflect"
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

// UnmarshalJSON unmarshal
func (e *SpinnakerError) UnmarshalJSON(bytes []byte) error {
	var errorMap map[string]interface{}
	if err := json.Unmarshal(bytes, &errorMap); err != nil {
		return err
	}

	errorMsg, ok := errorMap["error"].(string)
	if ok {
		e.ErrorMsg = errorMsg
	}
	exception, ok := errorMap["exception"].(string)
	if ok {
		e.Exception = exception
	}
	message, ok := errorMap["message"].(string)
	if ok {
		e.Message = message
	}
	status, ok := errorMap["status"].(float64)
	if ok {
		e.Status = int(status)
	}
	body, ok := errorMap["body"].(string)
	if ok {
		e.Body = body
	}

	timestampInterface := errorMap["timestamp"]
	timestampType := reflect.TypeOf(timestampInterface).String()
	switch timestampType {
	case "string":
		e.Timestamp = timestampInterface.(string)
	case "float64":
		e.Timestamp = fmt.Sprintf("%.0f", timestampInterface.(float64))
	default:
		return fmt.Errorf("Unknown timestamp type: %v", timestampType)
	}
	return nil
}

func (e *SpinnakerError) Error() string {
	return fmt.Sprintf("%d %v: %v%v\n%v", e.Status, e.ErrorMsg, e.Message,
		e.Body, e.Exception)
}
