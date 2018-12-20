package client

import (
	"testing"
)

func newStageTestPipeline() (*Pipeline, *BakeStage) {
	expected := NewBakeStage()
	expected.RefID = "stageId"
	return &Pipeline{
		Stages: []Stage{expected},
	}, expected
}

func TestGetStage(t *testing.T) {
	pipeline, expected := newStageTestPipeline()
	s, err := pipeline.GetStage(expected.RefID)
	if err != nil {
		t.Fatal(err)
	}
	if s != expected {
		t.Fatal("Not the expected stage")
	}
}

func TestUpdateStage(t *testing.T) {
	pipeline, expected := newStageTestPipeline()

	updateStage := BakeStage(*expected)
	updateStage.Name = "new name"
	err := pipeline.UpdateStage(&updateStage)
	if err != nil {
		t.Fatal(err)
	}
	if pipeline.Stages[0].GetName() == expected.Name {
		t.Fatal("Pipeline Stage was not updated")
	}
}

func TestDeleteStage(t *testing.T) {
	pipeline, expected := newStageTestPipeline()
	err := pipeline.DeleteStage(expected)
	if err != nil {
		t.Fatal(err)
	}
	if len(pipeline.Stages) != 0 {
		t.Fatal("Pipeline Stage was not deleted")
	}
}
