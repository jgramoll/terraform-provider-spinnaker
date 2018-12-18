package client

import (
	"testing"
)

var application *Application

func init() {
	application = NewApplication()
}

func TestNewApplication(t *testing.T) {
	expectedPort := 80
	if application.InstancePort != expectedPort {
		t.Fatalf("InstancePort should be %v, not \"%v\"", expectedPort, application.InstancePort)
	}
}

func TestNewProviderSettings(t *testing.T) {
	expectedUseAmiBlockDeviceMappings := false
	if application.ProviderSettings.AWS.UseAmiBlockDeviceMappings != expectedUseAmiBlockDeviceMappings {
		t.Fatalf("UseAmiBlockDeviceMappings should be %v, not \"%v\"", expectedUseAmiBlockDeviceMappings, application.ProviderSettings.AWS.UseAmiBlockDeviceMappings)
	}
}
