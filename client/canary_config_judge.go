package client

type CanaryConfigJudge struct {
	Name                string                 `json:"name"`
	JudgeConfigurations map[string]interface{} `json:"judgeConfigurations"`
}

func NewCanaryConfigJudge(name string) *CanaryConfigJudge {
	return &CanaryConfigJudge{
		Name:                name,
		JudgeConfigurations: map[string]interface{}{},
	}
}
