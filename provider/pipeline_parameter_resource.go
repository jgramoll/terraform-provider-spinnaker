package provider

import (
	"errors"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

var errInvalidParameterImportKey = errors.New("Invalid import key, must be pipelineID_parameterID")

func pipelineParameterResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineParameterCreate,
		Read:   resourcePipelineParameterRead,
		Update: resourcePipelineParameterUpdate,
		Delete: resourcePipelineParameterDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				id := strings.Split(d.Id(), "_")
				if len(id) != 2 {
					return nil, errInvalidParameterImportKey
				}
				d.Set(PipelineKey, id[0])
				d.SetId(id[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to add parameter",
				Required:    true,
				ForceNew:    true,
			},
			"default": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Default value",
				Optional:    true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": &schema.Schema{
				Type:        schema.TypeString,
				Description: "A label to display when users are triggering the pipeline manually",
				Optional:    true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"option": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"required": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourcePipelineParameterCreate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var parameter pipelineParameter
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &parameter); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	parameter.ID = id.String()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	clientParameter := toClientPipelineParameter(&parameter)
	pipeline.AppendParameter(clientParameter)

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline parameter:", id)
	d.SetId(id.String())
	return resourcePipelineParameterRead(d, m)
}

func resourcePipelineParameterRead(d *schema.ResourceData, m interface{}) error {
	pipelineID := d.Get(PipelineKey).(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var parameter *client.PipelineParameter
	parameter, err = pipeline.GetParameter(d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Parameter found:", err)
		d.SetId("")
	} else {
		d.SetId(parameter.ID)
		fromClientPipelineParameter(parameter).setResourceData(d)
	}

	return nil
}

func resourcePipelineParameterUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var parameter pipelineParameter
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &parameter); err != nil {
		return err
	}
	parameter.ID = d.Id()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	clientParameter := toClientPipelineParameter(&parameter)
	err = pipeline.UpdateParameter(clientParameter)
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline parameter:", d.Id())
	return resourcePipelineParameterRead(d, m)
}

func resourcePipelineParameterDelete(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = pipeline.DeleteParameter(d.Id())
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}
