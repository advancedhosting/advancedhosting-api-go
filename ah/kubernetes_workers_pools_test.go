package ah

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const WorkerPoolPublicResponse = `{
	"id": "e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
	"name": "KNP100000",
	"type": "public",
    "count": 1,
	"labels": {"labels.websa.com/test": "test"},
    "public_properties": {
		"plan_id": 111111111
	},
    "workers": [{
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

const WorkerPoolPrivateResponse = `{
	"id": "7a41f73b-009d-4e52-9f20-edf4b8c5b072",
	"name": "KNP100000",
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
	"workers": [{
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
	workerPoolPublicGetResponse  = fmt.Sprintf(`{"worker_pool": %s}`, WorkerPoolPublicResponse)
	workerPoolPrivateGetResponse = fmt.Sprintf(`{"worker_pool": %s}`, WorkerPoolPrivateResponse)
	workerPoolListResponse       = fmt.Sprintf(`{"worker_pools": [%s, %s]}`, WorkerPoolPublicResponse, WorkerPoolPrivateResponse)
)

func TestWorkerPoolPublicGet(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: workerPoolPublicGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesWorkerPoolRoot
	if err := json.Unmarshal([]byte(workerPoolPublicGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	workerPool, err := api.KubernetesClusters.GetWorkerPool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if workerPool == nil || workerPool.ID != "e312aa01-d123-4a90-9c6d-f7641d2e4cc7" {
		t.Errorf("Invalid response: %v", workerPool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesWorkerPool, workerPool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, workerPool)
	}
}

func TestWorkerPooPrivateGet(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: workerPoolPrivateGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools/7a41f73b-009d-4e52-9f20-edf4b8c5b072",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesWorkerPoolRoot
	if err := json.Unmarshal([]byte(workerPoolPrivateGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	workerPool, err := api.KubernetesClusters.GetWorkerPool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "7a41f73b-009d-4e52-9f20-edf4b8c5b072")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if workerPool == nil || workerPool.ID != "7a41f73b-009d-4e52-9f20-edf4b8c5b072" {
		t.Errorf("Invalid response: %v", workerPool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesWorkerPool, workerPool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, workerPool)
	}
}

func TestWorkerPoolList(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: workerPoolListResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesWorkerPoolsRoot
	if err := json.Unmarshal([]byte(workerPoolListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	workerPools, err := api.KubernetesClusters.ListWorkerPools(ctx, nil, "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if workerPools == nil || workerPools[0].ID != "e312aa01-d123-4a90-9c6d-f7641d2e4cc7" {
		t.Errorf("Invalid response: %v", workerPools)
	}

	if workerPools == nil || workerPools[1].ID != "7a41f73b-009d-4e52-9f20-edf4b8c5b072" {
		t.Errorf("Invalid response: %v", workerPools)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesWorkerPools, workerPools) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, workerPools)
	}
}

func TestWorkerPoolPublicCreate(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: workerPoolPublicGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesWorkerPoolRoot
	if err := json.Unmarshal([]byte(workerPoolPublicGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	publicProperties := &PublicProperties{PlanID: 111111111}

	request := &CreateKubernetesWorkerPoolRequest{
		Type:             "public",
		Count:            1,
		PublicProperties: publicProperties,
	}

	workerPool, err := api.KubernetesClusters.CreateWorkerPool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if workerPool == nil || workerPool.ID != "e312aa01-d123-4a90-9c6d-f7641d2e4cc7" {
		t.Errorf("Invalid response: %v", workerPool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesWorkerPool, workerPool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, workerPool)
	}
}

func TestWorkerPoolPrivateCreate(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: workerPoolPrivateGetResponse}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools",
		fakeResponse,
	)

	ctx := context.Background()

	var expectedResult KubernetesWorkerPoolRoot
	if err := json.Unmarshal([]byte(workerPoolPrivateGetResponse), &expectedResult); err != nil {
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

	request := &CreateKubernetesWorkerPoolRequest{
		Type:              "private",
		Count:             1,
		PrivateProperties: privateProperties,
	}

	workerPool, err := api.KubernetesClusters.CreateWorkerPool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if workerPool == nil || workerPool.ID != "7a41f73b-009d-4e52-9f20-edf4b8c5b072" {
		t.Errorf("Invalid response: %v", workerPool)
	}

	if !reflect.DeepEqual(expectedResult.KubernetesWorkerPool, workerPool) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, workerPool)
	}
}

func TestWorkerPoolUpdate(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
		fakeResponse,
	)

	ctx := context.Background()

	request := &UpdateKubernetesWorkerPoolRequest{
		Count:     1,
		AutoScale: false,
	}

	err := api.KubernetesClusters.UpdateWorkerPool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7", request)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestWorkerPoolDelete(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7",
		fakeResponse,
	)

	ctx := context.Background()

	err := api.KubernetesClusters.DeleteWorkerPool(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7", false)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
