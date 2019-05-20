package provider

import (
	"log"

	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

const (
	// ApplicationKey key for application in map
	ApplicationKey = "application"
)

func applicationResource() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationCreate,
		Read:   resourceApplicationRead,
		Update: resourceApplicationUpdate,
		Delete: resourceApplicationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Application Name",
				Required:    true,
				ForceNew:    true,
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Email Address",
				Required:    true,
			},
			"repo_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Email Address",
				Optional:    true,
			},
			"repo_project_key": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Email Address",
				Optional:    true,
			},
			"repo_slug": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Email Address",
				Optional:    true,
			},
			"cloud_providers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Cloud Providers",
				Required:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"provider_settings": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aws": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"use_ami_block_device_mappings": &schema.Schema{
										Description: "Prefer AMI Block Device Mappings",
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     false,
									},
								},
							},
						},
					},
				},
			},
			"instance_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"platform_health_only": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Consider only cloud provider health when executing tasks",
				Optional:    true,
				Default:     true,
			},
			"platform_health_only_show_override": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Show health override option for each operation",
				Optional:    true,
				Default:     false,
			},
			"enable_restart_running_executions": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Enable restarting running pipelines",
				Optional:    true,
				Default:     true,
			},
		},
	}
}

func resourceApplicationCreate(d *schema.ResourceData, m interface{}) error {
	var application application
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &application); err != nil {
		return err
	}

	log.Println("[DEBUG] Creating application:", application.Name)
	applicationService := m.(*Services).ApplicationService
	err := applicationService.CreateApplication(application.toClientApplication())
	if err != nil {
		return err
	}

	d.SetId(application.Name)

	// TODO
	// The process to create application is asynchronous
	// Need to update CreateApplication method to waiting/checking task before return
	time.Sleep(5 * time.Second)

	return resourceApplicationRead(d, m)
}

func resourceApplicationRead(d *schema.ResourceData, m interface{}) error {
	applicationService := m.(*Services).ApplicationService
	a, err := applicationService.GetApplicationByName(d.Id())
	if err != nil {
		if serr, ok := err.(*client.SpinnakerError); ok {
			if serr.Status == 404 {
				d.SetId("")
				return nil
			}
		}

		return err
	}

	log.Printf("[DEBUG] Got application %s", a.Name)
	return fromClientApplication(a).setResourceData(d)
}

func resourceApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	applicationService := m.(*Services).ApplicationService
	application, err := applicationService.GetApplicationByName(d.Id())
	if err != nil {
		return err
	}

	applicationFromResourceData(application, d)

	err = applicationService.UpdateApplication(application)
	if err != nil {
		return err
	}

	log.Println("[DEBUG] Updated application:", d.Id())
	return resourceApplicationRead(d, m)
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
	var a application
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &a); err != nil {
		return err
	}

	log.Println("[DEBUG] Deleting application:", d.Id())
	applicationService := m.(*Services).ApplicationService
	return applicationService.DeleteApplication(a.toClientApplication())
}
