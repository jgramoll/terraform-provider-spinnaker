package provider

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

// complete = "%v is done"
// failed = "%v is failed"
// starting = "%v is starting"

type message struct {
	Complete string
	Failed   string
	Starting string
}

type notification struct {
	ID      string
	Address string
	Level   string
	Message message
	Type    string
	When    []string
}

func pipelineNotificationResource() *schema.Resource {
	return &schema.Resource{
		Create: resourcePipelineNotificationCreate,
		Read:   resourcePipelineNotificationRead,
		Update: resourcePipelineNotificationUpdate,
		Delete: resourcePipelineNotificationDelete,

		Schema: map[string]*schema.Schema{
			"pipeline": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Id of the pipeline to send notification",
				Required:    true,
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
				Type:        schema.TypeMap,
				Description: "Custom messages",
				Optional:    true,
				// 	Elem: &schema.Resource{
				// 		Schema: map[string]*schema.Schema{
				// 			"pipeline_complete": {
				// 				Type:     schema.TypeMap,
				// 				Required: true,
				// 			},
				// 			"pipeline_failed": {
				// 				Type:     schema.TypeMap,
				// 				Required: true,
				// 			},
				// 			"pipeline_starting": {
				// 				Type:     schema.TypeMap,
				// 				Required: true,
				// 			},
				// 		},
				// 	},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of notification (slack, email, etc)",
				Required:    true,
			},
			"when": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "When to send notification (started, completed, failed)",
				Required:    true,
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

	c := m.(*client.Client)
	pipeline, err := c.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	pipeline.Notifications = append(pipeline.Notifications, notification.toClientNotification())
	err = c.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Creating pipeline notification:", id)
	d.SetId(id.String())
	return resourcePipelineNotificationRead(d, m)
}

func (n notification) toClientNotification() client.Notification {
	return client.Notification{
		ID:      n.ID,
		Address: n.Address,
		Level:   n.Level,
		Message: n.Message.toClientMessage(),
		Type:    n.Type,
		When:    n.When,
	}
}

func (m message) toClientMessage() client.Message {
	return client.Message{
		Complete: client.MessageText{Text: m.Complete},
		Failed:   client.MessageText{Text: m.Failed},
		Starting: client.MessageText{Text: m.Starting},
	}
}

func resourcePipelineNotificationRead(d *schema.ResourceData, m interface{}) error {
	pipelineID := d.Get("pipeline").(string)
	pipeline, err := m.(*client.Client).GetPipelineByID(pipelineID)
	if err != nil {
		log.Println("[WARN] No Pipeline found:", err)
		d.SetId("")
		return nil
	}

	var notification *client.Notification
	notification, err = getNotification(pipeline.Notifications, d.Id())
	if err != nil {
		log.Println("[WARN] No Pipeline Notification found:", err)
		d.SetId("")
	} else {
		d.SetId(notification.ID)
		d.Set("address", notification.Address)
		d.Set("level", notification.Level)
		d.Set("message", message{
			Complete: notification.Message.Complete.Text,
			Starting: notification.Message.Starting.Text,
			Failed:   notification.Message.Failed.Text,
		})
		d.Set("type", notification.Type)
		d.Set("when", notification.When)
	}

	return nil
}

func getNotification(notifications []client.Notification, notificationID string) (*client.Notification, error) {
	for _, notification := range notifications {
		if notification.ID == notificationID {
			return &notification, nil
		}
	}
	return nil, fmt.Errorf("Could not find notification %v", notificationID)
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

	c := m.(*client.Client)
	pipeline, err := c.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	err = updateNotifications(pipeline, notification.toClientNotification())
	if err != nil {
		return err
	}

	err = c.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated pipeline notifications:", d.Id())
	return resourcePipelineNotificationRead(d, m)
}

func updateNotifications(pipeline *client.Pipeline, notification client.Notification) error {
	for i, t := range pipeline.Notifications {
		if t.ID == notification.ID {
			pipeline.Notifications[i] = notification
			return nil
		}
	}
	return fmt.Errorf("Could not find notification %v", notification.ID)
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

	c := m.(*client.Client)
	pipeline, err := c.GetPipelineByID(d.Get("pipeline").(string))
	if err != nil {
		return err
	}

	err = deleteNotification(pipeline, notification.toClientNotification())
	if err != nil {
		return err
	}

	err = c.UpdatePipeline(pipeline)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func deleteNotification(pipeline *client.Pipeline, notification client.Notification) error {
	for i, t := range pipeline.Notifications {
		if t.ID == notification.ID {
			pipeline.Notifications = append(pipeline.Notifications[:i], pipeline.Notifications[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find notification %v", notification.ID)
}
