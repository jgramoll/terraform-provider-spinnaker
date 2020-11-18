package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineRunJobManifestStageResource() *schema.Resource {
	newRunJobManifestStageInterface := func() stage {
		return newRunJobManifestStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newRunJobManifestStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newRunJobManifestStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newRunJobManifestStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newRunJobManifestStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "A Spinnaker account corresponds to a physical Kubernetes cluster. If you are unsure which account to use, talk to your Spinnaker admin.",
				Required:    true,
			},
			"application": {
				Type:        schema.TypeString,
				Description: "This is the Spinnaker application that your manifest will be deployed to. An application is generally used to group resources that belong to a single service.",
				Required:    true,
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "The cloudprovider to handle the manifest",
				Required:    true,
			},
			"consume_artifact_source": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"credentials": {
				Type:        schema.TypeString,
				Description: "Spinnaker credientials to use to talk to cloud provider",
				Optional:    true,
			},
			"manifest": {
				Type:        schema.TypeString,
				Description: "Manifest Yaml as text",
				Required:    true,
			},
			"property_file": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source": {
				Type: schema.TypeString,
				Description: "Where the manifest file content is read from." +
					"text: The manifest is supplied statically to the pipeline from the below text-box." +
					"artifact: The manifest is read from an artifact supplied/created upstream. The expected artifact must be referenced here, and will be bound at runtime.",
				Required: true,
			},
		}),
	}
}
