/*
Copyright 2023 Advanced Hosting

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

const datacenterResponse = `{
	"id": "62893e3a-84e7-46a4-9cc2-ad892ef37a48",
	"name": "Ams1",
	"full_name": "NL, Amsterdam 1, Ams1",
	"slug": "ams1",
	"instances_running": 1,
	"private_nodes_count": 0,
	"region": {
		"id": "8bff9701-0e3d-44f5-9778-27c8716cdf86",
		"name": "Amsterdam 1",
		"country_code": "NL"
	}
}`

var (
	datacenterListResponse = fmt.Sprintf(`{"datacenters": [%s]}`, datacenterResponse)
	datacenterGetResponse  = fmt.Sprintf(`{"datacenter": %s}`, datacenterResponse)
)

func TestDatacenters_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: datacenterListResponse}
	server := newFakeServer("/api/v1/datacenters", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	datacenters, err := api.Datacenters.List(ctx, nil)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if datacenters == nil {
		t.Errorf("Unexpected response: %v", datacenters)
	}

	var expectedResult datacentersRoot
	if err = json.Unmarshal([]byte(datacenterListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.Datacenters, datacenters) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, datacenters)
	}

}

func TestDatacenters_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: datacenterGetResponse}
	server := newFakeServer("/api/v1/datacenters/62893e3a-84e7-46a4-9cc2-ad892ef37a48", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	datacenter, err := api.Datacenters.Get(ctx, "62893e3a-84e7-46a4-9cc2-ad892ef37a48")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if datacenter == nil {
		t.Errorf("Unexpected response: %v", datacenter)
	}

	var expectedResult datacenterRoot
	if err = json.Unmarshal([]byte(datacenterGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.Datacenter, datacenter) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, datacenter)
	}

}
