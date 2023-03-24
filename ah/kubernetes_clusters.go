/*
Copyright 2022 Advanced Hosting

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ah

import (
	"context"
	"fmt"
	"net/http"
)

// Cluster object
type Cluster struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	DatacenterID string `json:"datacenter_id,omitempty"`
	State        string `json:"state,omitempty"`
	Number       string `json:"number"`
	CreatedAt    string `json:"created_at"`
	Count        int    `json:"count"`
	PlanID       int    `json:"plan_id"`
	Vcpu         int    `json:"vcpu"`
	Ram          int    `json:"ram"`
	Disk         int    `json:"disk"`
}

// clusterConfig object
type clusterConfig struct {
	Config string `json:"config"`
}

// ClustersAPI is an interface for cluster API.
type ClustersAPI interface {
	Get(context.Context, string) (*Cluster, error)
	List(context.Context, *ListOptions) ([]Cluster, error)
	Create(context.Context, *ClusterCreateRequest) (*Cluster, error)
	Update(context.Context, string, *ClusterUpdateRequest) error
	GetConfig(context.Context, string) (string, error)
	Delete(context.Context, string) error
}

// ClustersService implements ClustersAPI interface.
type ClustersService struct {
	client *APIClient
}

type clusterRoot struct {
	Cluster *Cluster `json:"cluster,omitempty"`
}

type clustersRoot struct {
	Clusters []Cluster `json:"clusters,omitempty"`
}

// ClusterCreateRequest represents a request to create a cluster.
type ClusterCreateRequest struct {
	Name         string `json:"name"`
	DatacenterID string `json:"datacenter_id,omitempty"`
	PlanId       int    `json:"plan_id,omitempty"`
	Vcpu         int    `json:"vcpu,omitempty"`
	Ram          int    `json:"ram,omitempty"`
	Disk         int    `json:"disk,omitempty"`
	Count        int    `json:"count,omitempty"`
	PrivateCloud bool   `json:"private_cloud"`
}

// ClusterUpdateRequest represents a request to update a cluster
type ClusterUpdateRequest struct {
	Name string `json:"name,omitempty"`
}

// Create kubernetes cluster
func (kc *ClustersService) Create(ctx context.Context, createRequest *ClusterCreateRequest) (*Cluster, error) {
	req, err := kc.client.newRequest(http.MethodPost, "api/v1/kubernetes/clusters", createRequest)
	if err != nil {
		return nil, err
	}

	var clusterRoot clusterRoot
	if _, err := kc.client.Do(ctx, req, &clusterRoot); err != nil {
		return nil, err
	}

	return clusterRoot.Cluster, nil
}

// Get kubernetes cluster
func (kc *ClustersService) Get(ctx context.Context, clusterID string) (*Cluster, error) {
	path := fmt.Sprintf("api/v1/kubernetes/clusters/%s", clusterID)

	req, err := kc.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var clusterRoot clusterRoot
	_, err = kc.client.Do(ctx, req, &clusterRoot)
	if err != nil {
		return nil, err
	}

	return clusterRoot.Cluster, nil
}

// List returns list of kubernetes clusters
func (kc *ClustersService) List(ctx context.Context, options *ListOptions) ([]Cluster, error) {
	path := "/api/v1/kubernetes/clusters"

	var clustersRoot clustersRoot
	if err := kc.client.list(ctx, path, options, &clustersRoot); err != nil {
		return nil, err
	}

	return clustersRoot.Clusters, nil
}

// Update kubernetes cluster. Returns error
func (kc *ClustersService) Update(ctx context.Context, clusterId string, request *ClusterUpdateRequest) error {
	path := fmt.Sprintf("api/v1/kubernetes/clusters/%s", clusterId)

	req, err := kc.client.newRequest(http.MethodPatch, path, request)
	if err != nil {
		return err
	}

	_, err = kc.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete kubernetes cluster. Returns error
func (kc *ClustersService) Delete(ctx context.Context, clusterId string) error {
	path := fmt.Sprintf("api/v1/kubernetes/clusters/%s", clusterId)

	req, err := kc.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = kc.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig returns kubernetes cluster config
func (kc ClustersService) GetConfig(ctx context.Context, clusterId string) (string, error) {
	path := fmt.Sprintf("/api/v1/kubernetes/clusters/%s/kubeconfig", clusterId)

	req, err := kc.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	var configRoot clusterConfig
	_, err = kc.client.Do(ctx, req, &configRoot)
	if err != nil {
		return "", err
	}

	return configRoot.Config, nil
}
