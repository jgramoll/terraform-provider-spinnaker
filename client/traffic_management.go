package client

// TrafficManagement traffic
type TrafficManagement struct {
	Enabled bool                      `json:"enabled"`
	Options *TrafficManagementOptions `json:"options"`
}

// NewTrafficManagement new traffic
func NewTrafficManagement() *TrafficManagement {
	return &TrafficManagement{
		Enabled: false,
		Options: NewTrafficManagementOptions(),
	}
}
