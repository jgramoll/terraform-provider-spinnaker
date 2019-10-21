package client

type CanaryConfig struct {
	Id            string                  `json:"id,omitempty"`
	Name          string                  `json:"name"`
	Applications  []string                `json:"applications"`
	Description   string                  `json:"description"`
	Metrics       []*CanaryConfigMetric   `json:"metrics"`
	ConfigVersion string                  `json:"configVersion"`
	Templates     map[string]interface{}  `json:"templates"`
	Classifier    *CanaryConfigClassifier `json:"classifier"`
	Judge         *CanaryConfigJudge      `json:"judge"`
}

func NewCanaryConfig(judge *CanaryConfigJudge, name string, application string) *CanaryConfig {
	return &CanaryConfig{
		Name:          name,
		Applications:  []string{application},
		Metrics:       []*CanaryConfigMetric{},
		ConfigVersion: "1",
		Templates:     map[string]interface{}{},
		Classifier:    NewCanaryConfigClassifier(),
		Judge:         judge,
	}
}

func (config *CanaryConfig) AddGroup(group string, weight float64) {
	config.Classifier.GroupWeights[group] = weight
}

func (config *CanaryConfig) AddMetric(metric *CanaryConfigMetric) {
	config.Metrics = append(config.Metrics, metric)
}
