package client

// CanaryConfigJudge judge
type CanaryConfigJudge struct {
	Name                string                 `json:"name"`
	JudgeConfigurations map[string]interface{} `json:"judgeConfigurations"`
}

// NewCanaryConfigJudge new judge
func NewCanaryConfigJudge(name string) *CanaryConfigJudge {
	return &CanaryConfigJudge{
		Name:                name,
		JudgeConfigurations: map[string]interface{}{},
	}
}
