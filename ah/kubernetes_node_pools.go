package ah

import (
	"context"
	"fmt"
	"net/http"
)

// PublicProperties object
type PublicProperties struct {
	PlanID int `json:"plan_id,omitempty"`
}

// PrivateProperties object
type PrivateProperties struct {
	NetworkID     string `json:"network_id,omitempty"`
	ClusterID     string `json:"cluster_id,omitempty"`
	ClusterNodeID string `json:"cluster_node_id,omitempty"`
	Vcpu          int    `json:"vcpu,omitempty"`
	Ram           int    `json:"ram,omitempty"`
	Disk          int    `json:"disk,omitempty"`
}

// KubernetesNodePool object
type KubernetesNodePool struct {
	Labels            map[string]string `json:"labels,omitempty"`
	ID                string            `json:"id,omitempty"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	CreatedAt         string            `json:"created_at,omitempty"`
	Nodes             []KubernetesNodes `json:"nodes,omitempty"`
	PrivateProperties PrivateProperties `json:"private_properties,omitempty"`
	PublicProperties  PublicProperties  `json:"public_properties,omitempty"`
	Count             int               `json:"count,omitempty"`
	AutoScale         bool              `json:"autoscale,omitempty"`
	MinCount          int               `json:"min_count,omitempty"`
	MaxCount          int               `json:"max_count,omitempty"`
}

type KubernetesNodePoolRoot struct {
	KubernetesNodePool *KubernetesNodePool `json:"node_pool,omitempty"`
}

type KubernetesNodePoolsRoot struct {
	KubernetesNodePools []KubernetesNodePool `json:"node_pools,omitempty"`
}

// CreateKubernetesNodePoolRequest represents a request to create a node pool.
type CreateKubernetesNodePoolRequest struct {
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	PrivateProperties PrivateProperties `json:"private_properties,omitempty"`
	PublicProperties  PublicProperties  `json:"public_properties,omitempty"`
	Count             int               `json:"count,omitempty"`
	MinCount          int               `json:"min_count,omitempty"`
	MaxCount          int               `json:"max_count,omitempty"`
	AutoScale         bool              `json:"autoscale,omitempty"`
}

// UpdateKubernetesNodePoolRequest represents a request to update a node pool
type UpdateKubernetesNodePoolRequest struct {
	Labels    map[string]string `json:"labels,omitempty"`
	Name      string            `json:"name,omitempty"`
	Count     int               `json:"count,omitempty"`
	AutoScale bool              `json:"autoscale,omitempty"`
	MinCount  int               `json:"min_count,omitempty"`
	MaxCount  int               `json:"max_count,omitempty"`
}

// GetNodePool returns node pool
func (kcs *KubernetesClustersService) GetNodePool(ctx context.Context, clusterId, nodePoolId string) (*KubernetesNodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s", clusterId, nodePoolId)
	req, err := kcs.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	var nodePoolRoot KubernetesNodePoolRoot

	if _, err = kcs.client.Do(ctx, req, &nodePoolRoot); err != nil {
		return nil, err
	}

	return nodePoolRoot.KubernetesNodePool, nil
}

// ListNodePools returns list of node pools
func (kcs *KubernetesClustersService) ListNodePools(ctx context.Context, options *ListOptions, clusterId string) ([]KubernetesNodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools", clusterId)

	var NodePoolsRoot KubernetesNodePoolsRoot

	if err := kcs.client.list(ctx, path, options, &NodePoolsRoot); err != nil {
		return nil, err
	}

	return NodePoolsRoot.KubernetesNodePools, nil
}

// CreateNodePool creates node pool
func (kcs *KubernetesClustersService) CreateNodePool(ctx context.Context, clusterId string, request *CreateKubernetesNodePoolRequest) (*KubernetesNodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools", clusterId)
	req, err := kcs.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var nodePoolRoot KubernetesNodePoolRoot
	if _, err := kcs.client.Do(ctx, req, &nodePoolRoot); err != nil {
		return nil, err
	}

	return nodePoolRoot.KubernetesNodePool, nil
}

// UpdateNodePool updates node pool
func (kcs *KubernetesClustersService) UpdateNodePool(ctx context.Context, clusterId, nodePoolId string, request *UpdateKubernetesNodePoolRequest) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s", clusterId, nodePoolId)
	req, err := kcs.client.newRequest(http.MethodPut, path, request)
	if err != nil {
		return err
	}

	if _, err := kcs.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteNodePool deletes node pool
func (kcs *KubernetesClustersService) DeleteNodePool(ctx context.Context, clusterId string, nodePoolId string, replace bool) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s?replace=%v", clusterId, nodePoolId, replace)
	req, err := kcs.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if _, err := kcs.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
