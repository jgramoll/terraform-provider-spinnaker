package client

// PipelineCompleteKey for pipeline complete
const PipelineCompleteKey = "pipeline.complete"

// PipelineFailedKey for pipeline failed
const PipelineFailedKey = "pipeline.failed"

// PipelineStartingKey for pipeline starting
const PipelineStartingKey = "pipeline.starting"

func init() {
	messageFactories[NotificationLevelPipeline] = func() interface{} {
		return &PipelineMessage{}
	}
}

// PipelineMessage for Pipeline Notification
type PipelineMessage struct {
	Complete *MessageText `json:"pipeline.complete" mapstructure:"pipeline.complete"`
	Failed   *MessageText `json:"pipeline.failed" mapstructure:"pipeline.failed"`
	Starting *MessageText `json:"pipeline.starting" mapstructure:"pipeline.starting"`
}

// SetCompleteText for Message interface
func (m *PipelineMessage) SetCompleteText(text string) {
	m.Complete = &MessageText{Text: text}
}

// CompleteText for Message interface
func (m *PipelineMessage) CompleteText() string {
	if m.Complete != nil {
		return m.Complete.Text
	}
	return ""
}

// SetFailedText for Message interface
func (m *PipelineMessage) SetFailedText(text string) {
	m.Failed = &MessageText{Text: text}
}

// FailedText for Message interface
func (m *PipelineMessage) FailedText() string {
	if m.Failed != nil {
		return m.Failed.Text
	}
	return ""
}

// SetStartingText for Message interface
func (m *PipelineMessage) SetStartingText(text string) {
	m.Starting = &MessageText{Text: text}
}

// StartingText for Message interface
func (m *PipelineMessage) StartingText() string {
	if m.Starting != nil {
		return m.Starting.Text
	}
	return ""
}
