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

// ErrNotificationNotFound notification not found
var ErrNotificationNotFound = errors.New("Could not find notification")

var errInvalidNotificationImportKey = errors.New("Invalid import key, must be pipelineID_notificationID")

func pipelineNotificationResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineNotificationCreate,
		Read:   resourcePipelineNotificationRead,
		Update: resourcePipelineNotificationUpdate,
		Delete: resourcePipelineNotificationDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				id := strings.Split(d.Id(), "_")
				if len(id) != 2 {
					return nil, errInvalidNotificationImportKey
				}
				d.Set(PipelineKey, id[0])
				d.SetId(id[1])
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			PipelineKey: {
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
				ForceNew:    true,
			},
			"address": {
				Type:        schema.TypeString,
				Description: "Address of the notification (slack channel, email, etc)",
				Required:    true,
			},
			"message": {
				Type:        schema.TypeList,
				Description: "Custom messages",
				Optional:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complete": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"failed": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"starting": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Type of notification (slack, email, etc)",
				Required:    true,
			},
			"when": {
				Type:        schema.TypeList,
				Description: "When to send notification (started, completed, failed)",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"complete": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"failed": {
							Type:     schema.TypeBool,
							Optional: true,
						},
						"starting": {
							Type:     schema.TypeBool,
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

	var notification defaultNotification
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

	cn, err := notification.toClientNotification(client.NotificationLevelPipeline)
	if err != nil {
		return err
	}
	pipeline.AppendNotification(cn)

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating pipeline notification: %s\n", id)
	d.SetId(id.String())
	return resourcePipelineNotificationRead(d, m)
}

func resourcePipelineNotificationRead(d *schema.ResourceData, m interface{}) error {
	pipelineID := d.Get(PipelineKey).(string)
	pipelineService := m.(*Services).PipelineService
	pipeline, err := pipelineService.GetPipelineByID(pipelineID)
	if err != nil {
		log.Printf("[WARN] No Pipeline found: %s\n", err)
		d.SetId("")
		return nil
	}

	var notification *client.Notification
	notification, err = pipeline.GetNotification(d.Id())
	if err != nil {
		log.Printf("[WARN] No Pipeline Notification found: %s\n", err)
		d.SetId("")
	} else {
		d.SetId(notification.ID)
		newDefaultNotification().fromClientNotification(notification).setNotificationResourceData(d)
	}

	return nil
}

func resourcePipelineNotificationUpdate(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var notification defaultNotification
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

	cn, err := notification.toClientNotification(client.NotificationLevelPipeline)
	if err != nil {
		return err
	}

	err = pipeline.UpdateNotification(cn)
	if err != nil {
		return err
	}

	err = pipelineService.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updated pipeline notifications: %s\n", d.Id())
	return resourcePipelineNotificationRead(d, m)
}

func resourcePipelineNotificationDelete(d *schema.ResourceData, m interface{}) error {
	pipelineLock.Lock()
	defer pipelineLock.Unlock()

	var notification defaultNotification
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
