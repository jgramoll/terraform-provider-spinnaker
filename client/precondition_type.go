package client

var preconditionFactories = map[PreconditionType]func(map[string]interface{}) (Precondition, error){}

// PreconditionType type of precondition
type PreconditionType string

func (st PreconditionType) String() string {
	return string(st)
}
