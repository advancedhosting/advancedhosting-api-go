/*
Copyright 2021 Advanced Hosting

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

const privateNetworkResponse = `{
    "id": "06fb6b08-69f5-4b53-b88c-1694622e79c5",
    "number": "NET1001869",
    "cidr": "10.0.0.0/24",
    "name": "Default network",
    "state": "active",
    "created_at": "2020-04-15T10:51:15.765Z",
    "last_action": null,
    "instances_count": 0
}`

const privateNetworkInfoResponse = `{
	"id": "1bb35cbf-4b0f-467f-aa12-343e896e2d22",
	"number": "NET1002426",
	"cidr": "10.0.2.0/24",
	"name": "Kube1.18PrivateNetwork",
	"state": "active",
	"created_at": "2020-07-16T08:00:53.277Z",
	"last_action": {
		"id": "27108719-e5cd-4b73-8f66-efbcda96b166",
		"resource_id": "1bb35cbf-4b0f-467f-aa12-343e896e2d22",
		"state": "success",
		"created_at": "2020-08-05T08:28:30.738Z",
		"resource_type": "privatenetwork",
		"type": "update"
	},
	"instance_private_networks": [{
		"id": "ba34729d-ab52-414e-a407-f91ca83e29a2",
		"ip": "10.0.2.4",
		"mac_address": "00:16:3e:37:57:d2",
		"state": "connected",
		"connected_at": "2020-08-05T08:28:42.909Z",
		"instance": {
			"id": "b609e1a7-0e80-469c-8113-01efab7290fe",
			"image_id": "f0438a4b-7c4a-4a63-a593-8e619ec63d16",
			"name": "LoadBalancer",
			"number": "WVDS113904",
			"private_network_instructions": []
		}
	}]
}`

var (
	privateNetworkListResponse = fmt.Sprintf(`{"private_networks": [%s]}`, privateNetworkResponse)
	privateNetworkGetResponse  = fmt.Sprintf(`{"private_network": %s}`, privateNetworkInfoResponse)
)

func TestPrivateNetworks_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: privateNetworkListResponse}
	server := newFakeServer("/api/v1/private_networks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetworks, err := api.PrivateNetworks.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult privateNetworksRoot
	if err = json.Unmarshal([]byte(privateNetworkListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.PrivateNetworks, privateNetworks) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateNetworks)
	}
}

func TestPrivateNetworks_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: privateNetworkGetResponse}
	server := newFakeServer("/api/v1/private_networks/1bb35cbf-4b0f-467f-aa12-343e896e2d22", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult privateNetworkInfoRoot
	if err := json.Unmarshal([]byte(privateNetworkGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	privateNetwork, err := api.PrivateNetworks.Get(ctx, "1bb35cbf-4b0f-467f-aa12-343e896e2d22")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if privateNetwork == nil || privateNetwork.ID != "1bb35cbf-4b0f-467f-aa12-343e896e2d22" {
		t.Errorf("Invalid response: %v", privateNetwork)
	}

	if !reflect.DeepEqual(expectedResult.PrivateNetwork, privateNetwork) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateNetwork)
	}
}

func TestPrivateNetworks_Update(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: privateNetworkGetResponse}
	server := newFakeServer("/api/v1/private_networks/1bb35cbf-4b0f-467f-aa12-343e896e2d22", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult privateNetworkInfoRoot
	if err := json.Unmarshal([]byte(privateNetworkGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	request := &PrivateNetworkUpdateRequest{
		Name: "aaaa",
		CIDR: "10.0.3.6/24",
	}

	privateNetwork, err := api.PrivateNetworks.Update(ctx, "1bb35cbf-4b0f-467f-aa12-343e896e2d22", request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if privateNetwork == nil || privateNetwork.ID != "1bb35cbf-4b0f-467f-aa12-343e896e2d22" {
		t.Errorf("Invalid response: %v", privateNetwork)
	}

	if !reflect.DeepEqual(expectedResult.PrivateNetwork, privateNetwork) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateNetwork)
	}
}

func TestPrivateNetworks_Create(t *testing.T) {
	request := &PrivateNetworkCreateRequest{
		Name: "test-name",
		CIDR: "10.0.0.0/24",
		InstancePrivateNetworkAttributes: []InstancePrivateNetworkAttributes{
			{
				InstanceID: "test-id",
			},
		},
	}

	fakeResponse := &fakeServerResponse{
		responseBody: privateNetworkGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/private_networks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetwork, err := api.PrivateNetworks.Create(ctx, request)

	if privateNetwork == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestPrivateNetworks_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{}
	server := newFakeServer("/api/v1/private_networks/test_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.PrivateNetworks.Delete(ctx, "test_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
