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
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const instanceResponse = `{
	"id": "2a758843-b82c-435d-b2b2-65581361345b",
	"created_at": "2020-07-08T13:06:59.906Z",
	"updated_at": "2020-07-08T13:07:57.180Z",
	"number": "WVDS113987",
	"name": "ExternalLoadBalancerNewSchema",
	"state": "running",
	"state_description": null,
	"locked": false,
	"use_ssh_password": false,
	"product_id": "1a4cdeb2-6ca4-4745-819e-ac2ea99dc0cc",
	"vcpu": 2,
	"ram": 4096,
	"traffic": 5000,
	"tags": [],
	"primary_instance_ip_address_id": "84244ac6-359b-45cb-ab82-7492bbb234fb",
	"ip_scheme": "floating",
	"region": {
		"id": "d2cd5cc1-e822-46ec-a8f6-8d4a7d22fb04",
		"name": "Ashburn",
		"slug": "ash1",
		"country_code": "US",
		"country": "United States of America",
		"city": "Ashburn",
		"parent_id": "f669b77b-9dab-4eb6-9bf3-a84bfc8f9827",
		"group": false,
		"regions_count": 0,
		"services": [
			"vps"
		]
	},
	"datacenter": {
		"id": "c54e8896-53d8-479a-8ff1-4d7d9d856a50",
		"name": "ASH1",
		"full_name": "US, Ashburn, ASH1",
		"instances_running": 0,
		"private_nodes_count": 0,
		"region": {
			"id": "d2cd5cc1-e822-46ec-a8f6-8d4a7d22fb04",
			"name": "Ashburn",
			"country_code": "US"
		}
	},
	"features": [],
	"networks": {
		"v4": [{
			"type": "service",
			"ip_address": "100.124.252.1",
			"netmask": "255.192.0.0",
			"gateway": "100.64.0.1"
		}]
	},
	"current_action": null,
	"last_action": {
		"id": "7b6984e7-ff90-4c3e-8635-23515917d550",
		"resource_id": "2a758843-b82c-435d-b2b2-65581361345b",
		"state": "success",
		"created_at": "2020-07-08T13:07:48.049Z",
		"resource_type": "instance",
		"type": "create_instance_address"
	},
	"reason": null,
	"snapshot_by_schedule": false,
	"snapshot_period": null,
	"max_volumes_number": 0,
	"volumes": [],
	"instance_private_networks": [{
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
	}],
	"disk": 40,
	"instance_ip_addresses": [{
		"id": "84244ac6-359b-45cb-ab82-7492bbb234fb",
		"instance_id": "2a758843-b82c-435d-b2b2-65581361345b",
		"created_at": "2020-07-08T13:07:47.729Z",
		"updated_at": "2020-07-08T13:07:57.179Z",
		"ip_address_id": "ef5beda3-77f0-4e24-ae6b-5780f23e6b5b"
	}],
	"image": {
		"id": "f0438a4b-7c4a-4a63-a593-8e619ec63d16",
		"created_at": "2020-05-19T15:51:42.171Z",
		"updated_at": "2020-05-19T15:59:46.440Z",
		"type": "distribution",
		"name": "Ubuntu 20.04 x64",
		"distribution": "Ubuntu",
		"version": "20.04 LTS",
		"architecture": "x86_64",
		"slug": "Ubuntu-20-04-focal-server-cloudimg-amd64",
		"public": true
	}
}`

const instanceActionResponse = `{
	"id": "2d022304-585c-45f6-95ad-b2f7934cb0eb",
	"resource_id": "91c20eb5-a072-4604-b3e8-ee3de87b06a4",
	"state": "success",
	"created_at": "2020-09-04T16:31:28.189Z",
	"resource_type": "instance",
	"type": "set_primary_ip",
	"user_id": "247e6d8e-b700-4947-9ee4-7a7c89c646b6",
	"note": null,
	"updated_at": "2020-09-04T16:31:29.586Z",
	"started_at": "2020-09-04T16:31:28.198Z",
	"completed_at": "2020-09-04T16:31:29.584Z",
	"result_params": {}
}`

var (
	getResponse        = fmt.Sprintf(`{"instance": %s}`, instanceResponse)
	listResponse       = fmt.Sprintf(`{"instances": [%s]}`, instanceResponse)
	actionGetResponse  = fmt.Sprintf(`{"action": %s}`, instanceActionResponse)
	actionListResponse = fmt.Sprintf(`{"actions": [%s]}`, instanceActionResponse)
)

type fakeServerResponse struct {
	responseBody string
	statusCode   int
}

