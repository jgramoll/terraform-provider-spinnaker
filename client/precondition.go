package client

type Precondition interface {
	GetType() PreconditionType
}

type BasePrecondition struct {
	FailPipeline bool             `json:"failPipeline"`
	Type         PreconditionType `json:"type"`
}

func NewBasePrecondition(t PreconditionType) *BasePrecondition {
	return &BasePrecondition{
		FailPipeline: true,
		Type:         t,
	}
}

func (p *BasePrecondition) GetType() PreconditionType {
	return p.Type
}

// func (stage *BaseStage) parseBasePrecondition(preconditionType map[string]interface{}) error {
// 	notifications, err := parseNotifications(stageMap["notifications"])
// 	if err != nil {
// 		return err
// 	}
// 	stage.Notifications = notifications
// 	delete(stageMap, "notifications")

// 	return nil
// }
