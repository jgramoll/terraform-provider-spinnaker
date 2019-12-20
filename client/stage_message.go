package client

// TODO FRD

// StageCompleteKey for stage complete
const StageCompleteKey = "stage.complete"

// StageFailedKey for stage failed
const StageFailedKey = "stage.failed"

// StageStartingKey for pipeline starting
const StageStartingKey = "stage.starting"

// ManualJudgmentKey for manual judgement
const ManualJudgmentKey = "manualJudgment"

// ManualJudgmentContinueKey for manual judgment continue
const ManualJudgmentContinueKey = "manualJudgmentContinue"

// ManualJudgmentStopKey for manual judgement stop
const ManualJudgmentStopKey = "manualJudgmentStop"

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

	ManualJudgmentContinue *MessageText `json:"manualJudgmentContinue"`
	ManualJudgmentStop     *MessageText `json:"manualJudgmentStop"`
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

// SetManualJudgmentContinueText for Message interface
func (m *StageMessage) SetManualJudgmentContinueText(text string) {
	m.ManualJudgmentContinue = &MessageText{Text: text}
}

// ManualJudgmentContinueText for Message interface
func (m *StageMessage) ManualJudgmentContinueText() string {
	if m.ManualJudgmentContinue != nil {
		return m.ManualJudgmentContinue.Text
	}
	return ""
}

// SetManualJudgmentStopText for Message interface
func (m *StageMessage) SetManualJudgmentStopText(text string) {
	m.ManualJudgmentStop = &MessageText{Text: text}
}

// ManualJudgmentStopText for Message interface
func (m *StageMessage) ManualJudgmentStopText() string {
	if m.ManualJudgmentStop != nil {
		return m.ManualJudgmentStop.Text
	}
	return ""
}