func newFakeServer(url string, response *fakeServerResponse) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(url, func(rw http.ResponseWriter, r *http.Request) {
		if response.statusCode != 0 {
			rw.WriteHeader(response.statusCode)
		}
		rw.Write([]byte(response.responseBody))
	})
	return httptest.NewServer(mux)
}

func TestInstances_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: listResponse}
	server := newFakeServer("/api/v1/instances", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	instances, _, err := api.Instances.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult instancesRoot
	json.Unmarshal([]byte(listResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Instances, instances) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instances)
	}
}

func TestInstances_ListOptions(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: listResponse}
	server := newFakeServer("/api/v1/instances", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	options := &ListOptions{
		Meta: &ListMetaOptions{
			Page: 1,
		},
	}
	instances, _, err := api.Instances.List(ctx, options)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult instancesRoot
	json.Unmarshal([]byte(listResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Instances, instances) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instances)
	}
}

func TestInstance_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: getResponse}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	instance, err := api.Instances.Get(ctx, "2a758843-b82c-435d-b2b2-65581361345b")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult instanceRoot
	json.Unmarshal([]byte(getResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Instance, instance) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instance)
	}

}

func TestInstance_GetNonExisted(t *testing.T) {
	fakeResponse := &fakeServerResponse{
		responseBody: getResponse,
		statusCode:   404,
	}

	server := newFakeServer("/api/v1/instances/not_existed_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	instance, err := api.Instances.Get(ctx, "2a758843-b82c-435d-b2b2-65581361345b")
	if instance != nil {
		t.Errorf("Unexpected instance %v", instance)
	}

	if err != ErrResourceNotFound {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestInstance_Create(t *testing.T) {
	request := &InstanceCreateRequest{
		Name:                  "Test",
		DatacenterID:          "test-datacenter-id",
		ImageID:               "test-image-id",
		ProductID:             "test-product-id",
		CreatePublicIPAddress: true,
		UseSSHPassword:        true,
	}

	fakeResponse := &fakeServerResponse{
		responseBody: getResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/instances/", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	instance, err := api.Instances.Create(ctx, request)

	if instance == nil {
		t.Errorf("Invalid response %v", instance)
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestInstance_Rename(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: getResponse}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	instance, err := api.Instances.Rename(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "new_name")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult instanceRoot
	json.Unmarshal([]byte(getResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Instance, instance) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instance)
	}
}

func TestInstance_Upgrade(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: getResponse, statusCode: 202}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.Instances.Upgrade(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "new_product_id")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestInstance_Shutdown(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: getResponse, statusCode: 202}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.Instances.Shutdown(ctx, "2a758843-b82c-435d-b2b2-65581361345b")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestInstance_PowerOff(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: getResponse, statusCode: 202}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.Instances.PowerOff(ctx, "2a758843-b82c-435d-b2b2-65581361345b")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestInstance_Destroy(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: getResponse, statusCode: 202}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.Instances.Destroy(ctx, "2a758843-b82c-435d-b2b2-65581361345b")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestInstance_SetPrimaryIP(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse, statusCode: 202}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	action, err := api.Instances.SetPrimaryIP(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "testID")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult actionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}

}

func TestInstance_ActionInfo(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse, statusCode: 200}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions/2a758843-b82c-435d-b2b2-65581361345b", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	action, err := api.Instances.ActionInfo(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "2a758843-b82c-435d-b2b2-65581361345b")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult actionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestInstance_Actions(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionListResponse, statusCode: 200}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	actions, err := api.Instances.Actions(ctx, "2a758843-b82c-435d-b2b2-65581361345b")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult actionsRoot
	json.Unmarshal([]byte(actionListResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Actions, actions) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, actions)
	}
}

func TestInstance_AttachVolume(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse, statusCode: 200}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	action, err := api.Instances.AttachVolume(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "test_volume_id")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult actionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestInstance_DetachVolume(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse, statusCode: 200}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	action, err := api.Instances.DetachVolume(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "test_volume_id")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult actionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestInstance_AvailableVolumes(t *testing.T) {

	fakeResponse := &fakeServerResponse{responseBody: volumeListResponse, statusCode: 200}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/available_volumes", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	volumes, meta, err := api.Instances.AvailableVolumes(ctx, "2a758843-b82c-435d-b2b2-65581361345b", nil)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	var expectedResult volumesRoot
	json.Unmarshal([]byte(volumeListResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Volumes, volumes) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volumes)
	}
}

func TestInstance_CreateBackup(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse, statusCode: 202}

	server := newFakeServer("/api/v1/instances/2a758843-b82c-435d-b2b2-65581361345b/backups", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	action, err := api.Instances.CreateBackup(ctx, "2a758843-b82c-435d-b2b2-65581361345b", "test_backup_note")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult createBackupActionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}
