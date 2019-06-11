package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// Services used by provider
type Services struct {
	Config             client.Config
	ApplicationService client.ApplicationService
	PipelineService    client.PipelineService
}

// Config for provider
type Config struct {
	Address string
	Auth    Auth
}

// Auth for provider
type Auth struct {
	Enabled   bool
	CertPath  string `mapstructure:"cert_path"`
	KeyPath   string `mapstructure:"key_path"`
	UserEmail string `mapstructure:"user_email"`
}

// Provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_ADDRESS", nil),
				Description: "Address of spinnaker api",
			},

			"auth": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    false,
							Description: "Path to cert to authenticate with spinnaker api",
						},

						"cert_path": &schema.Schema{
							Type:        schema.TypeString,
							Required:    false,
							DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_CERT", nil),
							Description: "Path to cert to authenticate with spinnaker api",
						},

						"key_path": &schema.Schema{
							Type:        schema.TypeString,
							Required:    false,
							DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_KEY", nil),
							Description: "Path to key to authenticate with spinnaker api",
						},

						"user_email": &schema.Schema{
							Type:        schema.TypeString,
							Required:    false,
							DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_EMAIL", nil),
							Description: "Path to user_email to authenticate with spinnaker api",
						},
					},
				},
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"spinnaker_application":                         applicationResource(),
			"spinnaker_pipeline":                            pipelineResource(),
			"spinnaker_pipeline_bake_stage":                 pipelineBakeStageResource(),
			"spinnaker_pipeline_find_image_from_tags_stage": pipelineFindImageFromTagsStageResource(),
			"spinnaker_pipeline_deploy_stage":               pipelineDeployStageResource(),
			"spinnaker_pipeline_manual_judgment_stage":      pipelineManualJudgementStageResource(),
			"spinnaker_pipeline_destroy_server_group_stage": pipelineDestroyServerGroupStageResource(),
			"spinnaker_pipeline_jenkins_stage":              pipelineJenkinsStageResource(),
			"spinnaker_pipeline_pipeline_stage":             pipelinePipelineResource(),
			"spinnaker_pipeline_resize_server_group_stage":  pipelineResizeServerGroupStageResource(),
			"spinnaker_pipeline_rollback_cluster_stage":     pipelineRollbackClusterStageResource(),

			"spinnaker_pipeline_notification": pipelineNotificationResource(),

			"spinnaker_pipeline_parameter":     pipelineParameterResource(),
			"spinnaker_pipeline_trigger":       pipelineTriggerResource(),
			"spinnaker_pipeline_webhook_stage": pipelineWebhookStageResource(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var config Config
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing Spinnaker client")

	clientConfig := client.Config{
		Address: config.Address,
		Auth: client.Auth{
			Enabled:   config.Auth.Enabled,
			CertPath:  config.Auth.CertPath,
			KeyPath:   config.Auth.KeyPath,
			UserEmail: config.Auth.UserEmail,
		},
	}
	c := client.NewClient(clientConfig)
	return &Services{
		Config:             clientConfig,
		ApplicationService: client.ApplicationService{Client: c},
		PipelineService:    client.PipelineService{Client: c},
	}, nil
}
