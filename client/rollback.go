package client

type Rollback struct {
	OnFailure bool `json:"onFailure"`
}

func NewRollback() *Rollback {
	return &Rollback{
		OnFailure: false,
	}
}
