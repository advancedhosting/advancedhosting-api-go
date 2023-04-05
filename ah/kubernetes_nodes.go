package ah

// KubernetesNodes object
type KubernetesNodes struct {
	Labels           map[string]string `json:"labels,omitempty"`
	Id               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	State            string            `json:"state,omitempty"`
	Type             string            `json:"type,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	ExternalIpID     string            `json:"external_ip_id,omitempty"`
	PrivateNetworkID string            `json:"private_network_id,omitempty"`
	CloudServerID    string            `json:"cloud_server_id,omitempty"`
}
