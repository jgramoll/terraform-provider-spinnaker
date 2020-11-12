package client

type DeployCloudformationSource int

const (
	DeployCloudformationSourceUnknown DeployCloudformationSource = iota
	DeployCloudformationSourceText
	DeployCloudformationSourceArtifact
)
