package provider

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

func canaryConfigResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceCanaryConfigCreate,
		Read:   resourceCanaryConfigRead,
		Update: resourceCanaryConfigUpdate,
		Delete: resourceCanaryConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Canary Config Name",
				Required:    true,
			},
			"applications": {
				Type:        schema.TypeList,
				Description: "Applications",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:        schema.TypeString,
				Description: "Canary Config Description",
				Optional:    true,
			},
			"metric": {
				Type:        schema.TypeList,
				Description: "Canary Metrics",
				Required:    true,
				Elem:        canaryConfigMetricResource(),
			},
			"config_version": {
				Type:        schema.TypeString,
				Description: "Canary Config Version",
				Optional:    true,
				Default:     "1",
			},
			// Templates     map[string]interface{}  `mapstructure:"templates"`
			"classifier": {
				Type:        schema.TypeList,
				Description: "Canary Classifier",
				MaxItems:    1,
				Required:    true,
				Elem:        canaryConfigClassifierResource(),
			},
			"judge": {
				Type:        schema.TypeList,
				Description: "Canary Judge",
				MaxItems:    1,
				Required:    true,
				Elem:        canaryConfigJudgeResource(),
			},
		},
	}
}

func resourceCanaryConfigCreate(d *schema.ResourceData, m interface{}) error {
	var canaryConfig canaryConfig
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &canaryConfig); err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating canary config: %s\n", canaryConfig.Name)
	canaryConfigService := m.(*Services).CanaryConfigService
	id, err := canaryConfigService.CreateCanaryConfig(canaryConfig.toClientCanaryConfig(d.Id()))
	if err != nil {
		return err
	}

	d.SetId(id)
	return resourceCanaryConfigRead(d, m)
}

func resourceCanaryConfigRead(d *schema.ResourceData, m interface{}) error {
	canaryConfigService := m.(*Services).CanaryConfigService
	a, err := canaryConfigService.GetCanaryConfig(d.Id())
	if err != nil {
		if serr, ok := err.(*client.SpinnakerError); ok {
			if serr.Status == 404 {
				d.SetId("")
				return nil
			}
		}

		return err
	}

	log.Printf("[DEBUG] Got canary config: %s\n", a.Name)
	return fromClientCanaryConfig(a).setResourceData(d)
}

func resourceCanaryConfigUpdate(d *schema.ResourceData, m interface{}) error {
	var canaryConfig canaryConfig
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &canaryConfig); err != nil {
		return err
	}

	canaryConfigService := m.(*Services).CanaryConfigService
	err := canaryConfigService.UpdateCanaryConfig(canaryConfig.toClientCanaryConfig(d.Id()))
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updated canary config: %s\n", d.Id())
	return resourceCanaryConfigRead(d, m)
}

func resourceCanaryConfigDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] Deleting canary config: %s\n", d.Id())
	canaryConfigService := m.(*Services).CanaryConfigService
	err := canaryConfigService.DeleteCanaryConfig(d.Id())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
