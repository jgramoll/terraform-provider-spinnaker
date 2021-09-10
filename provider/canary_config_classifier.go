package provider

import (
	"strconv"

	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type canaryConfigClassifiers []*canaryConfigClassifier

type canaryConfigClassifier struct {
	GroupWeights map[string]string `mapstructure:"group_weights"`
}

func (classifiers *canaryConfigClassifiers) toClientClassifier() *client.CanaryConfigClassifier {
	for _, classifier := range *classifiers {
		return &client.CanaryConfigClassifier{
			GroupWeights: *classifier.toClientGroupWeights(),
		}
	}
	return nil
}

func (*canaryConfigClassifiers) fromClientClassifier(classifier *client.CanaryConfigClassifier) *canaryConfigClassifiers {
	return &canaryConfigClassifiers{&canaryConfigClassifier{
		GroupWeights: *fromClientGroupWeights(&classifier.GroupWeights),
	}}
}

func (classifier *canaryConfigClassifier) toClientGroupWeights() *map[string]float64 {
	weights := map[string]float64{}
	for k, v := range classifier.GroupWeights {
		if floatValue, err := strconv.ParseFloat(v, 32); err == nil {
			weights[k] = floatValue
		}
	}
	return &weights
}

func fromClientGroupWeights(clientGroupWeights *map[string]float64) *map[string]string {
	weights := map[string]string{}
	for k, v := range *clientGroupWeights {
		weights[k] = strconv.FormatFloat(v, 'f', -1, 32)
	}
	return &weights
}
