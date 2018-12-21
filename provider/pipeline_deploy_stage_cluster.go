package provider

import (
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployStageClusterCapacity struct {
	Desired int `mapstructure:"desired"`
	Max     int `mapstructure:"max"`
	Min     int `mapstructure:"min"`
}

type deployStageClusterMoniker struct {
	App    string `mapstructure:"app"`
	Detail string `mapstructure:"detail"`
	Stack  string `mapstructure:"stack"`
}

type deployStageCluster struct {
	Account                             string                       `mapstructure:"account"`
	Application                         string                       `mapstructure:"application"`
	AvailabilityZones                   []map[string][]string        `mapstructure:"availability_zones"`
	Capacity                            []deployStageClusterCapacity `mapstructure:"capacity"`
	CloudProvider                       string                       `mapstructure:"cloud_provider"`
	Cooldown                            int                          `mapstructure:"cooldown"`
	CopySourceCustomBlockDeviceMappings bool                         `mapstructure:"copy_source_custom_block_device_mappings"`
	EBSOptimized                        bool                         `mapstructure:"ebs_optimized"`
	EnabledMetrics                      []string                     `mapstructure:"enabled_metrics"`
	FreeFormDetails                     string                       `mapstructure:"free_form_details"`
	HealthCheckGracePeriod              int                          `mapstructure:"health_check_grace_period"`
	HealthCheckType                     string                       `mapstructure:"health_check_type"`
	IAMRole                             string                       `mapstructure:"iam_role"`
	InstanceMonitoring                  bool                         `mapstructure:"instance_monitoring"`
	InstanceType                        string                       `mapstructure:"instance_type"`
	KeyPair                             string                       `mapstructure:"key_pair"`
	LoadBalancers                       []string                     `mapstructure:"load_balancers"`
	Moniker                             []deployStageClusterMoniker  `mapstructure:"moniker"`
	Provider                            string                       `mapstructure:"provider"`
	SecurityGroups                      []string                     `mapstructure:"security_groups"`
	SpelLoadBalancers                   []string                     `mapstructure:"spel_load_balancers"`
	SpelTargetGroups                    []string                     `mapstructure:"spel_target_groups"`
	SpotPrice                           string                       `mapstructure:"spot_price"`
	Stack                               string                       `mapstructure:"stack"`
	Strategy                            string                       `mapstructure:"strategy"`
	SubnetType                          string                       `mapstructure:"subnet_type"`
	SuspendedProcesses                  []string                     `mapstructure:"suspended_processes"`
	Tags                                map[string]string            `mapstructure:"tags"`
	TargetGroups                        []string                     `mapstructure:"target_groups"`
	TargetHealthyDeployPercentage       int                          `mapstructure:"target_healthy_deploy_percentage"`
	TerminationPolicies                 []string                     `mapstructure:"termination_policies"`
	UseAmiBlockDeviceMappings           bool                         `mapstructure:"use_ami_block_device_mappings"`
	UseSourceCapacity                   bool                         `mapstructure:"use_source_capacity"`
}

func (c *deployStageCluster) clientCapacity() client.DeployStageClusterCapacity {
	if len(c.Capacity) > 0 {
		return client.DeployStageClusterCapacity(c.Capacity[0])
	}
	return client.DeployStageClusterCapacity{}
}

func (c *deployStageCluster) clientMoniker() client.DeployStageClusterMoniker {
	if len(c.Moniker) > 0 {
		return client.DeployStageClusterMoniker(c.Moniker[0])
	}
	return client.DeployStageClusterMoniker{}
}

func (c *deployStageCluster) toClientCluster() *client.DeployStageCluster {
	// TODO better way?
	return &client.DeployStageCluster{
		Account:           c.Account,
		Application:       c.Application,
		AvailabilityZones: c.AvailabilityZones[0],
		Capacity:          c.clientCapacity(),
		CloudProvider:     c.CloudProvider,
		Cooldown:          c.Cooldown,

		CopySourceCustomBlockDeviceMappings: c.CopySourceCustomBlockDeviceMappings,

		EBSOptimized:           c.EBSOptimized,
		EnabledMetrics:         c.EnabledMetrics,
		FreeFormDetails:        c.FreeFormDetails,
		HealthCheckGracePeriod: c.HealthCheckGracePeriod,
		HealthCheckType:        c.HealthCheckType,
		IAMRole:                c.IAMRole,
		InstanceMonitoring:     c.InstanceMonitoring,
		InstanceType:           c.InstanceType,
		KeyPair:                c.KeyPair,
		LoadBalancers:          c.LoadBalancers,
		Moniker:                c.clientMoniker(),
		Provider:               c.Provider,
		SecurityGroups:         c.SecurityGroups,
		SpelLoadBalancers:      c.SpelLoadBalancers,
		SpelTargetGroups:       c.SpelTargetGroups,
		SpotPrice:              c.SpotPrice,
		Stack:                  c.Stack,
		Strategy:               c.Strategy,
		SuspendedProcesses:     c.SuspendedProcesses,
		Tags:                   c.Tags,
		TargetGroups:           c.TargetGroups,

		TargetHealthyDeployPercentage: c.TargetHealthyDeployPercentage,
		TerminationPolicies:           c.TerminationPolicies,
		UseAmiBlockDeviceMappings:     c.UseAmiBlockDeviceMappings,
		UseSourceCapacity:             c.UseSourceCapacity,
	}
}
