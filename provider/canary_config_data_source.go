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

	log.Println("[DEBUG] Importing canary config", name)
	canaryConfigService := m.(*Services).CanaryConfigService
	canaryConfigs, err := canaryConfigService.GetCanaryConfigs()
	if err != nil {
		log.Println("[WARN] No canary configs found:", err)
		return err
	}

	for _, c := range *canaryConfigs {
		if c.Name == name {
			log.Println("[DEBUG] Imported canary config", c.Id)
			d.SetId(c.Id)

			return resourceCanaryConfigRead(d, m)
		}
	}

	log.Println("[WARN] No canary config found with name:", name)
	return errors.New("No canary config found")
}
