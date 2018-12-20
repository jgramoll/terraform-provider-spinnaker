package client

import (
	"testing"
)

var jenkinsStage JenkinsStage

func init() {
	jenkinsStage = *NewJenkinsStage()
}

func TestNewJenkinsStage(t *testing.T) {
	if jenkinsStage.Type != JenkinsStageType {
		t.Fatalf("Jenkins stage type should be %s, not \"%s\"", JenkinsStageType, jenkinsStage.Type)
	}
}

func TestJenkinsStageGetName(t *testing.T) {
	name := "New Jenkins"
	jenkinsStage.Name = name
	if jenkinsStage.GetName() != name {
		t.Fatalf("Jenkins stage GetName() should be %s, not \"%s\"", name, jenkinsStage.GetName())
	}
}

func TestJenkinsStageGetType(t *testing.T) {
	if jenkinsStage.GetType() != JenkinsStageType {
		t.Fatalf("Jenkins stage GetType() should be %s, not \"%s\"", JenkinsStageType, jenkinsStage.GetType())
	}
	if jenkinsStage.Type != JenkinsStageType {
		t.Fatalf("Jenkins stage Type should be %s, not \"%s\"", JenkinsStageType, jenkinsStage.Type)
	}
}
