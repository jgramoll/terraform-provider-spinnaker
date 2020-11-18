package client

// Rollback rollback
type Rollback struct {
	OnFailure bool `json:"onFailure"`
}

// NewRollback new rollback
func NewRollback() *Rollback {
	return &Rollback{
		OnFailure: false,
	}
}
