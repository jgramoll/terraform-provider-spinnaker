package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func pipelineDeployStageClusterResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"account": {
				Type:        schema.TypeString,
				Description: "Account to deploy cluster",
				Required:    true,
			},
			"application": {
				Type:        schema.TypeString,
				Description: "Application to deploy cluster",
				Required:    true,
			},
			"availability_zones": {
				Type:        schema.TypeList,
				Description: "Availability zones to deploy cluster",
				MaxItems:    1,
				Required:    true,
				Elem:        availabilityZonesResource(),
			},
			"capacity": {
				Type:        schema.TypeList,
				Description: "Capacity for cluster",
				MaxItems:    1,
				Optional:    true,
				Elem:        capacityResource(),
			},
			"cloud_provider": {
				Type:        schema.TypeString,
				Description: "Cloud Provider to deploy cluster",
				Required:    true,
			},
			"cooldown": {
				Type:        schema.TypeInt,
				Description: "Cooldown to deploy cluster",
				Optional:    true,
				Default:     10,
			},
			"copy_source_custom_block_device_mappings": {
				Type:        schema.TypeBool,
				Description: "Spinnaker will use the block device mappings of the existing server group when deploying a new server group.\nIn the event that there is no existing server group, the defaults for the selected instance type will be used.",
				Optional:    true,
				Default:     false,
			},
			"dirty": {
				Type:        schema.TypeMap,
				Description: "",
				Optional:    true,
			},
			"ebs_optimized": {
				Type:        schema.TypeBool,
				Description: "",
				Optional:    true,
				Default:     false,
			},
			"enabled_metrics": {
				Type:        schema.TypeList,
				Description: "Metrics to be enabled for cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"free_form_details": {
				Type:        schema.TypeString,
				Description: "Detail is a string of free-form alphanumeric characters and hyphens to describe any other variables.",
				Optional:    true,
			},
			"health_check_grace_period": {
				Type:        schema.TypeString,
				Description: "Health check grace period for cluster",
				Optional:    true,
				Default:     "300",
			},
			"health_check_type": {
				Type:        schema.TypeString,
				Description: "Type of health check for cluster (ELB, EC2)",
				Required:    true,
			},
			"iam_role": {
				Type:        schema.TypeString,
				Description: "IAM instance profile",
				Optional:    true,
			},
			"instance_monitoring": {
				Type:        schema.TypeBool,
				Description: "Instance Monitoring whether to enable detailed monitoring of instances. Group metrics must be disabled to update an ASG with Instance Monitoring set to false.",
				Optional:    true,
				Default:     false,
			},
			"instance_type": {
				Type:        schema.TypeString,
				Description: "Instance Type for cluster",
				Required:    true,
			},
			"key_pair": {
				Type:        schema.TypeString,
				Description: "Key pair name for cluster",
				Required:    true,
			},
			"max_remaining_asgs": {
				Type:        schema.TypeInt,
				Description: "Max amount of asgs to run",
				Optional:    true,
				Default:     2,
			},
			"load_balancers": {
				Type:        schema.TypeList,
				Description: "Load balancer to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"moniker": {
				Type:        schema.TypeList,
				Description: "Name to attach to cluster",
				Optional:    true,
				MaxItems:    1,
				Elem:        monikerResource(),
			},
			"provider": {
				Type:        schema.TypeString,
				Description: "Provider to deploy cluster",
				Required:    true,
			},
			"security_groups": {
				Type:        schema.TypeList,
				Description: "Security Groups to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"security_groups_expression": {
				Type:        schema.TypeString,
				Description: "Security Group expression -- will override other sg inputs",
				Optional:    true,
			},
			"spel_load_balancers": {
				Type:        schema.TypeList,
				Description: "spel load balancers to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"spel_target_groups": {
				Type:        schema.TypeList,
				Description: "spel target groups to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"spot_price": {
				Type:        schema.TypeString,
				Description: "spot price for cluster",
				Optional:    true,
			},
			"stack": {
				Type:        schema.TypeString,
				Description: "stack name for cluster",
				Optional:    true,
			},
			"strategy": {
				Type:        schema.TypeString,
				Description: "strategy for deploy (redblack, etc)",
				Required:    true,
			},
			"subnet_type": {
				Type:        schema.TypeString,
				Description: "subnet to deploy cluster",
				Required:    true,
			},
			"suspended_processes": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:        schema.TypeMap,
				Description: "tags to put on cluster",
				Optional:    true,
			},
			"target_groups": {
				Type:        schema.TypeList,
				Description: "target groups to attach to cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"target_healthy_deploy_percentage": {
				Type:        schema.TypeInt,
				Description: "Consider deployment successful when percent of instances are healthy",
				Optional:    true,
				Default:     100,
			},
			"termination_policies": {
				Type:        schema.TypeList,
				Description: "Termination policy names for cluster",
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_ami_block_device_mappings": {
				Type:        schema.TypeBool,
				Description: "Use the block device mappings from the selected AMI when deploying a new server group",
				Optional:    true,
				Default:     false,
			},
			"use_source_capacity": {
				Type:        schema.TypeBool,
				Description: "Spinnaker will use the current capacity of the existing server group when deploying a new server group.\nThis setting is intended to support a server group with auto-scaling enabled, where the bounds and desired capacity are controlled by an external process.\nIn the event that there is no existing server group, the deploy will fail.",
				Optional:    true,
				Default:     false,
			},
			"user_data": {
				Type:        schema.TypeString,
				Description: "UserData is a base64 encoded string.",
				Optional:    true,
			},
		},
	}
}
