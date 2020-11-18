package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineCanaryAnalysisConfigResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"canary_analysis_interval_mins": {
				Type: schema.TypeString,
				Description: "The frequency at which a canary score is generated. The recommended interval is at least 30 minutes." +
					"" +
					"If an interval is not specified, or the specified interval is larger than the overall time range, there will be one canary run over the full time range.",
				Optional: true,
			},
			"canary_config_id": {
				Type:        schema.TypeString,
				Description: "Id of the canary config.",
				Required:    true,
			},
			"lifetime_duration": {
				Type: schema.TypeString,
				Description: "The total time for which data will be collected and analyzed during this stage." +
					"Example: 'PT1H5M' means 1 hour 5 minutes",
				Required: true,
			},
			"metrics_account_name": {
				Type:        schema.TypeString,
				Description: "The name of the metrics account configured in spinnaker.",
				Required:    true,
			},
			"scope": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"control_location": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"control_scope": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"experiment_location": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"experiment_scope": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"extended_scope_params": {
							Type: schema.TypeMap,
							Description: "Metric source specific parameters which may be used to further alter the canary scope." +
								"" +
								"Also used to provide variable bindings for use in the expansion of custom filter templates within the canary config.",
							Optional: true,
						},
						"scope_name": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "default",
						},
						"step": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"score_thresholds": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"marginal": {
							Type: schema.TypeString,
							Description: "A canary stage can include multiple canary runs." +
								"" +
								"If a given canary run score is less than or equal to the marginal threshold, the canary stage will fail immediately." +
								"" +
								"If the canary run score is greater than the marginal threshold, the canary stage will not fail and will execute the remaining downstream canary runs.",
							Required: true,
						},
						"pass": {
							Type:        schema.TypeString,
							Description: "When all canary runs in a stage have executed, a canary stage is considered a success if the final (that is, the latest) canary run score is greater than or equal to the pass threshold. Otherwise, it is a failure.",
							Required:    true,
						},
					},
				},
			},
			"storage_account_name": {
				Type:        schema.TypeString,
				Description: "The name of the storage account configured in spinnaker.",
				Required:    true,
			},
		},
	}
}
