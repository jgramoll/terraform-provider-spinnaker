package client

type preconditionFactoryFunc func(map[string]interface{}) (Precondition, error)

var preconditionFactory = map[PreconditionType]preconditionFactoryFunc{}

// PreconditionType type of precondition
type PreconditionType string

func (st PreconditionType) String() string {
	return string(st)
}
