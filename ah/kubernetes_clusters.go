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

// KubernetesCluster object
type KubernetesCluster struct {
	ID                 string               `json:"id,omitempty"`
	Name               string               `json:"name,omitempty"`
	TokenID            string               `json:"token_id,omitempty"`
	DatacenterID       string               `json:"datacenter_id,omitempty"`
	DatacenterSlug     string               `json:"datacenter_slug,omitempty"`
	State              string               `json:"state,omitempty"`
	Number             string               `json:"number"`
	CreatedAt          string               `json:"created_at"`
	AccountID          string               `json:"account_id"`
	PrivateNetworkID   string               `json:"private_network_id"`
	PrivateNetworkName string               `json:"private_network_name,omitempty"`
	K8sVersion         string               `json:"k8s_version"`
	NodePools          []KubernetesNodePool `json:"node_pools,omitempty"`
}

// KubernetesClusterConfig object
type KubernetesClusterConfig struct {
	Config string `json:"config"`
}

// KubernetesClustersAPI is an interface for cluster API.
type KubernetesClustersAPI interface {
	Get(context.Context, string) (*KubernetesCluster, error)
	List(context.Context, *ListOptions) ([]KubernetesCluster, error)
	Create(context.Context, *KubernetesClusterCreateRequest) (*KubernetesCluster, error)
	Update(context.Context, string, *KubernetesClusterUpdateRequest) error
	GetConfig(context.Context, string) (string, error)
	Delete(context.Context, string) error
	GetKubernetesClustersVersions(context.Context) ([]string, error)
	GetNodePool(context.Context, string, string) (*KubernetesNodePool, error)
	ListNodePools(context.Context, *ListOptions, string) ([]KubernetesNodePool, error)
	CreateNodePool(context.Context, string, *CreateKubernetesNodePoolRequest) (*KubernetesNodePool, error)
	UpdateNodePool(context.Context, string, string, *UpdateKubernetesNodePoolRequest) error
	DeleteNodePool(context.Context, string, string, bool) error
}

// KubernetesClustersService implements ClustersAPI interface.
type KubernetesClustersService struct {
	client *APIClient
}

type KubernetesClusterRoot struct {
	KubernetesCluster *KubernetesCluster `json:"cluster,omitempty"`
}

type KubernetesClustersRoot struct {
	KubernetesClusters []KubernetesCluster `json:"clusters,omitempty"`
}

// KubernetesClusterCreateRequest represents a request to create a cluster.
type KubernetesClusterCreateRequest struct {
	Name         string                            `json:"name"`
	DatacenterID string                            `json:"datacenter_id,omitempty"`
	K8sVersion   string                            `json:"k8s_version"`
	NodePools    []CreateKubernetesNodePoolRequest `json:"node_pools"`
}

// KubernetesClusterUpdateRequest represents a request to update a cluster
type KubernetesClusterUpdateRequest struct {
	Name string `json:"name,omitempty"`
}

// Create kubernetes cluster
func (kcs *KubernetesClustersService) Create(ctx context.Context, createRequest *KubernetesClusterCreateRequest) (*KubernetesCluster, error) {
	req, err := kcs.client.newRequest(http.MethodPost, "api/v2/kubernetes/clusters", createRequest)
	if err != nil {
		return nil, err
	}

	var kubernetesClusterRoot KubernetesClusterRoot
	if _, err := kcs.client.Do(ctx, req, &kubernetesClusterRoot); err != nil {
		return nil, err
	}

	return kubernetesClusterRoot.KubernetesCluster, nil
}

// Get kubernetes cluster
func (kcs *KubernetesClustersService) Get(ctx context.Context, clusterID string) (*KubernetesCluster, error) {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s", clusterID)

	req, err := kcs.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var kubernetesClusterRoot KubernetesClusterRoot
	_, err = kcs.client.Do(ctx, req, &kubernetesClusterRoot)
	if err != nil {
		return nil, err
	}

	return kubernetesClusterRoot.KubernetesCluster, nil
}

// List returns list of kubernetes clusters
func (kcs *KubernetesClustersService) List(ctx context.Context, options *ListOptions) ([]KubernetesCluster, error) {
	path := "/api/v2/kubernetes/clusters"

	var kubernetesClustersRoot KubernetesClustersRoot
	if err := kcs.client.list(ctx, path, options, &kubernetesClustersRoot); err != nil {
		return nil, err
	}

	return kubernetesClustersRoot.KubernetesClusters, nil
}

// Update kubernetes cluster. Returns error
func (kcs *KubernetesClustersService) Update(ctx context.Context, clusterId string, request *KubernetesClusterUpdateRequest) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s", clusterId)

	req, err := kcs.client.newRequest(http.MethodPatch, path, request)
	if err != nil {
		return err
	}

	_, err = kcs.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete kubernetes cluster. Returns error
func (kcs *KubernetesClustersService) Delete(ctx context.Context, clusterId string) error {
	path := fmt.Sprintf("api/v2/kubernetes/clusters/%s", clusterId)

	req, err := kcs.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = kcs.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetKubernetesClustersVersions returns kubernetes version
func (kcs *KubernetesClustersService) GetKubernetesClustersVersions(ctx context.Context) ([]string, error) {
	path := "/api/v2/kubernetes/clusters/versions"

	req, err := kcs.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var versions []string

	_, err = kcs.client.Do(ctx, req, &versions)
	if err != nil {
		return nil, err
	}

	return versions, nil
}

// GetConfig returns kubernetes cluster config
func (kcs KubernetesClustersService) GetConfig(ctx context.Context, clusterId string) (string, error) {
	path := fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfig", clusterId)

	req, err := kcs.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	var configRoot KubernetesClusterConfig
	_, err = kcs.client.Do(ctx, req, &configRoot)
	if err != nil {
		return "", err
	}

	return configRoot.Config, nil
}
