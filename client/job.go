package client

// Job a job to run
type Job struct {
	Type        string       `json:"type"`
	Application *Application `json:"application"`
	User        string       `json:"user"`
}
