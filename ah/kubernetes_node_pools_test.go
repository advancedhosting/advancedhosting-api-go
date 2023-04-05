package ah

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const nodePoolPublicResponse = `{
	"id": "e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
	"name": "test",
	"type": "public",
    "count": 1,
	"labels": {"labels.websa.com/test": "test"},
    "public_properties": {
		"plan_id": 111111111
	},
    "nodes": [{
		"id": "339e3dd3-9734-40ec-8e5c-aa7c4c3be319",
        "name": "kub1000000-620d6",
        "labels": {},
        "private_network_id": "0a1004fb-6930-49f5-9a2e-1e713b50d850",
        "external_ip_id": "36ec9173-b77a-45b1-81c9-dc3c61c47630",
        "created_at": "2023-02-27T16:38:15.000+00:00",
        "state": "active",
        "type": "public",
        "cloud_server_id": "8dbcffb9-d3be-483c-938c-6f9b2c9a7bc4"
	}]
}`

const nodePoolPrivateResponse = `{
	"id": "7a41f73b-009d-4e52-9f20-edf4b8c5b072",
	"name": "test",
	"type": "public",
    "count": 0,
	"labels": {"labels.websa.com/test": "test"},
	"private_properties": {
	    "vcpu": 2,
	    "ram": 512,
	    "disk": 10,
	    "network_id": "0a1004fb-6930-49f5-9a2e-1e713b50d850",
	    "cluster_id": "ca2a8d11-8426-4493-8ee3-82b700c6092b",
	    "cluster_node_id": "33bb3e37-7e2d-4a55-936d-fae0223d5a00"
	},
	"nodes": [{
		"id": "339e3dd3-9734-40ec-8e5c-aa7c4c3be319",
        "name": "kub1000000-620d6",
        "labels": {},
        "private_network_id": "0a1004fb-6930-49f5-9a2e-1e713b50d850",
        "external_ip_id": "36ec9173-b77a-45b1-81c9-dc3c61c47630",
        "created_at": "2023-02-27T16:38:15.000+00:00",
        "state": "active",
        "type": "public",
        "cloud_server_id": "8dbcffb9-d3be-483c-938c-6f9b2c9a7bc4"
	}]
}`

var (
	nodePoolPublicGetResponse  = fmt.Sprintf(`{"node_pool": %s}`, nodePoolPublicResponse)
	nodePoolPrivateGetResponse = fmt.Sprintf(`{"node_pool": %s}`, nodePoolPrivateResponse)
	nodePoolListResponse       = fmt.Sprintf(`{"node_pools": [%s, %s]}`, nodePoolPublicResponse, nodePoolPrivateResponse)
)

func TestNodePoolPublicGet(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: nodePoolPublicGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesNodePoolRoot
	if err := json.Unmarshal([]byte(nodePoolPublicGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	nodePool, err := api.Clusters.GetKubernetesNodePool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if nodePool == nil || nodePool.ID != "e312aa01-d123-4a90-9c6d-f7641d2e4cc7" {
		t.Errorf("Invalid response: %v", nodePool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesNodePool, nodePool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, nodePool)
	}
}

func TestNodePooPrivatelGet(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: nodePoolPrivateGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools/7a41f73b-009d-4e52-9f20-edf4b8c5b072",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesNodePoolRoot
	if err := json.Unmarshal([]byte(nodePoolPrivateGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	nodePool, err := api.Clusters.GetKubernetesNodePool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "7a41f73b-009d-4e52-9f20-edf4b8c5b072")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if nodePool == nil || nodePool.ID != "7a41f73b-009d-4e52-9f20-edf4b8c5b072" {
		t.Errorf("Invalid response: %v", nodePool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesNodePool, nodePool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, nodePool)
	}
}

func TestNodePoolList(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: nodePoolListResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesNodePoolsRoot
	if err := json.Unmarshal([]byte(nodePoolListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	nodePools, err := api.Clusters.ListKubernetesNodePools(ctx, nil, "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if nodePools == nil || nodePools[0].ID != "e312aa01-d123-4a90-9c6d-f7641d2e4cc7" {
		t.Errorf("Invalid response: %v", nodePools)
	}

	if nodePools == nil || nodePools[1].ID != "7a41f73b-009d-4e52-9f20-edf4b8c5b072" {
		t.Errorf("Invalid response: %v", nodePools)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesNodePools, nodePools) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, nodePools)
	}
}

func TestNodePoolPublicCreate(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: nodePoolPublicGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesNodePoolRoot
	if err := json.Unmarshal([]byte(nodePoolPublicGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	publicProperties := &PublicProperties{PlanID: 111111111}

	request := &CreateKubernetesNodePoolRequest{
		Name:             "test",
		Type:             "public",
		Count:            1,
		Labels:           map[string]string{},
		PublicProperties: *publicProperties,
	}

	nodePool, err := api.Clusters.CreateKubernetesNodePool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if nodePool == nil || nodePool.ID != "e312aa01-d123-4a90-9c6d-f7641d2e4cc7" {
		t.Errorf("Invalid response: %v", nodePool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesNodePool, nodePool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, nodePool)
	}
}

func TestNodePoolPrivateCreate(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: nodePoolPrivateGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesNodePoolRoot
	if err := json.Unmarshal([]byte(nodePoolPrivateGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	privateProperties := &PrivateProperties{
		Vcpu:          4,
		Ram:           4096,
		Disk:          40,
		NetworkID:     "0a1004fb-6930-49f5-9a2e-1e713b50d850",
		ClusterID:     "ca2a8d11-8426-4493-8ee3-82b700c6092b",
		ClusterNodeID: "33bb3e37-7e2d-4a55-936d-fae0223d5a00",
	}

	request := &CreateKubernetesNodePoolRequest{
		Name:              "test",
		Type:              "private",
		Count:             1,
		Labels:            map[string]string{},
		PrivateProperties: *privateProperties,
	}

	nodePool, err := api.Clusters.CreateKubernetesNodePool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if nodePool == nil || nodePool.ID != "7a41f73b-009d-4e52-9f20-edf4b8c5b072" {
		t.Errorf("Invalid response: %v", nodePool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesNodePool, nodePool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, nodePool)
	}
}

func TestNodePoolUpdate(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
		fakeResponse,
	)

	ctx := context.Background()

	request := &UpdateKubernetesNodePoolRequest{
		Name:      "test",
		Count:     1,
		Labels:    map[string]string{},
		AutoScale: false,
	}

	err := api.Clusters.UpdateKubernetesNodePool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7", request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestNodePoolDelete(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/node_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
		fakeResponse,
	)

	ctx := context.Background()

	err := api.Clusters.DeleteKubernetesNodePool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7", false)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
