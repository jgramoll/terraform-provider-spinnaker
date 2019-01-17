package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDeployStageClusterResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"account": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Account to deploy cluster",
				Required:    true,
			},
			"application": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Application to deploy cluster",
				Required:    true,
			},
			"availability_zones": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Availability zones to deploy cluster",
				MaxItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"us_east_1": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"us_east_2": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"capacity": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Capacity for cluster",
				MaxItems:    1,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"desired": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"min": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"cloud_provider": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Cloud Provider to deploy cluster",
				Required:    true,
			},
			"cooldown": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Cooldown to deploy cluster",
				Optional:    true,
				Default:     10,
			},
			"copy_source_custom_block_device_mappings": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Spinnaker will use the block device mappings of the existing server group when deploying a new server group.\nIn the event that there is no existing server group, the defaults for the selected instance type will be used.",
				Optional:    true,
				Default:     false,
			},
			"ebs_optimized": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "",
				Optional:    true,
				Default:     false,
			},
			"enabled_metrics": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Metrics to be enabled for cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"free_form_details": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Detail is a string of free-form alphanumeric characters and hyphens to describe any other variables.",
				Optional:    true,
			},
			"health_check_grace_period": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Health check grace period for cluster",
				Optional:    true,
				Default:     "300",
			},
			"health_check_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Type of health check for cluster (ELB, EC2)",
				Required:    true,
			},
			"iam_role": &schema.Schema{
				Type:        schema.TypeString,
				Description: "IAM instane profile",
				Optional:    true,
			},
			"instance_monitoring": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Instance Monitoring whether to enable detailed monitoring of instances. Group metrics must be disabled to update an ASG with Instance Monitoring set to false.",
				Optional:    true,
				Default:     false,
			},
			"instance_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Instance Type for cluster",
				Required:    true,
			},
			"key_pair": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Key pair name for cluster",
				Required:    true,
			},
			"load_balancers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Load balancer to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"moniker": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Name to attach to cluster",
				Optional:    true,
				MaxItems:    1,
				Elem:        monikerResource(),
			},
			"provider": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Provider to deploy cluster",
				Required:    true,
			},
			"security_groups": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Security Groups to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"spel_load_balancers": &schema.Schema{
				Type:        schema.TypeList,
				Description: "spel load balancers to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"spel_target_groups": &schema.Schema{
				Type:        schema.TypeList,
				Description: "spel target groups to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"spot_price": &schema.Schema{
				Type:        schema.TypeString,
				Description: "spot price for cluster",
				Optional:    true,
			},
			"stack": &schema.Schema{
				Type:        schema.TypeString,
				Description: "stack name for cluster",
				Optional:    true,
			},
			"strategy": &schema.Schema{
				Type:        schema.TypeString,
				Description: "strategy for deploy (redblack, etc)",
				Required:    true,
			},
			"subnet_type": &schema.Schema{
				Type:        schema.TypeString,
				Description: "subnet to deploy cluster",
				Required:    true,
			},
			"suspended_processes": &schema.Schema{
				Type:        schema.TypeList,
				Description: "subnet to deploy cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": &schema.Schema{
				Type:        schema.TypeMap,
				Description: "tags to put on cluster",
				Optional:    true,
			},
			"target_groups": &schema.Schema{
				Type:        schema.TypeList,
				Description: "target groups to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_healthy_deploy_percentage": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Consider deployment successful when percent of instances are healthy",
				Optional:    true,
				Default:     100,
			},
			"termination_policies": &schema.Schema{
				Type:        schema.TypeList,
				Description: "Termination policy names for cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_ami_block_device_mappings": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Use the block device mappings from the selected AMI when deploying a new server group",
				Optional:    true,
				Default:     false,
			},
			"use_source_capacity": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Spinnaker will use the current capacity of the existing server group when deploying a new server group.\nThis setting is intended to support a server group with auto-scaling enabled, where the bounds and desired capacity are controlled by an external process.\nIn the event that there is no existing server group, the deploy will fail.",
				Optional:    true,
				Default:     false,
			},
		},
	}
}
