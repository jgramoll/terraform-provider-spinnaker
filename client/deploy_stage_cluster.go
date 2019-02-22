package client

//DeployStageClusterCapacity capacity for cluster
type DeployStageClusterCapacity struct {
	Desired int `json:"desired"`
	Max     int `json:"max"`
	Min     int `json:"min"`
}

// DeployStageCluster cluster to deploy
type DeployStageCluster struct {
	Account                             string                      `json:"account"`
	Application                         string                      `json:"application"`
	AvailabilityZones                   map[string][]string         `json:"availabilityZones"`
	Capacity                            *DeployStageClusterCapacity `json:"capacity"`
	CloudProvider                       string                      `json:"cloudProvider"`
	Cooldown                            int                         `json:"cooldown"`
	CopySourceCustomBlockDeviceMappings bool                        `json:"copySourceCustomBlockDeviceMappings"`
	Dirty                               map[string]interface{}      `json:"dirty"`
	EBSOptimized                        bool                        `json:"ebsOptimized"`
	EnabledMetrics                      []string                    `json:"enabledMetrics"`
	FreeFormDetails                     string                      `json:"freeFormDetails"`
	HealthCheckGracePeriod              string                      `json:"healthCheckGracePeriod"`
	HealthCheckType                     string                      `json:"healthCheckType"`
	IAMRole                             string                      `json:"iamRole"`
	InstanceMonitoring                  bool                        `json:"instanceMonitoring"`
	InstanceType                        string                      `json:"instanceType"`
	KeyPair                             string                      `json:"keyPair"`
	LoadBalancers                       []string                    `json:"loadBalancers"`
	Moniker                             *Moniker                    `json:"moniker"`
	Provider                            string                      `json:"provider"`
	SecurityGroups                      []string                    `json:"securityGroups"`
	SpelLoadBalancers                   []string                    `json:"spelLoadBalancers"`
	SpelTargetGroups                    []string                    `json:"spelTargetGroups"`
	SpotPrice                           string                      `json:"spotPrice"`
	Stack                               string                      `json:"stack"`
	Strategy                            string                      `json:"strategy"`
	SubnetType                          string                      `json:"subnetType"`
	SuspendedProcesses                  []string                    `json:"suspendedProcesses"`
	Tags                                map[string]string           `json:"tags"`
	TargetGroups                        []string                    `json:"targetGroups"`
	TargetHealthyDeployPercentage       int                         `json:"targetHealthyDeployPercentage"`
	TerminationPolicies                 []string                    `json:"terminationPolicies"`
	UseAmiBlockDeviceMappings           bool                        `json:"useAmiBlockDeviceMappings"`
	UseSourceCapacity                   bool                        `json:"useSourceCapacity"`
}
