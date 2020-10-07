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

const volumeResponse = `{
	"id": "e88cb60e-828f-416f-8ab0-e05ab4493b1a",
	"name": "test api volume",
	"file_system": "ext4",
	"number": "VOL1001188",
	"size": 10,
	"port": 2,
	"state": "attached",
	"original_id": null,
	"created_at": "2020-07-27T13:15:24.730Z",
	"attached_at": "2020-07-27T13:15:40.278Z",
	"product_id": "03bebb65-22d8-43c6-819b-5b85b5e49c82",
	"current_action": null,
	"last_action": {
		"id": "3305c748-d8e7-4ec8-be7e-f1c21385bb0d",
		"resource_id": "e88cb60e-828f-416f-8ab0-e05ab4493b1a",
		"state": "success",
		"created_at": "2020-07-27T13:15:24.736Z",
		"resource_type": "volume",
		"type": "create"
	},
	"instance": {
		"id": "2a758843-b82c-435d-b2b2-65581361345b",
		"name": "ExternalLoadBalancerNewSchema"
	},
	"product": {
		"id": "03bebb65-22d8-43c6-819b-5b85b5e49c82",
		"name": "HDD. Level 2 ASH1",
		"min_volume_size": 10,
		"max_volume_size": 10000
	},
	"volume_pool": {
		"name": "hdd2",
		"datacenter_ids": [
			"c54e8896-53d8-479a-8ff1-4d7d9d856a50"
		],
		"replication_level": 2
	}
}`

var (
	volumeListResponse = fmt.Sprintf(`{"volumes": [%s], "meta":{"page": 1,"per_page": 25,"total": 4}}`, volumeResponse)
	volumeGetResponse  = fmt.Sprintf(`{"volume": %s}`, volumeResponse)
)

func TestVolumes_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: volumeListResponse}
	server := newFakeServer("/api/v1/volumes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	volumes, meta, err := api.Volumes.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult volumesRoot
	json.Unmarshal([]byte(volumeListResponse), &expectedResult)

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	if !reflect.DeepEqual(expectedResult.Volumes, volumes) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volumes)
	}
}

func TestVolumes_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: volumeGetResponse}
	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult volumeRoot
	json.Unmarshal([]byte(volumeGetResponse), &expectedResult)

	volume, err := api.Volumes.Get(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if volume == nil || volume.ID != "e88cb60e-828f-416f-8ab0-e05ab4493b1a" {
		t.Errorf("Invalid response: %v", volume)
	}

	if !reflect.DeepEqual(expectedResult.Volume, volume) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volume)
	}
}

func TestVolumes_Update(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: volumeGetResponse}
	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult volumeRoot
	json.Unmarshal([]byte(volumeGetResponse), &expectedResult)

	request := &VolumeUpdateRequest{
		Name: "New Name",
	}

	volume, err := api.Volumes.Update(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a", request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if volume == nil || volume.ID != "e88cb60e-828f-416f-8ab0-e05ab4493b1a" {
		t.Errorf("Invalid response: %v", volume)
	}

	if !reflect.DeepEqual(expectedResult.Volume, volume) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volume)
	}
}

