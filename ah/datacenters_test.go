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

const datacenterResponse = `{
	"id": "62893e3a-84e7-46a4-9cc2-ad892ef37a48",
	"name": "Ams1",
	"full_name": "NL, Amsterdam 1, Ams1",
	"datacenter_slug": null,
	"instances_running": 1,
	"private_nodes_count": 0,
	"region": {
		"id": "8bff9701-0e3d-44f5-9778-27c8716cdf86",
		"name": "Amsterdam 1",
		"country_code": "NL"
	}
}`

var (
	datacenterListResponse = fmt.Sprintf(`{"products": [%s]}`, datacenterResponse)
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

	var expectedResult datacentersRoot
	json.Unmarshal([]byte(datacenterListResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.Datacenters, datacenters) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, datacenters)
	}

}