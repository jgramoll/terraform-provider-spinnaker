package client

import (
	"testing"
)

func newTriggerTestPipeline() (*Pipeline, *JenkinsTrigger) {
	expected := NewJenkinsTrigger()
	expected.ID = "triggerId"
	return &Pipeline{
		Triggers: []Trigger{expected},
	}, expected
}

func TestGetTrigger(t *testing.T) {
	pipeline, expected := newTriggerTestPipeline()
	actual, err := pipeline.GetTrigger(expected.GetID())
	if err != nil {
		t.Fatal(err)
	}
	if actual.GetID() != expected.GetID() {
		t.Fatal("Not the expected trigger")
	}
}

func TestUpdateTrigger(t *testing.T) {
	pipeline, expected := newTriggerTestPipeline()

	updateTrigger := JenkinsTrigger(*expected)
	updateTrigger.Master = "new name"
	err := pipeline.UpdateTrigger(&updateTrigger)
	if err != nil {
		t.Fatal(err)
	}
	actualTrigger, ok := pipeline.Triggers[0].(*JenkinsTrigger)
	if !ok {
		t.Fatal("Should be Jenkins Trigger")
	}
	if actualTrigger.Master != updateTrigger.Master {
		t.Fatal("Pipeline Trigger was not updated")
	}
}

func TestDeleteTrigger(t *testing.T) {
	pipeline, expected := newTriggerTestPipeline()
	err := pipeline.DeleteTrigger(expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(pipeline.Triggers) != 0 {
		t.Fatal("Pipeline Trigger was not deleted")
	}
}
