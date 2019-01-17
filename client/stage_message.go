package client

// TODO FRD

// StageCompleteKey for stage complete
const StageCompleteKey = "stage.complete"

// StageFailedKey for stage failed
const StageFailedKey = "stage.failed"

// StageStartingKey for pipeline starting
const StageStartingKey = "stage.starting"

func init() {
	messageFactories[NotificationLevelStage] = func() Message {
		return &StageMessage{}
	}
}

// StageMessage for Stage Notification
type StageMessage struct {
	Complete *MessageText `json:"stage.complete" mapstructure:"stage.complete"`
	Failed   *MessageText `json:"stage.failed" mapstructure:"stage.failed"`
	Starting *MessageText `json:"stage.starting" mapstructure:"stage.starting"`
}

// SetCompleteText for Message interface
func (m *StageMessage) SetCompleteText(text string) {
	m.Complete = &MessageText{Text: text}
}

// CompleteText for Message interface
func (m *StageMessage) CompleteText() string {
	if m.Complete != nil {
		return m.Complete.Text
	}
	return ""
}

// SetFailedText for Message interface
func (m *StageMessage) SetFailedText(text string) {
	m.Failed = &MessageText{Text: text}
}

// FailedText for Message interface
func (m *StageMessage) FailedText() string {
	if m.Failed != nil {
		return m.Failed.Text
	}
	return ""
}

// SetStartingText for Message interface
func (m *StageMessage) SetStartingText(text string) {
	m.Starting = &MessageText{Text: text}
}

// StartingText for Message interface
func (m *StageMessage) StartingText() string {
	if m.Starting != nil {
		return m.Starting.Text
	}
	return ""
}
