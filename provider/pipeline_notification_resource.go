package provider

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// ErrNotificationNotFound notification not found
var ErrNotificationNotFound = errors.New("Could not find notification")

const PipelineKey = "pipeline"

func pipelineNotificationResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineNotificationCreate,
		Read:   resourcePipelineNotificationRead,
		Update: resourcePipelineNotificationUpdate,
		Delete: resourcePipelineNotificationDelete,

		Schema: map[string]*schema.Schema{
			PipelineKey: &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
				ForceNew:    true,
			},
			"address": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Address of the notification (slack channel, email, etc)",
				Required:    true,
			},
			"level": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Level of the notification (pipeline, stage)",
				Required:    true,
			},
			"message": {
				Type:        schema.TypeList,
				Description: "Custom messages",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complete": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"failed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"starting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of notification (slack, email, etc)",
				Required:    true,
			},
			"when": &schema.Schema{
				// TODO validate more
				Type:        schema.TypeList,
				Description: "When to send notification (started, completed, failed)",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complete": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"failed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"starting": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourcePipelineNotificationCreate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var notification notification
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &notification); err != nil {
		return err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	notification.ID = id.String()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	pipeline.Notifications = append(pipeline.Notifications, *notification.toClientNotification())
	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline notification:", id)
	d.SetId(id.String())
	return resourcePipelineNotificationRead(d, m)
}

func resourcePipelineNotificationRead(d *schema.ResourceData, m interface{}) error {
	pipelineID := d.Get(PipelineKey).(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var notification *client.Notification
	notification, err = pipeline.GetNotification(d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Notification found:", err)
		d.SetId("")
	} else {
		d.SetId(notification.ID)
		d.Set("address", notification.Address)
		d.Set("level", notification.Level)
		newMessage := message{}
		if notification.Message.Complete != nil {
			newMessage.Complete = notification.Message.Complete.Text
		}
		if notification.Message.Starting != nil {
			newMessage.Starting = notification.Message.Starting.Text
		}
		if notification.Message.Failed != nil {
			newMessage.Failed = notification.Message.Failed.Text
		}
		d.Set("message", newMessage)
		d.Set("type", notification.Type)
		d.Set("when", notification.When)
	}

	return nil
}

func resourcePipelineNotificationUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var notification notification
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &notification); err != nil {
		return err
	}
	notification.ID = d.Id()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = pipeline.UpdateNotification(notification.toClientNotification())
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline notifications:", d.Id())
	return resourcePipelineNotificationRead(d, m)
}

func resourcePipelineNotificationDelete(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var notification notification
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &notification); err != nil {
		return err
	}
	notification.ID = d.Id()

	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(d.Get(PipelineKey).(string))
	if err != nil {
		return err
	}

	err = pipeline.DeleteNotification(notification.ID)
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
