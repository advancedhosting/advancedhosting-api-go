package ah

type Nodes struct {
	Id               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	State            string            `json:"state,omitempty"`
	Type             string            `json:"type,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	ExternalIpId     string            `json:"external_ip_id,omitempty"`
	PrivateNetworkId string            `json:"private_network_id,omitempty"`
	CloudServerId    string            `json:"cloud_server_id,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
}
