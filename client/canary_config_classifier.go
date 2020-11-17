package client

// CanaryConfigClassifier classifier
type CanaryConfigClassifier struct {
	GroupWeights map[string]float64 `json:"groupWeights"`
}

// NewCanaryConfigClassifier new classifier
func NewCanaryConfigClassifier() *CanaryConfigClassifier {
	return &CanaryConfigClassifier{
		GroupWeights: map[string]float64{},
	}
}
