package client

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPipeline(t *testing.T) {
	pipeline := NewPipeline()
	if pipeline.Disabled {
		t.Fatal("default pipeline should not be disabled")
	}
}

func TestParseEmptyPipeline(t *testing.T) {
	pipeline, err := parsePipeline(nil)
	if err != nil {
		t.Fatal(err)
	}
	if pipeline.Disabled {
		t.Fatal("default pipeline should not be disabled")
	}
}

func TestParsePipeline(t *testing.T) {
	expectedName := "test-expected"
	expectedStageName := "test-stage"
	expectedStageAmiName := "ami-test"
	pipeline, err := parsePipeline(map[string]interface{}{
		"name": expectedName,
		"stages": []interface{}{
			map[string]interface{}{
				"type":    BakeStageType.String(),
				"name":    expectedStageName,
				"amiName": expectedStageAmiName,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	expected := NewPipeline()
	expected.Name = expectedName
	expected.Stages = []Stage{
		NewBakeStage(),
	}
	stage := expected.Stages[0].(*BakeStage)
	stage.Name = expectedStageName
	stage.AmiName = expectedStageAmiName
	err = pipeline.equals(expected)
	if err != nil {
		t.Fatal(err)
	}
}

func (pipeline *Pipeline) equals(expected *Pipeline) error {
	if !reflect.DeepEqual(pipeline, expected) {
		return fmt.Errorf("Pipeline %v does not match %v", pipeline, expected)
	}
	return nil
}
