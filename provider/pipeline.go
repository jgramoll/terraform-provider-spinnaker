package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Pipeline deploy pipeline in application
type Pipeline struct {
	Application          string                 `mapstructure:"application"`
	AppConfig            map[string]interface{} `mapstructure:"appConfig"`
	Disabled             bool                   `mapstructure:"disabled"`
	ID                   string                 `mapstructure:"id"`
	KeepWaitingPipelines bool                   `mapstructure:"keep_waiting_pipelines"`
	LimitConcurrent      bool                   `mapstructure:"limit_concerrent"`
	Name                 string                 `mapstructure:"name"`
	Index                int                    `mapstructure:"index"`
	ParameterConfig      *[]*pipelineParameter  `mapstructure:"parameter"`
	Roles                *[]string              `mapstructure:"roles"`
	ServiceAccount       string                 `mapstructure:"serviceAccount"`
}

// ToClientPipeline convert to client pipeline
// TODO better way?
func (pipeline *Pipeline) ToClientPipeline() *client.Pipeline {
	return &client.Pipeline{
		SerializablePipeline: client.SerializablePipeline{
			Application:          pipeline.Application,
			AppConfig:            pipeline.AppConfig,
			Disabled:             pipeline.Disabled,
			ID:                   pipeline.ID,
			KeepWaitingPipelines: pipeline.KeepWaitingPipelines,
			LimitConcurrent:      pipeline.LimitConcurrent,
			Name:                 pipeline.Name,
			Index:                pipeline.Index,
			ParameterConfig:      toClientPipelineConfig(pipeline.ParameterConfig),
			Roles:                pipeline.Roles,
		},
	}
}

func SetResourceData(pipeline *client.Pipeline, d *schema.ResourceData) error {
	d.SetId(pipeline.ID)
	err := d.Set(ApplicationKey, pipeline.Application)
	if err != nil {
		return err
	}
	err = d.Set("name", pipeline.Name)
	if err != nil {
		return err
	}
	err = d.Set("index", pipeline.Index)
	if err != nil {
		return err
	}
	err = d.Set("disabled", pipeline.Disabled)
	if err != nil {
		return err
	}
	err = d.Set("keep_waiting_pipelines", pipeline.KeepWaitingPipelines)
	if err != nil {
		return err
	}
	err = d.Set("limit_concurrent", pipeline.LimitConcurrent)
	if err != nil {
		return err
	}
	err = d.Set("parameter", pipeline.ParameterConfig)
	if err != nil {
		return err
	}
	err = d.Set("roles", pipeline.Roles)
	if err != nil {
		return err
	}
	err = d.Set("service_account", pipeline.ServiceAccount)
	if err != nil {
		return err
	}
	return nil
}

// PipelineFromResourceData get pipeline from resource data
func PipelineFromResourceData(pipeline *client.Pipeline, d *schema.ResourceData) {
	pipeline.Index = d.Get("index").(int)
	pipeline.Application = d.Get(ApplicationKey).(string)
	pipeline.Name = d.Get("name").(string)
	pipeline.Disabled = d.Get("disabled").(bool)
	pipeline.KeepWaitingPipelines = d.Get("keep_waiting_pipelines").(bool)
	pipeline.LimitConcurrent = d.Get("limit_concurrent").(bool)
	pipeline.ParameterConfig = pipelineParametersFromResourceData(d)
	pipeline.Roles = pipelineRolesFromResourceData(d)

	serviceAccount, ok := d.GetOk("service_account")
	if ok {
		pipeline.ServiceAccount = serviceAccount.(string)
	}
}

func pipelineRolesFromResourceData(d *schema.ResourceData) *[]string {
	rolesInterface, ok := d.GetOk("roles")
	if !ok {
		return nil
	}

	roles := []string{}
	for _, role := range rolesInterface.([]interface{}) {
		roles = append(roles, role.(string))
	}
	return &roles
}
