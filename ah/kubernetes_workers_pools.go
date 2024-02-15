package ah

import (
	"context"
	"fmt"
	"net/http"
)

// Labels object
type Labels map[string]string

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

// KubernetesWorkerPool object
type KubernetesWorkerPool struct {
	Labels            Labels             `json:"labels,omitempty"`
	ID                string             `json:"id,omitempty"`
	Name              string             `json:"name,omitempty"`
	Type              string             `json:"type"`
	CreatedAt         string             `json:"created_at,omitempty"`
	Workers           []KubernetesWorker `json:"workers,omitempty"`
	PrivateProperties PrivateProperties  `json:"private_properties,omitempty"`
	PublicProperties  PublicProperties   `json:"public_properties,omitempty"`
	Count             int                `json:"count,omitempty"`
	AutoScale         bool               `json:"autoscale,omitempty"`
	MinCount          int                `json:"min_count,omitempty"`
	MaxCount          int                `json:"max_count,omitempty"`
}

type KubernetesWorkerPoolRoot struct {
	KubernetesWorkerPool *KubernetesWorkerPool `json:"worker_pool,omitempty"`
}

type KubernetesWorkerPoolsRoot struct {
	KubernetesWorkerPools []KubernetesWorkerPool `json:"worker_pools,omitempty"`
}

// CreateKubernetesWorkerPoolRequest represents a request to create a worker pool.
type CreateKubernetesWorkerPoolRequest struct {
	PrivateProperties *PrivateProperties `json:"private_properties,omitempty"`
	PublicProperties  *PublicProperties  `json:"public_properties,omitempty"`
	Labels            *Labels            `json:"labels,omitempty"`
	Type              string             `json:"type"`
	Count             int                `json:"count,omitempty"`
	MinCount          int                `json:"min_count,omitempty"`
	MaxCount          int                `json:"max_count,omitempty"`
	AutoScale         bool               `json:"autoscale,omitempty"`
}

// UpdateKubernetesWorkerPoolRequest represents a request to update a worker pool
type UpdateKubernetesWorkerPoolRequest struct {
	Labels    *Labels `json:"labels,omitempty"`
	Count     int     `json:"count,omitempty"`
	AutoScale bool    `json:"autoscale,omitempty"`
	MinCount  int     `json:"min_count,omitempty"`
	MaxCount  int     `json:"max_count,omitempty"`
}

// GetWorkerPool returns worker pool
func (kcs *KubernetesClustersService) GetWorkerPool(ctx context.Context, clusterId, workerPoolId string) (*KubernetesWorkerPool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/worker_pools/%s", clusterId, workerPoolId)
	req, err := kcs.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	var workerPoolRoot KubernetesWorkerPoolRoot

	if _, err = kcs.client.Do(ctx, req, &workerPoolRoot); err != nil {
		return nil, err
	}

	return workerPoolRoot.KubernetesWorkerPool, nil
}

// ListWorkerPools returns list of worker pools
func (kcs *KubernetesClustersService) ListWorkerPools(ctx context.Context, options *ListOptions, clusterId string) ([]KubernetesWorkerPool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/worker_pools", clusterId)

	var WorkerPoolsRoot KubernetesWorkerPoolsRoot

	if err := kcs.client.list(ctx, path, options, &WorkerPoolsRoot); err != nil {
		return nil, err
	}

	return WorkerPoolsRoot.KubernetesWorkerPools, nil
}

// CreateWorkerPool creates worker pool
func (kcs *KubernetesClustersService) CreateWorkerPool(ctx context.Context, clusterId string, request *CreateKubernetesWorkerPoolRequest) (*KubernetesWorkerPool, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/worker_pools", clusterId)
	req, err := kcs.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var workerPoolRoot KubernetesWorkerPoolRoot
	if _, err := kcs.client.Do(ctx, req, &workerPoolRoot); err != nil {
		return nil, err
	}

	return workerPoolRoot.KubernetesWorkerPool, nil
}

// UpdateWorkerPool updates worker pool
func (kcs *KubernetesClustersService) UpdateWorkerPool(ctx context.Context, clusterId, workerPoolId string, request *UpdateKubernetesWorkerPoolRequest) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/worker_pools/%s", clusterId, workerPoolId)
	req, err := kcs.client.newRequest(http.MethodPatch, path, request)
	if err != nil {
		return err
	}

	if _, err := kcs.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}

// DeleteWorkerPool deletes worker pool
func (kcs *KubernetesClustersService) DeleteWorkerPool(ctx context.Context, clusterId string, workerPoolId string, replace bool) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s/worker_pools/%s?replace=%v", clusterId, workerPoolId, replace)
	req, err := kcs.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if _, err := kcs.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
