package client

import (
	"testing"
)

func newTriggerTestPipeline() (*Pipeline, *Trigger) {
	expected := Trigger{}
	expected.ID = "triggerId"
	return &Pipeline{
		SerializablePipeline: SerializablePipeline{
			Triggers: []Trigger{expected},
		},
	}, &expected
}

func TestGetTrigger(t *testing.T) {
	pipeline, expected := newTriggerTestPipeline()
	s, err := pipeline.GetTrigger(expected.ID)
	if err != nil {
		t.Fatal(err)
	}
	if s.ID != expected.ID {
		t.Fatal("Not the expected trigger")
	}
}

func TestUpdateTrigger(t *testing.T) {
	pipeline, expected := newTriggerTestPipeline()

	updateTrigger := Trigger(*expected)
	updateTrigger.Master = "new name"
	err := pipeline.UpdateTrigger(&updateTrigger)
	if err != nil {
		t.Fatal(err)
	}
	if pipeline.Triggers[0].Master != updateTrigger.Master {
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
