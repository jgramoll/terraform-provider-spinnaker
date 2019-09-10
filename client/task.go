package client

// Task a job within an application
type Task struct {
	Job         *[]*Job `json:"job"`
	Application string  `json:"application"`
	Description string  `json:"description"`
}

// TaskResponse ref for task execution
type TaskResponse struct {
	Ref string `json:"ref"`
}

// TaskExecution get status of task
type TaskExecution struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Application string `json:"application"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	BuildTime   int64  `json:"buildTime"`
	Status      string `json:"status"`
}
