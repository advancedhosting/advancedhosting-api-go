/*
 Copyright 2023 Advanced Hosting

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

// ConfigFile object
type ConfigFile any

// KubernetesClustersToken object
type KubernetesClustersToken struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

// KubernetesClustersTokenRoot object
type kubernetesClustersTokenRoot struct {
	KubernetesTokens []KubernetesClustersToken `json:"tokens,omitempty"`
}

// KubernetesClustersConfig object
type KubernetesClustersConfig struct {
	Config ConfigFile `json:"config"`
}

// KubernetesClustersTokenCreateRequest object
type KubernetesClustersTokenCreateRequest struct {
	Name string `json:"name"`
}

// KubernetesClustersTokensAPI is an interface for kubernetes tokens.
type KubernetesClustersTokensAPI interface {
	List(context.Context, string, *ListOptions) ([]KubernetesClustersToken, error)
	Get(context.Context, string, string) (ConfigFile, error)
	GetDefault(context.Context, string) (ConfigFile, error)
	Create(context.Context, string, *KubernetesClustersTokenCreateRequest) error
	Delete(context.Context, string, string) error
}

// KubernetesClustersTokensService implements KubernetesTokensAPI interface.
type KubernetesClustersTokensService struct {
	client *APIClient
}

// Get returns a kubernetes clusters token by ID
func (kts *KubernetesClustersTokensService) Get(ctx context.Context, clusterID string, tokenID string) (ConfigFile, error) {
	path := fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs/%s", clusterID, tokenID)

	req, err := kts.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	var kubernetesClustersConfig KubernetesClustersConfig
	_, err = kts.client.Do(ctx, req, &kubernetesClustersConfig)
	if err != nil {
		return "", err
	}

	return kubernetesClustersConfig.Config, nil
}

// GetDefault returns a default kubernetes clusters token
func (kts *KubernetesClustersTokensService) GetDefault(ctx context.Context, clusterID string) (ConfigFile, error) {
	path := fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs/default", clusterID)

	req, err := kts.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")

	var kubernetesClustersConfig KubernetesClustersConfig
	_, err = kts.client.Do(ctx, req, &kubernetesClustersConfig)
	if err != nil {
		return "", err
	}

	return kubernetesClustersConfig.Config, nil
}

// List returns a list of kubernetes clusters tokens
func (kts *KubernetesClustersTokensService) List(ctx context.Context, clusterID string, options *ListOptions) ([]KubernetesClustersToken, error) {
	path := fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs", clusterID)

	var kubernetesClustersTokens kubernetesClustersTokenRoot
	if err := kts.client.list(ctx, path, options, &kubernetesClustersTokens); err != nil {
		return nil, err
	}

	return kubernetesClustersTokens.KubernetesTokens, nil
}

// Create creates a kubernetes clusters token
func (kts *KubernetesClustersTokensService) Create(ctx context.Context, clusterID string, req *KubernetesClustersTokenCreateRequest) error {
	path := fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs", clusterID)
	request, err := kts.client.newRequest(http.MethodPost, path, req)
	if err != nil {
		return err
	}

	_, err = kts.client.Do(ctx, request, nil)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a kubernetes clusters token
func (kts *KubernetesClustersTokensService) Delete(ctx context.Context, clusterID string, tokenID string) error {
	path := fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs/%s", clusterID, tokenID)
	req, err := kts.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = kts.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
