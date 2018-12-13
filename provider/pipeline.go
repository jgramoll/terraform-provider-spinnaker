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
	var pipeline client.Pipeline
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}

	if pipeline.Name == "" {
		return fmt.Errorf("pipeline name must be provided")
	}
	if pipeline.Application == "" {
		return fmt.Errorf("pipeline application must be provided")
	}

	c := m.(*client.Client)
	err := c.CreatePipeline(&pipeline)
	if err != nil {
		return err
	}

	id := fmt.Sprintf("%s_%s", pipeline.Application, pipeline.Name)
	log.Println("[DEBUG] Creating pipeline:", id)
	d.SetId(id)
	return nil
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {
	application := d.Get("application").(string)
	name := d.Get("name").(string)
	if name == "" {
		log.Println("[WARN] No Pipeline name", d.Id())
	}

	c := m.(*client.Client)
	pipeline, err := c.GetPipeline(application, name)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Got Pipeline %s_%s\n", pipeline.Application, pipeline.Name)
	d.Set("name", pipeline.Name)
	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	// log.Println("[DEBUG] Updating pipeline:", id)
	// Enable partial state mode
	// d.Partial(true)

	// if d.HasChange("address") {
	//   // Try updating the address
	//   if err := updateAddress(d, m); err != nil {
	//           return err
	//   }

	//   d.SetPartial("address")
	// }

	// d.Partial(false)

	return nil
	// return resourceServerRead(d, m)
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	var pipeline client.Pipeline
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}

	if pipeline.Name == "" {
		return fmt.Errorf("pipeline name must be provided")
	}
	if pipeline.Application == "" {
		return fmt.Errorf("pipeline application must be provided")
	}

	log.Println("[DEBUG] Deleting pipeline:", d.Id())
	d.SetId("")
	c := m.(*client.Client)
	return c.DeletePipeline(&pipeline)
}
