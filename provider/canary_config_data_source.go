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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Canary Config Name",
				Required:    true,
			},
			"applications": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Applications",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Canary Config Description",
				Optional:    true,
			},
			"metric": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Canary Metrics",
				Optional:    true,
				Elem:        canaryConfigMetricResource(),
			},
			"config_version": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Canary Config Version",
				Optional:    true,
				Default:     "1",
			},
			// Templates     map[string]interface{}  `mapstructure:"templates"`
			"classifier": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Canary Classifier",
				MaxItems:    1,
				Optional:    true,
				Elem:        canaryConfigClassifierResource(),
			},
			"judge": &schema.Schema{
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

	log.Printf("[DEBUG] Importing canary config %s", name)
	canaryConfigService := m.(*Services).CanaryConfigService
	canaryConfigs, err := canaryConfigService.GetCanaryConfigs()
	if err != nil {
		log.Printf("[WARN] No canary configs found: %s", err)
		return err
	}

	for _, c := range *canaryConfigs {
		if c.Name == name {
			log.Printf("[DEBUG] Imported canary config %s", c.Id)
			d.SetId(c.Id)

			return resourceCanaryConfigRead(d, m)
		}
	}

	log.Printf("[WARN] No canary config found with name: %s", name)
	return errors.New("No canary config found")
}
