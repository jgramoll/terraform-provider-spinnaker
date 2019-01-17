package client

var stageFactories = map[StageType]func(map[string]interface{}) (Stage, error){}

// StageType type of stage
type StageType string

func (st StageType) String() string {
	return string(st)
}