func TestVolumes_Copy(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse}
	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a/actions", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult volumeActionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	request := &VolumeCopyActionRequest{
		Name:      "new name",
		ProductID: "test_product_id",
	}

	action, err := api.Volumes.Copy(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a", request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if action == nil {
		t.Errorf("Invalid response: %v", action)
	}

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestVolumes_CopyWithProductSlug(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse}
	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a/actions", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult volumeActionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	request := &VolumeCopyActionRequest{
		Name:        "new name",
		ProductSlug: "test_product_slug",
	}

	action, err := api.Volumes.Copy(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a", request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if action == nil {
		t.Errorf("Invalid response: %v", action)
	}

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestVolumes_Resize(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse}
	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a/actions", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult actionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	action, err := api.Volumes.Resize(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a", 20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if action == nil {
		t.Errorf("Invalid response: %v", action)
	}

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestVolumes_Create(t *testing.T) {
	request := &VolumeCreateRequest{
		Name:       "test-name",
		Size:       50,
		ProductID:  "Test_product_id",
		FileSystem: "ext4",
		InstanceID: "test_instance_id",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: volumeGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/volumes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	volume, err := api.Volumes.Create(ctx, request)

	if volume == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestVolumes_CreateWithSlug(t *testing.T) {
	request := &VolumeCreateRequest{
		Name:        "test-name",
		Size:        50,
		ProductSlug: "Test_product_id",
		FileSystem:  "ext4",
		InstanceID:  "test_instance_id",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: volumeGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/volumes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	volume, err := api.Volumes.Create(ctx, request)

	if volume == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}
}

func TestVolumes_ActionInfo(t *testing.T) {
	actionGetResponse := `{
		"action": {
			"id": "7dc9faa7-6049-432e-8576-00313cb0cafe",
			"resource_id": "f90558e9-c66c-4ad9-8760-a26a162b07e2",
			"state": "success",
			"created_at": "2020-09-21T08:37:18.047Z",
			"resource_type": "volume",
			"type": "copy",
			"user_id": "de1c6534-0782-45c7-948f-522c644c9240",
			"note": null,
			"updated_at": "2020-09-21T08:37:33.868Z",
			"started_at": "2020-09-21T08:37:18.071Z",
			"completed_at": "2020-09-21T08:37:33.860Z",
			"result_params": {
				"copied_volume_id": "fcd60ac7-b119-4a5e-bd96-6d90983a3e22"
			}
		}
	}
	`
	fakeResponse := &fakeServerResponse{responseBody: actionGetResponse, statusCode: 200}

	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a/actions/7dc9faa7-6049-432e-8576-00313cb0cafe", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	action, err := api.Volumes.ActionInfo(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a", "7dc9faa7-6049-432e-8576-00313cb0cafe")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult volumeActionRoot
	json.Unmarshal([]byte(actionGetResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Action, action) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, action)
	}
}

func TestVolumes_Actions(t *testing.T) {
	actionListResponse := `{
		"actions": [
			{
				"id": "7dc9faa7-6049-432e-8576-00313cb0cafe",
				"resource_id": "f90558e9-c66c-4ad9-8760-a26a162b07e2",
				"state": "success",
				"created_at": "2020-09-21T08:37:18.047Z",
				"resource_type": "volume",
				"type": "copy",
				"user_id": "de1c6534-0782-45c7-948f-522c644c9240",
				"note": null,
				"updated_at": "2020-09-21T08:37:33.868Z",
				"started_at": "2020-09-21T08:37:18.071Z",
				"completed_at": "2020-09-21T08:37:33.860Z",
				"result_params": {
					"copied_volume_id": "fcd60ac7-b119-4a5e-bd96-6d90983a3e22"
				}
			},
			{
				"id": "6dbdc31e-e1e3-4858-a62f-de1a91a7f3e6",
				"resource_id": "f90558e9-c66c-4ad9-8760-a26a162b07e2",
				"state": "success",
				"created_at": "2020-06-19T11:48:41.603Z",
				"resource_type": "volume",
				"type": "build",
				"user_id": null,
				"note": null,
				"updated_at": "2020-06-19T11:48:48.725Z",
				"started_at": "2020-06-19T11:48:41.614Z",
				"completed_at": "2020-06-19T11:48:48.721Z",
				"result_params": {}
			}
		]
	}`
	fakeResponse := &fakeServerResponse{responseBody: actionListResponse, statusCode: 200}

	server := newFakeServer("/api/v1/volumes/e88cb60e-828f-416f-8ab0-e05ab4493b1a/actions", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	actions, err := api.Volumes.Actions(ctx, "e88cb60e-828f-416f-8ab0-e05ab4493b1a")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	var expectedResult volumeActionsRoot
	json.Unmarshal([]byte(actionListResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Actions, actions) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, actions)
	}

	if actions[0].ResultParams.CopiedVolumeID != "fcd60ac7-b119-4a5e-bd96-6d90983a3e22" {
		t.Errorf("unexpected copied_volume_id, expected fcd60ac7-b119-4a5e-bd96-6d90983a3e22. got: %s", actions[0].ResultParams.CopiedVolumeID)
	}
}

func TestVolumes_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{}
	server := newFakeServer("/api/v1/volumes/test_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.Volumes.Delete(ctx, "test_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
