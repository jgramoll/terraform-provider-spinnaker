package provider

import (
	"log"
	"os/user"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// Services used by provider
type Services struct {
	config             client.Config
	ApplicationService client.ApplicationService
	PipelineService    client.PipelineService
}

// Config for provider
type Config struct {
	Address   string
	CertPath  string `mapstructure:"cert_path"`
	KeyPath   string `mapstructure:"key_path"`
	UserEmail string `mapstructure:"user_email"`
}

// Provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},

			"cert_path": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},

			"key_path": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},

			"user_email": &schema.Schema{
				Type: schema.TypeString,
				// Required: true,
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"spinnaker_pipeline":              pipelineResource(),
			"spinnaker_pipeline_bake_stage":   pipelineBakeStageResource(),
			"spinnaker_pipeline_deploy_stage": pipelineDeployStageResource(),

			"spinnaker_pipeline_destroy_server_group_stage": pipelineDestroyServerGroupStageResource(),
			"spinnaker_pipeline_jenkins_stage":              pipelineJenkinsStageResource(),
			"spinnaker_pipeline_notification":               pipelineNotificationResource(),
			"spinnaker_pipeline_pipeline_stage":             pipelinePipelineResource(),
			"spinnaker_pipeline_rollback_cluster_stage":     pipelineRollbackClusterStageResource(),

			"spinnaker_pipeline_trigger": pipelineTriggerResource(),
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

	u, _ := user.Current()

	// TODO Why do we need this...
	// debug.PrintStack()
	// fmt.Println("config", config)
	// fmt.Println("config", client.Config(config))
	config.Address = "https://api.spinnaker.inseng.net"
	config.CertPath = u.HomeDir + "/.spin/client.crt"
	config.KeyPath = u.HomeDir + "/.spin/client.key"
	config.UserEmail = "rcolvin@instructure.com"

	clientConfig := client.Config(config)
	c := client.NewClient(clientConfig)
	return &Services{
		config:             clientConfig,
		ApplicationService: client.ApplicationService{Client: c},
		PipelineService:    client.PipelineService{Client: c},
	}, nil
}
