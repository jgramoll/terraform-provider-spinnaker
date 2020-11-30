package provider

import (
	"log"
	"strings"

	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

type deployStageClusters []*deployStageCluster

type deployStageCluster struct {
	Account                             string                 `mapstructure:"account"`
	Application                         string                 `mapstructure:"application"`
	AvailabilityZones                   []map[string][]string  `mapstructure:"availability_zones"`
	Capacity                            *[]*capacity           `mapstructure:"capacity"`
	CloudProvider                       string                 `mapstructure:"cloud_provider"`
	Cooldown                            int                    `mapstructure:"cooldown"`
	CopySourceCustomBlockDeviceMappings bool                   `mapstructure:"copy_source_custom_block_device_mappings"`
	Dirty                               map[string]interface{} `mapstructure:"dirty"`
	EBSOptimized                        bool                   `mapstructure:"ebs_optimized"`
	EnabledMetrics                      []string               `mapstructure:"enabled_metrics"`
	FreeFormDetails                     string                 `mapstructure:"free_form_details"`
	HealthCheckGracePeriod              string                 `mapstructure:"health_check_grace_period"`
	HealthCheckType                     string                 `mapstructure:"health_check_type"`
	IAMRole                             string                 `mapstructure:"iam_role"`
	InstanceMonitoring                  bool                   `mapstructure:"instance_monitoring"`
	InstanceType                        string                 `mapstructure:"instance_type"`
	KeyPair                             string                 `mapstructure:"key_pair"`
	MaxRemainingAsgs                    int                    `mapstructure:"max_remaining_asgs"`
	LoadBalancers                       []string               `mapstructure:"load_balancers"`
	Moniker                             *[]*moniker            `mapstructure:"moniker"`
	Provider                            string                 `mapstructure:"provider"`
	SecurityGroups                      interface{}            `mapstructure:"security_groups"`
	SecurityGroupsExpression            string                 `mapstructure:"security_groups_expression"`
	SpelLoadBalancers                   []string               `mapstructure:"spel_load_balancers"`
	SpelTargetGroups                    []string               `mapstructure:"spel_target_groups"`
	SpotPrice                           string                 `mapstructure:"spot_price"`
	Stack                               string                 `mapstructure:"stack"`
	Strategy                            string                 `mapstructure:"strategy"`
	SubnetType                          string                 `mapstructure:"subnet_type"`
	SuspendedProcesses                  []string               `mapstructure:"suspended_processes"`
	Tags                                map[string]string      `mapstructure:"tags"`
	TargetGroups                        []string               `mapstructure:"target_groups"`
	TargetHealthyDeployPercentage       int                    `mapstructure:"target_healthy_deploy_percentage"`
	TerminationPolicies                 []string               `mapstructure:"termination_policies"`
	Base64UserData                      string                 `mapstructure:"user_data"`
	UseAmiBlockDeviceMappings           bool                   `mapstructure:"use_ami_block_device_mappings"`
	UseSourceCapacity                   bool                   `mapstructure:"use_source_capacity"`
}

func (c *deployStageCluster) clientAvailabilityZones() *map[string][]string {
	newAZ := map[string][]string{}
	for _, regions := range c.AvailabilityZones {
		for region, zones := range regions {
			if len(zones) == 0 {
				continue
			}
			// TODO unit test
			newAZ[strings.Replace(region, "_", "-", -1)] = zones
		}
	}
	return &newAZ
}

func (c *deployStageCluster) importAvailabilityZones(clientCluster *client.DeploymentCluster) {
	for region, zones := range clientCluster.AvailabilityZones {
		newZone := map[string][]string{
			strings.Replace(region, "-", "_", -1): zones,
		}
		// TODO unit test
		c.AvailabilityZones = append(c.AvailabilityZones, newZone)
	}
}

func (c *deployStageCluster) toClientCluster() *client.DeploymentCluster {
	// TODO better way?
	clientCluster := client.NewDeploymentCluster()
	clientCluster.Account = c.Account
	clientCluster.Application = c.Application
	clientCluster.AvailabilityZones = *c.clientAvailabilityZones()
	clientCluster.Capacity = toClientCapacity(c.Capacity)
	clientCluster.CloudProvider = c.CloudProvider
	clientCluster.Cooldown = c.Cooldown
	clientCluster.CopySourceCustomBlockDeviceMappings = c.CopySourceCustomBlockDeviceMappings
	clientCluster.EBSOptimized = c.EBSOptimized
	clientCluster.EnabledMetrics = c.EnabledMetrics
	clientCluster.FreeFormDetails = c.FreeFormDetails
	clientCluster.HealthCheckGracePeriod = c.HealthCheckGracePeriod
	clientCluster.HealthCheckType = c.HealthCheckType
	clientCluster.IAMRole = c.IAMRole
	clientCluster.InstanceMonitoring = c.InstanceMonitoring
	clientCluster.InstanceType = c.InstanceType
	clientCluster.KeyPair = c.KeyPair
	clientCluster.MaxRemainingAsgs = c.MaxRemainingAsgs
	clientCluster.LoadBalancers = c.LoadBalancers
	clientCluster.Moniker = toClientMoniker(c.Moniker)
	clientCluster.Provider = c.Provider
	if c.SecurityGroupsExpression != "" {
		clientCluster.SecurityGroups = []string{c.SecurityGroupsExpression}
	} else {
		clientCluster.SecurityGroups = c.SecurityGroups
	}
	clientCluster.SpelLoadBalancers = c.SpelLoadBalancers
	clientCluster.SpelTargetGroups = c.SpelTargetGroups
	clientCluster.SpotPrice = c.SpotPrice
	clientCluster.Stack = c.Stack
	clientCluster.Strategy = c.Strategy
	clientCluster.SubnetType = c.SubnetType
	clientCluster.SuspendedProcesses = c.SuspendedProcesses
	clientCluster.Tags = c.Tags
	clientCluster.TargetGroups = c.TargetGroups
	clientCluster.TargetHealthyDeployPercentage = c.TargetHealthyDeployPercentage
	clientCluster.TerminationPolicies = c.TerminationPolicies
	clientCluster.UseAmiBlockDeviceMappings = c.UseAmiBlockDeviceMappings
	clientCluster.UseSourceCapacity = c.UseSourceCapacity
	clientCluster.Base64UserData = c.Base64UserData
	return clientCluster
}

func (s *deployStageClusters) toClientClusters() *[]*client.DeploymentCluster {
	if len(*s) == 0 {
		return nil
	}
	clusters := []*client.DeploymentCluster{}
	for _, c := range *s {
		clusters = append(clusters, c.toClientCluster())
	}
	return &clusters
}

func fromClientCluster(c *client.DeploymentCluster) *deployStageCluster {
	var sgs []string
	var sgExpression string
	switch v := c.SecurityGroups.(type) {
	default:
		log.Printf("[WARN] unknown security group type: %s\n", v)
	case string:
		sgExpression = v
	case []interface{}:
		varray := []string{}
		for _, i := range v {
			s, ok := i.(string)
			if ok {
				varray = append(varray, s)
			} else {
				log.Printf("[WARN] unknown security group type, should be string: %v\n", i)
			}
		}
		sgs = varray
	case []string:
		sgs = v
	}

	newCluster := deployStageCluster{
		Account:       c.Account,
		Application:   c.Application,
		CloudProvider: c.CloudProvider,
		Cooldown:      c.Cooldown,

		CopySourceCustomBlockDeviceMappings: c.CopySourceCustomBlockDeviceMappings,

		EBSOptimized:             c.EBSOptimized,
		EnabledMetrics:           c.EnabledMetrics,
		FreeFormDetails:          c.FreeFormDetails,
		HealthCheckGracePeriod:   c.HealthCheckGracePeriod,
		HealthCheckType:          c.HealthCheckType,
		IAMRole:                  c.IAMRole,
		InstanceMonitoring:       c.InstanceMonitoring,
		InstanceType:             c.InstanceType,
		KeyPair:                  c.KeyPair,
		MaxRemainingAsgs:         c.MaxRemainingAsgs,
		LoadBalancers:            c.LoadBalancers,
		Provider:                 c.Provider,
		SecurityGroups:           sgs,
		SecurityGroupsExpression: sgExpression,
		SpelLoadBalancers:        c.SpelLoadBalancers,
		SpelTargetGroups:         c.SpelTargetGroups,
		SpotPrice:                c.SpotPrice,
		Stack:                    c.Stack,
		Strategy:                 c.Strategy,
		SubnetType:               c.SubnetType,
		SuspendedProcesses:       c.SuspendedProcesses,
		Tags:                     c.Tags,
		TargetGroups:             c.TargetGroups,

		TargetHealthyDeployPercentage: c.TargetHealthyDeployPercentage,
		TerminationPolicies:           c.TerminationPolicies,
		UseAmiBlockDeviceMappings:     c.UseAmiBlockDeviceMappings,
		UseSourceCapacity:             c.UseSourceCapacity,
		Base64UserData:                c.Base64UserData,
	}
	newCluster.importAvailabilityZones(c)
	newCluster.Capacity = fromClientCapacity(c.Capacity)
	newCluster.Moniker = fromClientMoniker(c.Moniker)
	return &newCluster
}

func (*deployStageClusters) fromClientClusters(clientClusters *[]*client.DeploymentCluster) *deployStageClusters {
	if clientClusters == nil || len(*clientClusters) == 0 {
		return nil
	}

	newClusters := deployStageClusters{}
	for _, c := range *clientClusters {
		cluster := fromClientCluster(c)
		newClusters = append(newClusters, cluster)
	}
	return &newClusters
}
