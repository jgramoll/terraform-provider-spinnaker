package provider

import (
	"errors"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func canaryConfigDataSource() *schema.Resource {
	return &schema.Resource{
		Read: canaryConfigDataSourceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Canary Config Name",
				Required:    true,
			},
			"applications": {
				Type:        schema.TypeList,
				Description: "Applications",
				Optional:    true,
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
				Optional:    true,
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
				Optional:    true,
				Elem:        canaryConfigClassifierResource(),
			},
			"judge": {
				Type:        schema.TypeList,
				Description: "Canary Judge",
				MaxItems:    1,
				Optional:    true,
				Elem:        canaryConfigJudgeResource(),
			},
		},
	}
}

func canaryConfigDataSourceRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)

	log.Printf("[DEBUG] Importing canary config: %s\n", name)
	canaryConfigService := m.(*Services).CanaryConfigService
	canaryConfigs, err := canaryConfigService.GetCanaryConfigs()
	if err != nil {
		log.Printf("[WARN] No canary configs found: %s\n", err)
		return err
	}

	for _, c := range *canaryConfigs {
		if c.Name == name {
			log.Printf("[DEBUG] Imported canary config: %s\n", c.ID)
			d.SetId(c.ID)

			return resourceCanaryConfigRead(d, m)
		}
	}

	log.Printf("[WARN] No canary config found with name: %s\n", name)
	return errors.New("No canary config found")
}
