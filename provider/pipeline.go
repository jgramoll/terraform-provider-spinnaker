package provider

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

func pipelineResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineCreate,
		Read:   resourcePipelineRead,
		Update: resourcePipelineUpdate,
		Delete: resourcePipelineDelete,

		Schema: map[string]*schema.Schema{
			"application": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the application where the pipeline lives",
				Required:    true,
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Name of the pipeline",
				Required:    true,
			},
		},
	}
}

func resourcePipelineCreate(d *schema.ResourceData, m interface{}) error {
	var pipeline *client.Pipeline
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}

	log.Println(d.Get("name"))
	log.Println(pipeline)

	if pipeline.Name == "" {
		return fmt.Errorf("pipeline name must be provided")
	}
	if pipeline.Application == "" {
		return fmt.Errorf("pipeline application must be provided")
	}

	id := fmt.Sprintf("%s_%s", pipeline.Application, pipeline.Name)
	log.Printf("[DEBUG] Creating pipeline configuration: %#v", id)

	// c := m.(*client.Client)
	// ck, err := client.Checks.Create(check)
	// if err != nil {
	// 	return err
	// }

	d.SetId(id)
	return nil
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*client.Client)

	pipeline, err := c.GetPipeline("TODO.appName", d.Id())
	if err != nil {
		log.Printf("[WARN] No Server found: %s", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] got pipeline %s", pipeline.Id)
	d.Set("name", pipeline.Name)
	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	d.Partial(true)

	// if d.HasChange("address") {
	//   // Try updating the address
	//   if err := updateAddress(d, m); err != nil {
	//           return err
	//   }

	//   d.SetPartial("address")
	// }

	d.Partial(false)

	return nil
	// return resourceServerRead(d, m)
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
