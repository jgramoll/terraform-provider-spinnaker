package provider

import (
	"fmt"
	"log"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

var pipelineLock sync.Mutex

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
			"index": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Index of the pipeline",
				Optional:    true,
				Default:     0,
			},
		},
	}
}

func resourcePipelineCreate(d *schema.ResourceData, m interface{}) error {
	var pipeline client.CreatePipelineRequest
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline:", pipeline.Application, pipeline.Name)
	err := m.(*client.Client).CreatePipeline(&pipeline)
	if err != nil {
		return err
	}

	pipelineWithID, err := m.(*client.Client).GetPipeline(pipeline.Application, pipeline.Name)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		return err
	}

	log.Println("[DEBUG] New pipeline ID", pipelineWithID.ID)
	d.SetId(pipelineWithID.ID)
	// create can't update index...
	return resourcePipelineUpdate(d, m)
}

func resourcePipelineRead(d *schema.ResourceData, m interface{}) error {
	pipeline, err := m.(*client.Client).GetPipelineByID(d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline found:", d.Id())
		d.SetId("")
		return nil
	}

	log.Printf("[INFO] Got Pipeline %s", pipeline.ID)
	d.SetId(pipeline.ID)
	d.Set("application", pipeline.Application)
	d.Set("name", pipeline.Name)
	d.Set("index", pipeline.Index)
	return nil
}

func resourcePipelineUpdate(d *schema.ResourceData, m interface{}) error {
	var pipeline *client.Pipeline
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &pipeline); err != nil {
		return err
	}
	pipeline.ID = d.Id()

	err := m.(*client.Client).UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	return resourcePipelineRead(d, m)
}

func resourcePipelineDelete(d *schema.ResourceData, m interface{}) error {
	var pipeline *client.Pipeline
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
	return m.(*client.Client).DeletePipeline(pipeline)
}
