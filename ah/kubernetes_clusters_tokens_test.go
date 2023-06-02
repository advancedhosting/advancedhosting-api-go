package ah

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

var (
	clusterID = "5839cebe-c7a5-4a27-8253-7bd619ca430d"
	tokenID   = "1eaa2266-29da-4b23-b986-e3cf3ff0611f"
)

const KubernetesClustersConfigResponse = `{
    "apiVersion": "v1",
    "kind": "Config",
    "preferences": {},
    "clusters": [
      {
        "name": "kub1000000",
        "cluster": {
          "certificate-authority-data": "TEST",
          "server": "https://kub1000000.com"
        }
      }
    ],
    "users": [
      {
        "name": "admin",
        "user": {
          "client-certificate-data": "TEST",
          "client-key-data": "TEST"
        }
      }
    ],
    "current-context": "admin@kub1000000",
    "contexts": [
      {
        "name": "admin@kub1000000",
        "context": {
          "cluster": "kub1000000",
          "user": "admin"
        }
      }
    ]
}`

const KubernetesClustersTokenResponse = `{
	"id": "1eaa2266-29da-4b23-b986-e3cf3ff0611f",
	"name": "Test Kubernetes Token",
	"created_at": "2020-04-15T10:51:15.765Z"
}`

var (
	kubernetesClustersConfigGetResponse = fmt.Sprintf(`{"config": %s}`, KubernetesClustersConfigResponse)
	kubernetesClustersTokenListResponse = fmt.Sprintf(`{"tokens": [%s]}`, KubernetesClustersTokenResponse)
)

func TestKubernetesClustersTokensService_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: kubernetesClustersConfigGetResponse, statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs/%s", clusterID, tokenID),
		fakeResponse,
	)

	ctx := context.Background()

	config, err := api.KubernetesClustersTokens.Get(ctx, clusterID, tokenID)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if config == "" {
		t.Errorf("Empty response")
	}
}

func TestKubernetesClustersTokensService_GetDefault(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: kubernetesClustersConfigGetResponse, statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs/default", clusterID),
		fakeResponse,
	)

	ctx := context.Background()

	config, err := api.KubernetesClustersTokens.GetDefault(ctx, clusterID)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if config == "" {
		t.Errorf("Empty response")
	}
}

func TestKubernetesClustersTokensService_List(t *testing.T) {
	var expectedResult kubernetesClustersTokenRoot

	fakeResponse := &fakeServerResponse{responseBody: kubernetesClustersTokenListResponse, statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs", clusterID),
		fakeResponse,
	)

	ctx := context.Background()

	tokens, err := api.KubernetesClustersTokens.List(ctx, clusterID, nil)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if err := json.Unmarshal([]byte(kubernetesClustersTokenListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}
	if !reflect.DeepEqual(expectedResult.KubernetesTokens, tokens) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, tokens)
	}
}

func TestKubernetesClustersTokensService_Create(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs", clusterID),
		fakeResponse,
	)

	createRequest := KubernetesClustersTokenCreateRequest{Name: "Test Kubernetes Token"}

	ctx := context.Background()

	err := api.KubernetesClustersTokens.Create(ctx, clusterID, &createRequest)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}

func TestKubernetesClustersTokensService_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 200}
	api, _ := newFakeAPIClient(
		fmt.Sprintf("/api/v2/kubernetes/clusters/%s/kubeconfigs/%s", clusterID, tokenID),
		fakeResponse,
	)

	ctx := context.Background()

	err := api.KubernetesClustersTokens.Delete(ctx, clusterID, tokenID)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
