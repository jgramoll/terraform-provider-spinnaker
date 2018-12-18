package client

var stageFactories = map[StageType]func() interface{}{}

// StageType type of stage
type StageType string

func (st StageType) String() string {
	return string(st)
}
