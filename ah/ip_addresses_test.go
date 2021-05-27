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

const ipAddressResponse = `{
		"id" : "5a52fcb2-e319-4f14-ae13-901c04341098",
		"reverse_dns" : "ip-185-189-68-65.ah-server.com",
		"updated_at" : "2020-07-08T09:09:57.584Z",
		"created_at" : "2020-07-08T09:09:57.584Z",
		"datacenter_full_name" : "US, Ashburn, ASH1",
		"network_used_for_private_cluster" : false,
		"address" : "185.189.68.65",
		"address_type" : "public",
		"instance_ids" : [
			"a010d05c-e377-4924-a10b-9a23b0bb8663"
		],
		"delete_protection" : false
	}`

var (
	ipAddressesListResponse = fmt.Sprintf(`{"ip_addresses": [%s]}`, ipAddressResponse)
	ipAddressesGetResponse  = fmt.Sprintf(`{"ip_address": %s}`, ipAddressResponse)
)

func TestIPAddresses_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: ipAddressesListResponse}
	server := newFakeServer("/api/v1/ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	ipAddresses, err := api.IPAddresses.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult ipAddressesRoot
	if err = json.Unmarshal([]byte(ipAddressesListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.IPAddresses, ipAddresses) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, ipAddresses)
	}
}

func TestIPAddress_Create(t *testing.T) {
	request := &IPAddressCreateRequest{
		Type:         "public",
		DatacenterID: "test-datacenter-id",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: ipAddressesGetResponse,
		statusCode:   201,
	}

	server := newFakeServer("/api/v1/ip_addresses/", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	ipAddress, err := api.IPAddresses.Create(ctx, request)

	if ipAddress == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestIPAddress_CreateWithSlug(t *testing.T) {
	request := &IPAddressCreateRequest{
		Type:           "public",
		DatacenterSlug: "test-datacenter-slug",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: ipAddressesGetResponse,
		statusCode:   201,
	}

	server := newFakeServer("/api/v1/ip_addresses/", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	ipAddress, err := api.IPAddresses.Create(ctx, request)

	if ipAddress == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestIPAddresses_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: ipAddressesListResponse}
	server := newFakeServer("/api/v1/ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	ipAddr, err := api.IPAddresses.Get(ctx, "5a52fcb2-e319-4f14-ae13-901c04341098")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ipAddr == nil || ipAddr.ID != "5a52fcb2-e319-4f14-ae13-901c04341098" {
		t.Errorf("Invalid response: %v", ipAddr)
	}
}

func TestIPAddresses_GetEmpty(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: `{"ip_addresses": []}`}
	server := newFakeServer("/api/v1/ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	_, err := api.IPAddresses.Get(ctx, "5a52fcb2-e319-4f14-ae13-901c04341098")
	if err != ErrResourceNotFound {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestIPAddresses_GetMultiply(t *testing.T) {
	multipleResponse := fmt.Sprintf(`{"ip_addresses": [%s,%s]}`, ipAddressResponse, ipAddressResponse)
	fakeResponse := &fakeServerResponse{responseBody: multipleResponse}
	server := newFakeServer("/api/v1/ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	_, err := api.IPAddresses.Get(ctx, "5a52fcb2-e319-4f14-ae13-901c04341098")
	if err != ErrResourceNotFound {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestIPAddresses_Update(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: ipAddressesGetResponse}
	server := newFakeServer("/api/v1/ip_addresses/5a52fcb2-e319-4f14-ae13-901c04341098", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	request := &IPAddressUpdateRequest{
		ReverseDNS: "ip-185-189-68-65.ah-server2.com",
	}

	ipAddr, err := api.IPAddresses.Update(ctx, "5a52fcb2-e319-4f14-ae13-901c04341098", request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ipAddr == nil || ipAddr.ID != "5a52fcb2-e319-4f14-ae13-901c04341098" {
		t.Errorf("Invalid response: %v", ipAddr)
	}
}

func TestIPAddresses_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{}
	server := newFakeServer("/api/v1/ip_addresses/test_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.IPAddresses.Delete(ctx, "test_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
