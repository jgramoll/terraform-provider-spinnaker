package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type targetServerGroupStage struct {
	// baseStage
	Name                              string                   `mapstructure:"name"`
	RefID                             string                   `mapstructure:"ref_id"`
	Type                              client.StageType         `mapstructure:"type"`
	RequisiteStageRefIds              []string                 `mapstructure:"requisite_stage_ref_ids"`
	Notifications                     *[]*notification         `mapstructure:"notification"`
	StageEnabled                      *[]*stageEnabled         `mapstructure:"stage_enabled"`
	CompleteOtherBranchesThenFail     bool                     `mapstructure:"complete_other_branches_then_fail"`
	ContinuePipeline                  bool                     `mapstructure:"continue_pipeline"`
	FailOnFailedExpressions           bool                     `mapstructure:"fail_on_failed_expressions"`
	FailPipeline                      bool                     `mapstructure:"fail_pipeline"`
	OverrideTimeout                   bool                     `mapstructure:"override_timeout"`
	RestrictExecutionDuringTimeWindow bool                     `mapstructure:"restrict_execution_during_time_window"`
	RestrictedExecutionWindow         *[]*stageExecutionWindow `mapstructure:"restricted_execution_window"`
	// End baseStage

	CloudProvider     string      `mapstructure:"cloud_provider"`
	CloudProviderType string      `mapstructure:"cloud_provider_type"`
	Cluster           string      `mapstructure:"cluster"`
	Credentials       string      `mapstructure:"credentials"`
	Moniker           *[]*moniker `mapstructure:"moniker"`
	Regions           []string    `mapstructure:"regions"`
	Target            string      `mapstructure:"target"`
}

func newTargetServerGroupStage(stageType client.StageType) *targetServerGroupStage {
	return &targetServerGroupStage{Type: stageType}
}

func targetServerGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		PipelineKey: {
			Type:        schema.TypeString,
			Description: "Id of the pipeline to send notification",
			Required:    true,
			ForceNew:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "Name of the stage",
			Required:    true,
		},
		"complete_other_branches_then_fail": {
			Type:        schema.TypeBool,
			Description: "halt this branch and fail the pipeline once other branches complete. Prevents any stages that depend on this stage from running, but allows other branches of the pipeline to run. The pipeline will be marked as failed once complete.",
			Optional:    true,
			Default:     false,
		},
		"continue_pipeline": {
			Type:        schema.TypeBool,
			Description: "If false, marks the stage as successful right away without waiting for the jenkins job to complete",
			Optional:    true,
			Default:     false,
		},
		"fail_pipeline": {
			Type:        schema.TypeBool,
			Description: "If the stage fails, immediately halt execution of all running stages and fails the entire execution",
			Optional:    true,
			Default:     true,
		},
		"fail_on_failed_expressions": {
			Type:        schema.TypeBool,
			Description: "The stage will be marked as failed if it contains any failed expressions",
			Optional:    true,
			Default:     false,
		},
		"override_timeout": {
			Type:        schema.TypeBool,
			Description: "Allows you to override the amount of time the stage can run before failing.\nNote: this represents the overall time the stage has to complete (the sum of all the task times).",
			Optional:    true,
			Default:     false,
		},
		"restrict_execution_during_time_window": {
			Type:        schema.TypeBool,
			Description: "Restrict execution to specific time windows",
			Optional:    true,
			Default:     false,
		},
		"restricted_execution_window": {
			Type:        schema.TypeList,
			Description: "Time windows to restrict execution",
			Optional:    true,
			MaxItems:    1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"days": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Schema{
							Type: schema.TypeInt,
						},
					},
					"jitter": {
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
					"whitelist": {
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
		"requisite_stage_ref_ids": {
			Type:        schema.TypeList,
			Description: "Stage(s) that must be complete before this one",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"notification": {
			Type:        schema.TypeList,
			Description: "Notifications to send for stage results",
			Optional:    true,
			Elem:        notificationResource(),
		},
		"stage_enabled": {
			Type:        schema.TypeList,
			Description: "Stage will only execute when the supplied expression evaluates true.\nThe expression does not need to be wrapped in ${ and }.\nIf this expression evaluates to false, the stages following this stage will still execute.",
			Optional:    true,
			MaxItems:    1,
			Elem:        stageEnabledResource(),
		},
		"cloud_provider": {
			Type:        schema.TypeString,
			Description: "Cloud provider to use (aws)",
			Optional:    true,
		},
		"cloud_provider_type": {
			Type:        schema.TypeString,
			Description: "Cloud provider to use (aws)",
			Optional:    true,
		},
		"cluster": {
			Type:        schema.TypeString,
			Description: "Name of the cluster",
			Required:    true,
		},
		"credentials": {
			Type:        schema.TypeString,
			Description: "Name of the credentials to use",
			Optional:    true,
		},
		"moniker": {
			Type:        schema.TypeList,
			Description: "Name to attach to cluster",
			Optional:    true,
			MaxItems:    1,
			Elem:        monikerResource(),
		},
		"regions": {
			Type:        schema.TypeList,
			Description: "regions to target (us-east-1)",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"target": {
			Type:        schema.TypeString,
			Description: "Which server group to destroy (oldest_asg_dynamic, ancestor_asg_dynamic, current_asg_dynamic)",
			Optional:    true,
		},
	}
}

func (s *targetServerGroupStage) toClientStage(config *client.Config) (client.Stage, error) {
	// baseStage
	notifications, err := toClientNotifications(s.Notifications)
	if err != nil {
		return nil, err
	}

	cs := client.NewTargetServerGroupStage(s.Type)
	cs.Name = s.Name
	cs.RefID = s.RefID
	cs.RequisiteStageRefIds = s.RequisiteStageRefIds
	cs.Notifications = notifications
	cs.SendNotifications = notifications != nil && len(*notifications) > 0
	cs.StageEnabled = toClientStageEnabled(s.StageEnabled)
	cs.CompleteOtherBranchesThenFail = s.CompleteOtherBranchesThenFail
	cs.ContinuePipeline = s.ContinuePipeline
	cs.FailOnFailedExpressions = s.FailOnFailedExpressions
	cs.FailPipeline = s.FailPipeline
	cs.OverrideTimeout = s.OverrideTimeout
	cs.RestrictExecutionDuringTimeWindow = s.RestrictExecutionDuringTimeWindow
	cs.RestrictedExecutionWindow = toClientExecutionWindow(s.RestrictedExecutionWindow)
	// End baseStage

	cs.CloudProvider = s.CloudProvider
	cs.CloudProviderType = s.CloudProviderType
	cs.Cluster = s.Cluster
	cs.Credentials = s.Credentials
	cs.Moniker = toClientMoniker(s.Moniker)
	cs.Regions = s.Regions
	cs.Target = s.Target

	return cs, nil
}

func (s *targetServerGroupStage) fromClientStage(cs client.Stage) stage {
	clientStage := cs.(*client.TargetServerGroupStage)
	newStage := &targetServerGroupStage{}

	// baseStage
	newStage.Name = clientStage.Name
	newStage.RefID = clientStage.RefID
	newStage.RequisiteStageRefIds = clientStage.RequisiteStageRefIds
	newStage.Notifications = fromClientNotifications(clientStage.Notifications)
	newStage.StageEnabled = fromClientStageEnabled(clientStage.StageEnabled)
	newStage.CompleteOtherBranchesThenFail = clientStage.CompleteOtherBranchesThenFail
	newStage.ContinuePipeline = clientStage.ContinuePipeline
	newStage.FailOnFailedExpressions = clientStage.FailOnFailedExpressions
	newStage.FailPipeline = clientStage.FailPipeline
	newStage.OverrideTimeout = clientStage.OverrideTimeout
	newStage.RestrictExecutionDuringTimeWindow = clientStage.RestrictExecutionDuringTimeWindow
	newStage.RestrictedExecutionWindow = fromClientExecutionWindow(clientStage.RestrictedExecutionWindow)
	// end baseStage

	newStage.Type = clientStage.Type
	newStage.CloudProvider = clientStage.CloudProvider
	newStage.CloudProviderType = clientStage.CloudProviderType
	newStage.Cluster = clientStage.Cluster
	newStage.Credentials = clientStage.Credentials
	newStage.Moniker = fromClientMoniker(clientStage.Moniker)
	newStage.Regions = clientStage.Regions
	newStage.Target = clientStage.Target

	return newStage
}

func (s *targetServerGroupStage) SetResourceData(d *schema.ResourceData) error {
	// baseStage
	err := d.Set("name", s.Name)
	if err != nil {
		return err
	}
	err = d.Set("requisite_stage_ref_ids", s.RequisiteStageRefIds)
	if err != nil {
		return err
	}
	err = d.Set("notification", s.Notifications)
	if err != nil {
		return err
	}
	err = d.Set("stage_enabled", s.StageEnabled)
	if err != nil {
		return err
	}
	err = d.Set("complete_other_branches_then_fail", s.CompleteOtherBranchesThenFail)
	if err != nil {
		return err
	}
	err = d.Set("continue_pipeline", s.ContinuePipeline)
	if err != nil {
		return err
	}
	err = d.Set("fail_on_failed_expressions", s.FailOnFailedExpressions)
	if err != nil {
		return err
	}
	err = d.Set("fail_pipeline", s.FailPipeline)
	if err != nil {
		return err
	}
	err = d.Set("override_timeout", s.OverrideTimeout)
	if err != nil {
		return err
	}
	err = d.Set("restrict_execution_during_time_window", s.RestrictExecutionDuringTimeWindow)
	if err != nil {
		return err
	}
	err = d.Set("restricted_execution_window", s.RestrictedExecutionWindow)
	if err != nil {
		return err
	}
	// End baseStage

	err = d.Set("cloud_provider", s.CloudProvider)
	if err != nil {
		return err
	}
	err = d.Set("cloud_provider_type", s.CloudProviderType)
	if err != nil {
		return err
	}
	err = d.Set("cluster", s.Cluster)
	if err != nil {
		return err
	}
	err = d.Set("credentials", s.Credentials)
	if err != nil {
		return err
	}
	err = d.Set("moniker", s.Moniker)
	if err != nil {
		return err
	}
	err = d.Set("regions", s.Regions)
	if err != nil {
		return err
	}
	return d.Set("target", s.Target)
}

func (s *targetServerGroupStage) SetRefID(id string) {
	s.RefID = id
}

func (s *targetServerGroupStage) GetRefID() string {
	return s.RefID
}
