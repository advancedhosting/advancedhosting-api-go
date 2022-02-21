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
}

// ClustersAPI is an interface for load balancers.
type ClustersAPI interface {
	Get(context.Context, string) (*Cluster, error)
}

// ClustersService implements ClustersAPI interface.
type ClustersService struct {
	client *APIClient
}

type clusterRoot struct {
	Cluster *Cluster `json:"cluster,omitempty"`
}

// Get kubernetes cluster
func (kc *ClustersService) Get(ctx context.Context, clusterID string) (*Cluster, error) {
	path := fmt.Sprintf("api/v1/kubernetes/clusters/%s", clusterID)

	req, err := kc.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var csRoot clusterRoot
	_, err = kc.client.Do(ctx, req, &csRoot)

	if err != nil {
		return nil, err
	}

	return csRoot.Cluster, nil
}
