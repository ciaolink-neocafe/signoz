package model

// Interface for source of otel collector config recommendations.
type AgentConfigProvider interface {
	// Generate recommended config for an agent based on its `currentConfYaml`
	// and current state of user facing settings for agent based features.
	RecommendAgentConfig(currentConfYaml []byte) (
		recommendedConfYaml []byte,
		// Opaque id of the recommended config, used for reporting deployment status updates
		configId string,
		err error,
	)

	// Report deployment status for config recommendations generated by RecommendAgentConfig
	ReportConfigDeploymentStatus(
		agentId string,
		configId string,
		err error,
	)
}
