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
	Config              *client.Config
	ApplicationService  *client.ApplicationService
	CanaryConfigService *client.CanaryConfigService
	PipelineService     *client.PipelineService
}

// Provider for terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		TerraformVersion: ">= 0.12",
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_ADDRESS", nil),
				Description: "Address of spinnaker api",
			},

			"cert_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_CERT", nil),
				Description: "Path to cert to authenticate with spinnaker api",
			},

			"key_path": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_KEY", nil),
				Description: "Path to key to authenticate with spinnaker api",
			},

			"cert_content": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_CERT_CONTENT", nil),
				Description: "Cert string in base64 to authenticate with spinnaker api",
			},

			"key_content": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_KEY_CONTENT", nil),
				Description: "Key string in base64 to authenticate with spinnaker api",
			},

			"user_email": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SPINNAKER_EMAIL", nil),
				Description: "Path to user_email to authenticate with spinnaker api",
			},

			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If http client should skip ssl validation",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"spinnaker_canary_config": canaryConfigDataSource(),
			"spinnaker_pipeline":      pipelineDataSource(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"spinnaker_application":   applicationResource(),
			"spinnaker_canary_config": canaryConfigResource(),

			"spinnaker_pipeline":            pipelineResource(),
			"spinnaker_pipeline_bake_stage": pipelineBakeStageResource(),

			"spinnaker_pipeline_bake_manifest_stage":         pipelineBakeManifestStageResource(),
			"spinnaker_pipeline_canary_analysis_stage":       pipelineCanaryAnalysisStageResource(),
			"spinnaker_pipeline_check_preconditions_stage":   pipelineCheckPreconditionsStageResource(),
			"spinnaker_pipeline_delete_manifest_stage":       pipelineDeleteManifestStageResource(),
			"spinnaker_pipeline_deploy_cloudformation_stage": pipelineDeployCloudformationStageResource(),
			"spinnaker_pipeline_deploy_manifest_stage":       pipelineDeployManifestStageResource(),
			"spinnaker_pipeline_deploy_stage":                pipelineDeployStageResource(),
			"spinnaker_pipeline_destroy_server_group_stage":  pipelineDestroyServerGroupStageResource(),
			"spinnaker_pipeline_disable_manifest_stage":      pipelineDisableManifestStageResource(),
			"spinnaker_pipeline_disable_server_group_stage":  pipelineDisableServerGroupStageResource(),
			"spinnaker_pipeline_enable_server_group_stage":   pipelineEnableServerGroupStageResource(),
			"spinnaker_pipeline_enable_manifest_stage":       pipelineEnableManifestStageResource(),
			"spinnaker_pipeline_evaluate_variables_stage":    pipelineEvaluateVariablesStageResource(),

			"spinnaker_pipeline_find_artifacts_from_resource_stage": pipelineFindArtifactsFromResourceStageResource(),

			"spinnaker_pipeline_find_image_from_tags_stage":  pipelineFindImageFromTagsStageResource(),
			"spinnaker_pipeline_jenkins_stage":               pipelineJenkinsStageResource(),
			"spinnaker_pipeline_manual_judgment_stage":       pipelineManualJudgementStageResource(),
			"spinnaker_pipeline_notification":                pipelineNotificationResource(),
			"spinnaker_pipeline_patch_manifest_stage":        pipelinePatchManifestStageResource(),
			"spinnaker_pipeline_pipeline_stage":              pipelinePipelineResource(),
			"spinnaker_pipeline_resize_server_group_stage":   pipelineResizeServerGroupStageResource(),
			"spinnaker_pipeline_rollback_cluster_stage":      pipelineRollbackClusterStageResource(),
			"spinnaker_pipeline_run_job_manifest_stage":      pipelineRunJobManifestStageResource(),
			"spinnaker_pipeline_scale_manifest_stage":        pipelineScaleManifestStageResource(),
			"spinnaker_pipeline_undo_rollout_manifest_stage": pipelineUndoRolloutManifestStageResource(),
			"spinnaker_pipeline_webhook_stage":               pipelineWebhookStageResource(),
			"spinnaker_pipeline_script_stage":                pipelineScriptStageResource(),

			"spinnaker_pipeline_parameter": pipelineParameterResource(),

			"spinnaker_pipeline_trigger":          pipelineJenkinsTriggerResource("use spinnaker_pipeline_jenkins_trigger"),
			"spinnaker_pipeline_docker_trigger":   pipelineDockerTriggerResource(),
			"spinnaker_pipeline_jenkins_trigger":  pipelineJenkinsTriggerResource(""),
			"spinnaker_pipeline_pipeline_trigger": pipelinePipelineTriggerResource(),
			"spinnaker_pipeline_webhook_trigger":  pipelineWebhookTriggerResource(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := newProviderConfig()
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.WeakDecode(configRaw, &config); err != nil {
		return nil, err
	}

	log.Println("[INFO] Initializing Spinnaker client")

	clientConfig := config.toClientConfig()
	c, err := client.NewClient(clientConfig)
	if err != nil {
		return nil, err
	}
	return &Services{
		Config:              clientConfig,
		ApplicationService:  &client.ApplicationService{Client: c},
		CanaryConfigService: &client.CanaryConfigService{Client: c},
		PipelineService:     &client.PipelineService{Client: c},
	}, nil
}
