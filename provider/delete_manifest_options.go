package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deleteManifestOptions struct {
	Cascading bool `mapstructure:"cascading"`
}

func newDeleteManifestOptions() *deleteManifestOptions {
	return &deleteManifestOptions{
		Cascading: true,
	}
}

func (o *deleteManifestOptions) toClientOptions() *client.DeleteManifestOptions {
	newOptions := client.NewDeleteManifestOptions()
	newOptions.Cascading = o.Cascading
	return newOptions
}

func fromClientDeleteManifestOptions(o *client.DeleteManifestOptions) *deleteManifestOptions {
	newOptions := newDeleteManifestOptions()
	newOptions.Cascading = o.Cascading
	return newOptions
}
