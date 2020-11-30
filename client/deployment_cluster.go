package client

// DeploymentCluster cluster to deploy
type DeploymentCluster struct {
	Account                             string                 `json:"account"`
	Application                         string                 `json:"application"`
	AvailabilityZones                   map[string][]string    `json:"availabilityZones"`
	Capacity                            *Capacity              `json:"capacity"`
	CloudProvider                       string                 `json:"cloudProvider"`
	Cooldown                            int                    `json:"cooldown"`
	CopySourceCustomBlockDeviceMappings bool                   `json:"copySourceCustomBlockDeviceMappings"`
	DelayBeforeDisableSec               int                    `json:"delayBeforeDisableSec"`
	DelayBeforeScaleDownSec             int                    `json:"delayBeforeScaleDownSec"`
	Dirty                               map[string]interface{} `json:"dirty"`
	EBSOptimized                        bool                   `json:"ebsOptimized"`
	EnabledMetrics                      []string               `json:"enabledMetrics"`
	FreeFormDetails                     string                 `json:"freeFormDetails"`
	HealthCheckGracePeriod              string                 `json:"healthCheckGracePeriod"`
	HealthCheckType                     string                 `json:"healthCheckType"`
	IAMRole                             string                 `json:"iamRole"`
	InstanceMonitoring                  bool                   `json:"instanceMonitoring"`
	InstanceType                        string                 `json:"instanceType"`
	KeyPair                             string                 `json:"keyPair"`
	LoadBalancers                       []string               `json:"loadBalancers"`
	MaxRemainingAsgs                    int                    `json:"maxRemainingAsgs"`
	Moniker                             *Moniker               `json:"moniker"`
	Provider                            string                 `json:"provider"`
	Rollback                            *Rollback              `json:"rollback"`
	ScaleDown                           bool                   `json:"scaleDown"`
	SecurityGroups                      interface{}            `json:"securityGroups"`
	SpelLoadBalancers                   []string               `json:"spelLoadBalancers"`
	SpelTargetGroups                    []string               `json:"spelTargetGroups"`
	SpotPrice                           string                 `json:"spotPrice"`
	Stack                               string                 `json:"stack"`
	Strategy                            string                 `json:"strategy"`
	SubnetType                          string                 `json:"subnetType"`
	SuspendedProcesses                  []string               `json:"suspendedProcesses"`
	Tags                                map[string]string      `json:"tags"`
	TargetGroups                        []string               `json:"targetGroups"`
	TargetHealthyDeployPercentage       int                    `json:"targetHealthyDeployPercentage"`
	TerminationPolicies                 []string               `json:"terminationPolicies"`
	UseAmiBlockDeviceMappings           bool                   `json:"useAmiBlockDeviceMappings"`
	UseSourceCapacity                   bool                   `json:"useSourceCapacity"`
	Base64UserData                      string                 `json:"base64UserData"`
}

// NewDeploymentCluster new cluster
func NewDeploymentCluster() *DeploymentCluster {
	return &DeploymentCluster{
		DelayBeforeDisableSec:   0,
		DelayBeforeScaleDownSec: 0,
		Dirty:                   map[string]interface{}{},
		EnabledMetrics:          []string{},
		LoadBalancers:           []string{},
		MaxRemainingAsgs:        2,
		Rollback:                NewRollback(),
		ScaleDown:               false,
		SpelLoadBalancers:       []string{},
		SpelTargetGroups:        []string{},
		SuspendedProcesses:      []string{},
	}
}
