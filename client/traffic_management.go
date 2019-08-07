package client

type TrafficManagement struct {
	Enabled bool                      `json:"enabled"`
	Options *TrafficManagementOptions `json:"options"`
}

func NewTrafficManagement() *TrafficManagement {
	return &TrafficManagement{
		Enabled: false,
		Options: NewTrafficManagementOptions(),
	}
}
