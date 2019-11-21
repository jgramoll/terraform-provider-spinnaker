package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Pipeline deploy pipeline in application
type pipeline struct {
	Application          string                 `mapstructure:"application"`
	AppConfig            map[string]interface{} `mapstructure:"appConfig"`
	Disabled             bool                   `mapstructure:"disabled"`
	ID                   string                 `mapstructure:"id"`
	KeepWaitingPipelines bool                   `mapstructure:"keep_waiting_pipelines"`
	LimitConcurrent      bool                   `mapstructure:"limit_concerrent"`
	Name                 string                 `mapstructure:"name"`
	Index                int                    `mapstructure:"index"`
	Roles                *[]string              `mapstructure:"roles"`
	ServiceAccount       string                 `mapstructure:"serviceAccount"`
}

func (p *pipeline) toClientPipeline() *client.Pipeline {
	return &client.Pipeline{
		Application:          p.Application,
		AppConfig:            p.AppConfig,
		Disabled:             p.Disabled,
		ID:                   p.ID,
		KeepWaitingPipelines: p.KeepWaitingPipelines,
		LimitConcurrent:      p.LimitConcurrent,
		Name:                 p.Name,
		Index:                p.Index,
		Roles:                p.Roles,
	}
}

func fromClientPipeline(p *client.Pipeline) *pipeline {
	return &pipeline{
		Application:          p.Application,
		AppConfig:            p.AppConfig,
		Disabled:             p.Disabled,
		ID:                   p.ID,
		KeepWaitingPipelines: p.KeepWaitingPipelines,
		LimitConcurrent:      p.LimitConcurrent,
		Name:                 p.Name,
		Index:                p.Index,
		Roles:                p.Roles,
	}
}

func (p *pipeline) setResourceData(d *schema.ResourceData) error {
	d.SetId(p.ID)
	err := d.Set(ApplicationKey, p.Application)
	if err != nil {
		return err
	}
	err = d.Set("name", p.Name)
	if err != nil {
		return err
	}
	err = d.Set("index", p.Index)
	if err != nil {
		return err
	}
	err = d.Set("disabled", p.Disabled)
	if err != nil {
		return err
	}
	err = d.Set("keep_waiting_pipelines", p.KeepWaitingPipelines)
	if err != nil {
		return err
	}
	err = d.Set("limit_concurrent", p.LimitConcurrent)
	if err != nil {
		return err
	}
	err = d.Set("roles", p.Roles)
	if err != nil {
		return err
	}
	err = d.Set("service_account", p.ServiceAccount)
	if err != nil {
		return err
	}
	return nil
}

// pipelineFromResourceData get pipeline from resource data
func pipelineFromResourceData(pipeline *client.Pipeline, d *schema.ResourceData) {
	pipeline.Index = d.Get("index").(int)
	pipeline.Application = d.Get(ApplicationKey).(string)
	pipeline.Name = d.Get("name").(string)
	pipeline.Disabled = d.Get("disabled").(bool)
	pipeline.KeepWaitingPipelines = d.Get("keep_waiting_pipelines").(bool)
	pipeline.LimitConcurrent = d.Get("limit_concurrent").(bool)
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
