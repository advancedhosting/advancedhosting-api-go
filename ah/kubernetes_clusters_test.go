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

const clusterResponse = `{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"name": "string",
	"datacenter_id": "5839cebe-c7a5-4a27-8253-7bd619ca430d",
	"state": "defined",
	"nodes_count": 0,
	"created_at": "2019-08-24T14:15:22Z",
	"plan_id": 0,
	"number": "string"
}`

const configResponse = "Cluster config"

var (
	clusterGetResponse    = fmt.Sprintf(`{"cluster": %s}`, clusterResponse)
	clusterListResponse   = fmt.Sprintf(`{"clusters": [%s]}`, clusterResponse)
	clusterConfigResponse = fmt.Sprintf(`{"config": "%s"}`, configResponse)
)

func TestCluster_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: clusterGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v1/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult clusterRoot
	if err := json.Unmarshal([]byte(clusterGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	cluster, err := api.Clusters.Get(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cluster == nil || cluster.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", cluster)
	}

	if !reflect.DeepEqual(expectedResult.Cluster, cluster) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, cluster)
	}
}

func TestClusters_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: clusterListResponse}
	api, _ := newFakeAPIClient("/api/v1/kubernetes/clusters", fakeResponse)

	ctx := context.Background()

	clusters, err := api.Clusters.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult clustersRoot
	if err := json.Unmarshal([]byte(clusterListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.Clusters, clusters) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, clusters)
	}
}

func TestCluster_Create(t *testing.T) {
	request := &ClusterCreateRequest{
		Name:         "New Kubernetes Cluster",
		DatacenterID: "5839cebe-c7a5-4a27-8253-7bd619ca430d",
		PlanId:       0,
		NodesCount:   1,
		PrivateCloud: true,
	}

	fakeResponse := &fakeServerResponse{
		responseBody: clusterGetResponse,
		statusCode:   201,
	}
	api, _ := newFakeAPIClient("/api/v1/kubernetes/clusters", fakeResponse)

	ctx := context.Background()

	cluster, err := api.Clusters.Create(ctx, request)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if cluster == nil {
		t.Errorf("Empty response")
	}
}

func TestCluster_Update(t *testing.T) {
	var clusterId = "5839cebe-c7a5-4a27-8253-7bd619ca430d"
	request := &ClusterUpdateRequest{
		Name: "Updated Kubernetes Cluster Name",
	}
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 204}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v1/kubernetes/clusters/%s", clusterId),
		fakeResponse,
	)

	ctx := context.Background()

	err := api.Clusters.Update(ctx, clusterId, request)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestCluster_Delete(t *testing.T) {
	var clusterId = "5839cebe-c7a5-4a27-8253-7bd619ca430d"

	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v1/kubernetes/clusters/%s", clusterId),
		fakeResponse,
	)

	ctx := context.Background()

	err := api.Clusters.Delete(ctx, clusterId)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestCluster_GetConfig(t *testing.T) {
	var clusterId = "5839cebe-c7a5-4a27-8253-7bd619ca430d"

	fakeResponse := &fakeServerResponse{responseBody: clusterConfigResponse, statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v1/kubernetes/clusters/%s/kubeconfig", clusterId),
		fakeResponse,
	)

	ctx := context.Background()

	config, err := api.Clusters.GetConfig(ctx, clusterId)
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
