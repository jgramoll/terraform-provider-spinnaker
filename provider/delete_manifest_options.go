package provider

import (
	"github.com/get-bridge/terraform-provider-spinnaker/client"
)

type deleteManifestOptions struct {
	Cascading bool `mapstructure:"cascading"`
}

func newDeleteManifestOptions() *deleteManifestOptions {
	return &deleteManifestOptions{
		Cascading: true,
	}
}

func toClientOptions(options *[]*deleteManifestOptions) *client.DeleteManifestOptions {
	if options != nil {
		for _, o := range *options {
			newOptions := client.NewDeleteManifestOptions()
			newOptions.Cascading = o.Cascading
			return newOptions
		}
	}
	return nil
}

func fromClientDeleteManifestOptions(o *client.DeleteManifestOptions) *[]*deleteManifestOptions {
	if o == nil {
		return nil
	}
	newOptions := newDeleteManifestOptions()
	newOptions.Cascading = o.Cascading
	array := []*deleteManifestOptions{newOptions}
	return &array
}
