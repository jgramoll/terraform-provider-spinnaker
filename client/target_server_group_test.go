package client

import (
	"fmt"
	"testing"
)

func TestControlServerGroupStageGetName(t *testing.T) {
	instances := map[StageType]*TargetServerGroupStage{
		DestroyServerGroupStageType: NewTargetServerGroupStage(DestroyServerGroupStageType),
		DisableServerGroupStageType: NewTargetServerGroupStage(DisableServerGroupStageType),
	}
	instances[DisableServerGroupStageType].Name = "New Disable Server Group"
	instances[DestroyServerGroupStageType].Name = "New Destroy Server Group"

	testCases := []*struct {
		StageName string
		Type      StageType
	}{{
		StageName: "New Destroy Server Group",
		Type:      DestroyServerGroupStageType,
	}, {
		StageName: "New Disable Server Group",
		Type:      DisableServerGroupStageType,
	}}

	for idx := range testCases {
		testCase := testCases[idx]
		t.Run(fmt.Sprintf("stage-%s", testCase.Type), func(t *testing.T) {
			if instances[testCase.Type].GetName() != testCase.StageName {
				t.Fatalf("%s instance GetName() method should be %q, not %q", testCase.Type, testCase.StageName, instances[testCase.Type].GetName())
			}
			if instances[testCase.Type].GetType() != testCase.Type {
				t.Fatalf("%s instance GetType() method should be %q, not %q", testCase.Type, testCase.Type, instances[testCase.Type].GetType())
			}
		})
	}
}
