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
	"fmt"
	"testing"
)

const ipAddressAssignmentResponse = `{
	"id": "aa292b3d-09eb-440e-be80-9bba7556a862",
	"instance_id": "b609e1a7-0e80-469c-8113-01efab7290fe",
	"ip_address_id": "fc106c9b-df92-4f8c-823c-2f3ce1972d5d",
	"state": "attaching",
	"created_at": "2020-08-21T09:45:40.778Z",
	"updated_at": "2020-08-21T09:45:40.778Z"
}`

var (
	ipAddressAssignmentGetResponse = fmt.Sprintf(`{"instance_ip_address": %s}`, ipAddressResponse)
)

func TestIPAddressAssignment_Create(t *testing.T) {
	request := &IPAddressAssignmentCreateRequest{
		IPAddressID: "fc106c9b-df92-4f8c-823c-2f3ce1972d5d",
		InstanceID:  "b609e1a7-0e80-469c-8113-01efab7290fe",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: ipAddressAssignmentGetResponse,
		statusCode:   200,
	}

	server := newFakeServer("/api/v1/instance_ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	ipAddressAssignment, err := api.IPAddressAssignments.Create(ctx, request)

	if ipAddressAssignment == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestIPAddressAssignment_Delete(t *testing.T) {
	server := newFakeServer("/api/v1/instance_ip_addresses/test_id", &fakeServerResponse{})

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.IPAddressAssignments.Delete(ctx, "test_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
