package client

// Task a job within an application
type Task struct {
	Job         *[]*Job `json:"job"`
	Application string  `json:"application"`
	Description string  `json:"description"`
}
