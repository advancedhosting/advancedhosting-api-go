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

// ClusterNodePool object
type ClusterNodePool struct {
	Labels            map[string]string `json:"labels,omitempty"`
	ID                string            `json:"id,omitempty"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	CreatedAt         string            `json:"created_at,omitempty"`
	Nodes             []ClusterNodes    `json:"nodes,omitempty"`
	PrivateProperties PrivateProperties `json:"private_properties,omitempty"`
	PublicProperties  PublicProperties  `json:"public_properties,omitempty"`
	Count             int               `json:"count,omitempty"`
	AutoScale         bool              `json:"autoscale,omitempty"`
	MinCount          int               `json:"min_count,omitempty"`
	MaxCount          int               `json:"max_count,omitempty"`
}

type ClusterNodePoolRoot struct {
	ClusterNodePool *ClusterNodePool `json:"node_pool,omitempty"`
}

type ClusterNodePoolsRoot struct {
	ClusterNodePools []ClusterNodePool `json:"node_pools,omitempty"`
}

// CreateClusterNodePoolRequest represents a request to create a node pool.
type CreateClusterNodePoolRequest struct {
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

// UpdateClusterNodePoolRequest represents a request to update a node pool
type UpdateClusterNodePoolRequest struct {
	Labels    map[string]string `json:"labels,omitempty"`
	Name      string            `json:"name,omitempty"`
	Count     int               `json:"count,omitempty"`
	AutoScale bool              `json:"autoscale,omitempty"`
	MinCount  int               `json:"min_count,omitempty"`
	MaxCount  int               `json:"max_count,omitempty"`
}

// GetNodePool returns node pool
func (kc *ClustersService) GetNodePool(ctx context.Context, clusterId, nodePoolId string) (*ClusterNodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s", clusterId, nodePoolId)
	req, err := kc.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	var nodePoolRoot ClusterNodePoolRoot

	if _, err = kc.client.Do(ctx, req, &nodePoolRoot); err != nil {
		return nil, err
	}

	return nodePoolRoot.ClusterNodePool, nil
}

// ListNodePools returns list of node pools
func (kc *ClustersService) ListNodePools(ctx context.Context, options *ListOptions, clusterId string) ([]ClusterNodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools", clusterId)

	var NodePoolsRoot ClusterNodePoolsRoot

	if err := kc.client.list(ctx, path, options, &NodePoolsRoot); err != nil {
		return nil, err
	}

	return NodePoolsRoot.ClusterNodePools, nil
}

// CreateNodePool creates node pool
func (kc *ClustersService) CreateNodePool(ctx context.Context, clusterId string, request *CreateClusterNodePoolRequest) (*ClusterNodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools", clusterId)
	req, err := kc.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var nodePoolRoot ClusterNodePoolRoot
	if _, err := kc.client.Do(ctx, req, &nodePoolRoot); err != nil {
		return nil, err
	}

	return nodePoolRoot.ClusterNodePool, nil
}

// UpdateNodePool updates node pool
func (kc *ClustersService) UpdateNodePool(ctx context.Context, clusterId, nodePoolId string, request *UpdateClusterNodePoolRequest) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s", clusterId, nodePoolId)
	req, err := kc.client.newRequest(http.MethodPut, path, request)
	if err != nil {
		return err
	}

	if _, err := kc.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteNodePool deletes node pool
func (kc *ClustersService) DeleteNodePool(ctx context.Context, clusterId string, nodePoolId string, replace bool) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s?replace=%v", clusterId, nodePoolId, replace)
	req, err := kc.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if _, err := kc.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
