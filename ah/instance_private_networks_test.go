/*
Copyright 2020 Advanced Hosting

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

const instancePrivateNetworkResponse = `{
	"id": "bb68f57c-15c6-4fd4-b388-a6c57cff36ff",
	"ip": "10.0.1.3",
	"mac_address": "00:16:3e:61:76:c6",
	"state": "connected",
	"connected_at": "2020-07-08T13:09:00.947Z",
	"instance": {
		"id": "2a758843-b82c-435d-b2b2-65581361345b",
		"image_id": "f0438a4b-7c4a-4a63-a593-8e619ec63d16",
		"name": "ExternalLoadBalancerNewSchema",
		"number": "WVDS113987",
		"private_network_instructions": []
	},
	"private_network": {
		"id": "17a1b879-5354-4118-b572-a73688523035",
		"number": "NET1002373",
		"cidr": "10.0.1.0/24",
		"name": "KubePrivateNetworkNewSchema",
		"state": "active",
		"created_at": "2020-07-08T10:23:43.938Z",
		"last_action": {
			"id": "b55d976a-e7df-4d31-b357-4d67e4bef42e",
			"resource_id": "17a1b879-5354-4118-b572-a73688523035",
			"state": "success",
			"created_at": "2020-07-08T13:08:50.543Z",
			"resource_type": "privatenetwork",
			"type": "update"
		},
		"instances_count": 3
	}
}
`

var (
	instancePrivateNetworkGetResponse = fmt.Sprintf(`{"instance_private_network": %s}`, instancePrivateNetworkResponse)
)

func TestInstancePrivateNetworks_Create(t *testing.T) {
	request := &InstancePrivateNetworkCreateRequest{
		PrivateNetworkID: "test_network_id",
		InstanceID:       "test_instance_id",
		IP:               "10.10.0.1",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: instancePrivateNetworkGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/instance_private_networks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetwork, err := api.InstancePrivateNetworks.Create(ctx, request)

	if privateNetwork == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestInstancePrivateNetworks_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: instancePrivateNetworkGetResponse}
	server := newFakeServer("/api/v1/instance_private_networks/bb68f57c-15c6-4fd4-b388-a6c57cff36ff", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	instancePrivateNetwork, err := api.InstancePrivateNetworks.Get(ctx, "bb68f57c-15c6-4fd4-b388-a6c57cff36ff")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if instancePrivateNetwork == nil || instancePrivateNetwork.ID != "bb68f57c-15c6-4fd4-b388-a6c57cff36ff" {
		t.Errorf("Invalid response: %v", "bb68f57c-15c6-4fd4-b388-a6c57cff36ff")
	}

	var expectedResult instancePrivateNetworkInfoRoot
	json.Unmarshal([]byte(instancePrivateNetworkGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.InstancePrivateNetwork, instancePrivateNetwork) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instancePrivateNetwork)
	}

}

func TestInstancePrivateNetworks_Update(t *testing.T) {
	request := &InstancePrivateNetworkUpdateRequest{
		IP: "10.10.0.2",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: instancePrivateNetworkGetResponse,
		statusCode:   200,
	}

	server := newFakeServer("/api/v1/instance_private_networks/1bb35cbf-4b0f-467f-aa12-343e896e2d22", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetwork, err := api.InstancePrivateNetworks.Update(ctx, "1bb35cbf-4b0f-467f-aa12-343e896e2d22", request)

	if privateNetwork == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestInstancePrivateNetworks_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{
		responseBody: instancePrivateNetworkGetResponse,
		statusCode:   200,
	}

	server := newFakeServer("/api/v1/instance_private_networks/1bb35cbf-4b0f-467f-aa12-343e896e2d22", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetwork, err := api.InstancePrivateNetworks.Delete(ctx, "1bb35cbf-4b0f-467f-aa12-343e896e2d22")

	if privateNetwork == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}
