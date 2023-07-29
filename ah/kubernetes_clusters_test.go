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
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var (
	clusterResponse = fmt.Sprintf(`{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"name": "New Kubernetes Cluster",
	"token_id": "ad85a5d3-99ad-4f05-a5ac-24eb27b9cd8c",
	"datacenter_id": "5839cebe-c7a5-4a27-8253-7bd619ca430d",
	"datacenter_slug": "ams1",
	"state": "active",
	"created_at": "2019-08-24T14:15:22Z",
	"number": "KUB1000000",
	"account_id": "085592e9-dd2f-490d-b972-7c70142f91b0",
	"private_network_id": "c2a96164-d7f4-45db-b61a-e757de64483e",
	"private_network_name": "NET10000000",
	"k8s_version": "1.19.3",
	"worker_pools": [%s,%s]
}`, WorkerPoolPublicResponse, WorkerPoolPrivateResponse)
	kubernetesVersions = `["v1.19.3", "v1.18.10", "v1.17.13"]`
)

const configResponse = "Cluster config"

var (
	clusterGetResponse    = fmt.Sprintf(`{"cluster": %s}`, clusterResponse)
	clusterListResponse   = fmt.Sprintf(`{"clusters": [%s]}`, clusterResponse)
	clusterConfigResponse = fmt.Sprintf(`{"config": "%s"}`, configResponse)
)

func TestCluster_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: clusterGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesClusterRoot
	if err := json.Unmarshal([]byte(clusterGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	cluster, err := api.KubernetesClusters.Get(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cluster == nil || cluster.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", cluster)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesCluster, cluster) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, cluster)
	}
}

func TestClusters_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: clusterListResponse}
	api, _ := newFakeAPIClient("/api/v2/kubernetes/clusters", fakeResponse)

	ctx := context.Background()

	clusters, err := api.KubernetesClusters.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult KubernetesClustersRoot
	if err := json.Unmarshal([]byte(clusterListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesClusters, clusters) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, clusters)
	}
}

func TestCluster_Create(t *testing.T) {
	request := &KubernetesClusterCreateRequest{
		Name:         "New Kubernetes Cluster",
		DatacenterID: "5839cebe-c7a5-4a27-8253-7bd619ca430d",
		K8sVersion:   "1.19.3",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: clusterGetResponse,
		statusCode:   201,
	}
	api, _ := newFakeAPIClient("/api/v2/kubernetes/clusters", fakeResponse)

	ctx := context.Background()

	cluster, err := api.KubernetesClusters.Create(ctx, request)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if cluster == nil {
		t.Errorf("Empty response")
	}
}

func TestCluster_Update(t *testing.T) {
	var clusterId = "5839cebe-c7a5-4a27-8253-7bd619ca430d"
	request := &KubernetesClusterUpdateRequest{
		Name: "Updated Kubernetes Cluster Name",
	}
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 204}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s", clusterId),
		fakeResponse,
	)

	ctx := context.Background()

	err := api.KubernetesClusters.Update(ctx, clusterId, request)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestCluster_Delete(t *testing.T) {
	var clusterId = "5839cebe-c7a5-4a27-8253-7bd619ca430d"

	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s", clusterId),
		fakeResponse,
	)

	ctx := context.Background()

	err := api.KubernetesClusters.Delete(ctx, clusterId)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestCluster_GetConfig(t *testing.T) {
	var clusterId = "5839cebe-c7a5-4a27-8253-7bd619ca430d"

	fakeResponse := &fakeServerResponse{responseBody: clusterConfigResponse, statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfig", clusterId),
		fakeResponse,
	)

	ctx := context.Background()

	config, err := api.KubernetesClusters.GetConfig(ctx, clusterId)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if config == "" {
		t.Errorf("Empty response")
	}
	if config != configResponse {
		t.Errorf("Unexpected response")
	}
}

func TestKubernetesClustersVersions(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: kubernetesVersions, statusCode: 200}

	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/versions",
		fakeResponse,
	)

	ctx := context.Background()

	versions, err := api.KubernetesClusters.GetKubernetesClustersVersions(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if versions == nil {
		t.Errorf("Empty response")
	}

	found := false
	for _, version := range versions {
		if version == "v1.19.3" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Version v1.19.3 not found in the response")
	}
}
