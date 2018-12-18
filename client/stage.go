package client

// Stage interface for Pipeline stages
type Stage interface {
	GetName() string
	GetType() StageType
}

// BaseStage attributes common to all Pipeline stages
type BaseStage struct {
	Name  string    `json:"name"`
	RefID string    `json:"refId"`
	Type  StageType `json:"type"`
}
