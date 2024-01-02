package ah

import (
	"context"
	"fmt"
	"net/http"
)

// KubernetesWorker object
type KubernetesWorker struct {
	Labels           map[string]string `json:"labels,omitempty"`
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	State            string            `json:"state,omitempty"`
	Type             string            `json:"type,omitempty"`
	CreatedAt        string            `json:"created_at,omitempty"`
	ExternalIpID     string            `json:"external_ip_id,omitempty"`
	PrivateNetworkID string            `json:"private_network_id,omitempty"`
	CloudServerID    string            `json:"cloud_server_id,omitempty"`
}

type ClusterDeleteWorkerRequest struct {
	Replace bool `json:"replace,omitempty"`
}

// DeleteWorker deletes worker pool
func (kcs *KubernetesClustersService) DeleteWorker(ctx context.Context, clusterID, workerPoolID, workerID string, request *ClusterDeleteWorkerRequest) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/worker_pools/%s/workers/%s", clusterID, workerPoolID, workerID)
	req, err := kcs.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if request != nil {
		q := req.URL.Query()
		replace := "0"
		if request.Replace {
			replace = "1"
		}
		q.Add("replace", replace)
		req.URL.RawQuery = q.Encode()

	}
	fmt.Println(req.URL.String())

	if _, err := kcs.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
