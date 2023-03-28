package ah

import (
	"context"
	"fmt"
	"net/http"
)

// PublicProperties object
type PublicProperties struct {
	PlanId int `json:"plan_id,omitempty"`
}

// PrivateProperties object
type PrivateProperties struct {
	NetworlId     string `json:"network_id,omitempty"`
	ClusterId     string `json:"cluster_id,omitempty"`
	ClusterNodeId string `json:"cluster_node_id,omitempty"`
	Vcpu          int    `json:"vcpu,omitempty"`
	Ram           int    `json:"ram,omitempty"`
	Disk          int    `json:"disk,omitempty"`
}

// NodePool object
type NodePool struct {
	Labels            map[string]string `json:"labels,omitempty"`
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	CreatedAt         string            `json:"created_at"`
	Nodes             []Nodes           `json:"nodes,omitempty"`
	PrivateProperties PrivateProperties `json:"private_properties,omitempty"`
	PublicProperties  PublicProperties  `json:"public_properties,omitempty"`
	Count             int               `json:"count,omitempty"`
	Autoscale         bool              `json:"autoscale,omitempty"`
	MinCount          int               `json:"min_count,omitempty"`
	MaxCount          int               `json:"max_count,omitempty"`
}

type NodePoolRoot struct {
	NodePool *NodePool `json:"node_pool,omitempty"`
}

type NodePoolsRoot struct {
	NodePools []NodePool `json:"node_pools,omitempty"`
}

// CreateNodePoolRequest represents a request to create a node pool.
type CreateNodePoolRequest struct {
	Labels            map[string]string `json:"labels,omitempty"`
	Name              string            `json:"name"`
	Type              string            `json:"type"`
	PrivateProperties PrivateProperties `json:"private_properties,omitempty"`
	PublicProperties  PublicProperties  `json:"public_properties,omitempty"`
	Count             int               `json:"count,omitempty"`
	MinCount          int               `json:"min_count,omitempty"`
	MaxCount          int               `json:"max_count,omitempty"`
	Autoscale         bool              `json:"autoscale,omitempty"`
}

// NodePoolUpdateRequest represents a request to update a node pool
type NodePoolUpdateRequest struct {
	Labels    map[string]string `json:"labels,omitempty"`
	Name      string            `json:"name,omitempty"`
	Count     int               `json:"count,omitempty"`
	Autoscale bool              `json:"autoscale,omitempty"`
	MinCount  int               `json:"min_count,omitempty"`
	MaxCount  int               `json:"max_count,omitempty"`
}

// NodePoolsAPI is an interface for node pools API.
type NodePoolsAPI interface {
	Get(context.Context, string, string) (*NodePool, error)
	List(context.Context, *ListOptions, string) ([]NodePool, error)
	Create(context.Context, string, *CreateNodePoolRequest) (*NodePool, error)
	Update(context.Context, string, string, *NodePoolUpdateRequest) error
	Delete(context.Context, string, string, bool) error
}

// NodePoolsService implements NodePoolsAPI interface.
type NodePoolsService struct {
	client *APIClient
}

func (nps *NodePoolsService) Get(ctx context.Context, clusterId, nodePoolId string) (*NodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s", clusterId, nodePoolId)
	req, err := nps.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	var NodePoolRoot NodePoolRoot

	if _, err = nps.client.Do(ctx, req, &NodePoolRoot); err != nil {
		return nil, err
	}

	return NodePoolRoot.NodePool, nil
}

func (nps *NodePoolsService) List(ctx context.Context, options *ListOptions, clusterId string) ([]NodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools", clusterId)

	var NodePoolsRoot NodePoolsRoot

	if err := nps.client.list(ctx, path, options, &NodePoolsRoot); err != nil {
		return nil, err
	}

	return NodePoolsRoot.NodePools, nil
}

func (nps *NodePoolsService) Create(ctx context.Context, clusterId string, request *CreateNodePoolRequest) (*NodePool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools", clusterId)
	req, err := nps.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var NodePoolRoot NodePoolRoot
	if _, err := nps.client.Do(ctx, req, &NodePoolRoot); err != nil {
		return nil, err
	}

	return NodePoolRoot.NodePool, nil
}

func (nps *NodePoolsService) Update(ctx context.Context, clusterId, nodePoolId string, request *NodePoolUpdateRequest) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s", clusterId, nodePoolId)
	req, err := nps.client.newRequest(http.MethodPut, path, request)
	if err != nil {
		return err
	}

	if _, err := nps.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

func (nps *NodePoolsService) Delete(ctx context.Context, clusterId string, nodePoolId string, replace bool) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/node_pools/%s?replace=%v", clusterId, nodePoolId, replace)
	req, err := nps.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if _, err := nps.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
