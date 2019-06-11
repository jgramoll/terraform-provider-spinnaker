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

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
				ForceNew:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the stage",
				Required:    true,
			},
			"requisite_stage_ref_ids": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage(s) that must be complete before this one",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"notification": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Notifications to send for stage results",
				Optional:    true,
				Elem:        notificationResource(),
			},
			"complete_other_branches_then_fail": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "halt this branch and fail the pipeline once other branches complete. Prevents any stages that depend on this stage from running, but allows other branches of the pipeline to run. The pipeline will be marked as failed once complete.",
				Optional:    true,
				Default:     false,
			},
			"continue_pipeline": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If false, marks the stage as successful right away without waiting for the jenkins job to complete",
				Optional:    true,
				Default:     false,
			},
			"fail_pipeline": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "If the stage fails, immediately halt execution of all running stages and fails the entire execution",
				Optional:    true,
				Default:     true,
			},
			"fail_on_failed_expressions": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "The stage will be marked as failed if it contains any failed expressions",
				Optional:    true,
				Default:     false,
			},
			"override_timeout": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Allows you to override the amount of time the stage can run before failing.\nNote: this represents the overall time the stage has to complete (the sum of all the task times).",
				Optional:    true,
				Default:     false,
			},
			"restrict_execution_during_time_window": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Restrict execution to specific time windows",
				Optional:    true,
				Default:     false,
			},
			"restricted_execution_window": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Time windows to restrict execution",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"jitter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
									"max_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"min_delay": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"skip_manual": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
									},
								},
							},
						},
						"whitelist": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end_hour": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"end_min": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_hour": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"start_min": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"stage_enabled": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Stage will only execute when the supplied expression evaluates true.\nThe expression does not need to be wrapped in ${ and }.\nIf this expression evaluates to false, the stages following this stage will still execute.",
				Optional:    true,
				MaxItems:    1,
				Elem:        stageEnabledResource(),
			},

			"canceled_statuses": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Comma-separated list of strings that will be considered as CANCELED status.",
				Optional:    true,
			},
			"custom_headers": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "Key-value pairs to be sent as additional headers to the service.",
				Optional:    true,
			},
			"fail_fast_status_codes": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Comma-separated HTTP status codes (4xx or 5xx) that will cause this webhook stage to fail without retrying.",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"method": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Config the HTTP method used for the webhook.",
				Optional:    true,
				Default:     "GET",
			},
			"payload": &schema.Schema{
				Type:        schema.TypeString,
				Description: "JSON payload to be added to the webhook call.",
				Optional:    true,
			},
			"progress_json_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "JSON path to a descriptive message about the progress in the webhook's response JSON. (e.g. $.buildInfo.progress)",
				Optional:    true,
			},
			"status_json_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "JSON path to the status information in the webhook's response JSON. (e.g. $.buildInfo.status)",
				Optional:    true,
			},
			"status_url_json_path": &schema.Schema{
				Type:        schema.TypeString,
				Description: "JSON path to the status url in the webhook's response JSON. (i.e. $.buildInfo.url)",
				Optional:    true,
			},
			"status_url_resolution": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Set the technique to lookup the overall status: webhookResponse - GET method against webhook URL; locationHeader - From the Location header; getMethod - From webhook’s response.",
				Optional:    true,
			},
			"success_statuses": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Comma-separated list of strings that will be considered as SUCCESS status.",
				Optional:    true,
			},
			"terminal_statuses": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Comma-separated list of strings that will be considered as TERMINAL status.",
				Optional:    true,
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Config the url for the webhook.",
				Required:    true,
			},
		},
	}
}
