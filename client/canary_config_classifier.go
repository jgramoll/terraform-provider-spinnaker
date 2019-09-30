package client

type CanaryConfigClassifier struct {
	GroupWeights map[string]float64 `json:"groupWeights"`
}

func NewCanaryConfigClassifier() *CanaryConfigClassifier {
	return &CanaryConfigClassifier{
		GroupWeights: map[string]float64{},
	}
}
