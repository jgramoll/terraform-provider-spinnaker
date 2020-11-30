package provider

import (
	"errors"
	"log"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
	"github.com/mitchellh/mapstructure"
)

const (
	// ApplicationKey key for application in map
	ApplicationKey = "application"
)

var (
	applicationNameRegex = regexp.MustCompile("^[a-zA-Z_0-9.-]*$")
	emailRegex           = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
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
			"accounts": {
				Type:        schema.TypeList,
				Description: "Accounts",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"cloud_providers": {
				Type:        schema.TypeList,
				Description: "Cloud Providers",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"email": {
				Type:        schema.TypeString,
				Description: "Email Address",
				Required:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if !emailRegex.MatchString(v) {
						errs = append(errs, errors.New("a valid email address is required"))
					}
					return
				},
			},
			"enable_restart_running_executions": {
				Type:        schema.TypeBool,
				Description: "Enable restarting running pipelines",
				Optional:    true,
				Default:     true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "Application Name",
				Required:    true,
				ForceNew:    true,
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if len(v) > 249 {
						errs = append(errs, errors.New("applicaiton name must be shorter than 250 characters"))
					}

					if !applicationNameRegex.MatchString(v) {
						errs = append(errs, errors.New("application name can't have special characters or spaces"))
					}
					return
				},
			},
			"instance_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"permissions": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"execute": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"read": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"write": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"platform_health_only": {
				Type:        schema.TypeBool,
				Description: "Consider only cloud provider health when executing tasks",
				Optional:    true,
				Default:     true,
			},
			"platform_health_only_show_override": {
				Type:        schema.TypeBool,
				Description: "Show health override option for each operation",
				Optional:    true,
				Default:     false,
			},
			"provider_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aws": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"use_ami_block_device_mappings": {
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
			"repo_project_key": {
				Type:        schema.TypeString,
				Description: "Repository project key",
				Optional:    true,
			},
			"repo_slug": {
				Type:        schema.TypeString,
				Description: "Repository slug",
				Optional:    true,
			},
			"repo_type": {
				Type:        schema.TypeString,
				Description: "Repository type",
				Optional:    true,
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

	log.Printf("[DEBUG] Creating application: %s\n", application.Name)
	applicationService := m.(*Services).ApplicationService
	err := applicationService.CreateApplication(application.toClientApplication())
	if err != nil {
		return err
	}

	// Retry until app exists
	_, err = applicationService.GetApplicationByNameWithRetries(application.Name)
	if err != nil {
		if serr, ok := err.(*client.SpinnakerError); ok {
			if serr.Status == 404 {
				d.SetId("")
				return nil
			}
		}

		return err
	}

	d.SetId(application.Name)
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

	log.Printf("[DEBUG] Got application: %s\n", a.Name)
	return fromClientApplication(a).setResourceData(d)
}

func resourceApplicationUpdate(d *schema.ResourceData, m interface{}) error {
	var application application
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &application); err != nil {
		return err
	}

	applicationService := m.(*Services).ApplicationService
	if err := applicationService.UpdateApplication(application.toClientApplication()); err != nil {
		return err
	}

	log.Printf("[DEBUG] Updated application %s\n", d.Id())
	return resourceApplicationRead(d, m)
}

func resourceApplicationDelete(d *schema.ResourceData, m interface{}) error {
	var a application
	configRaw := d.Get("").(map[string]interface{})
	if err := mapstructure.Decode(configRaw, &a); err != nil {
		return err
	}

	log.Printf("[DEBUG] Deleting application %s\n", d.Id())
	applicationService := m.(*Services).ApplicationService
	err := applicationService.DeleteApplication(a.toClientApplication())
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}
