package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineWebhookStageResource() *schema.Resource {
	newWebhookStageInterface := func() stage {
		return newWebhookStage()
	}
	return &schema.Resource{
		Create: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageCreate(d, m, newWebhookStageInterface)
		},
		Read: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageRead(d, m, newWebhookStageInterface)
		},
		Update: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageUpdate(d, m, newWebhookStageInterface)
		},
		Delete: func(d *schema.ResourceData, m interface{}) error {
			return resourcePipelineStageDelete(d, m, newWebhookStageInterface)
		},
		Importer: &schema.ResourceImporter{
			State: resourcePipelineImporter,
		},

		Schema: stageResource(map[string]*schema.Schema{
			"canceled_statuses": {
				Type:        schema.TypeString,
				Description: "Comma-separated list of strings that will be considered as CANCELED status.",
				Optional:    true,
			},
			"custom_headers": {
				Type:        schema.TypeMap,
				Description: "Key-value pairs to be sent as additional headers to the service.",
				Optional:    true,
			},
			"fail_fast_status_codes": {
				Type:        schema.TypeList,
				Description: "Comma-separated HTTP status codes (4xx or 5xx) that will cause this webhook stage to fail without retrying.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"method": {
				Type:        schema.TypeString,
				Description: "Config the HTTP method used for the webhook.",
				Optional:    true,
				Default:     "GET",
			},
			"payload_string": {
				Type:        schema.TypeString,
				Description: "JSON payload to be added to the webhook call.",
				Optional:    true,
			},
			"progress_json_path": {
				Type:        schema.TypeString,
				Description: "JSON path to a descriptive message about the progress in the webhook's response JSON. (e.g. $.buildInfo.progress)",
				Optional:    true,
			},
			"status_json_path": {
				Type:        schema.TypeString,
				Description: "JSON path to the status information in the webhook's response JSON. (e.g. $.buildInfo.status)",
				Optional:    true,
			},
			"status_url_json_path": {
				Type:        schema.TypeString,
				Description: "JSON path to the status url in the webhook's response JSON. (i.e. $.buildInfo.url)",
				Optional:    true,
			},
			"status_url_resolution": {
				Type:        schema.TypeString,
				Description: "Set the technique to lookup the overall status: webhookResponse - GET method against webhook URL; locationHeader - From the Location header; getMethod - From webhook’s response.",
				Optional:    true,
			},
			"success_statuses": {
				Type:        schema.TypeString,
				Description: "Comma-separated list of strings that will be considered as SUCCESS status.",
				Optional:    true,
			},
			"terminal_statuses": {
				Type:        schema.TypeString,
				Description: "Comma-separated list of strings that will be considered as TERMINAL status.",
				Optional:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "Config the url for the webhook.",
				Required:    true,
			},
		}),
	}
}
