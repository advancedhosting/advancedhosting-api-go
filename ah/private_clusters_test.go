/*
Copyright 2022 Advanced Hosting

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

const privateClusterResponse = `{
	"id": "6770e666-7a7b-4e9f-816d-cfc98a52c84d",
	"name": "PrivateCloud AMS1 Test ",
	"ip_addresses_count": 0,
	"datacenter_id": "1b1ae192-d44e-451b-8d39-a8670c58e97d",
	"networks": [{
		"id": "c1dc64cd-075a-486d-ad45-3b63614eade5",
		"name": "185.162.85.176/28"
	}, {
		"id": "4d572f8b-e386-4357-8ae1-81210abb1f91",
		"name": "185.191.170.0/24"
	}],
	"nodes": [{
		"id": "2486b2f8-f7a6-4207-979b-9b94d93c174e",
		"name": "DS5199",
		"cpu": 40,
		"vcpu": 400,
		"disk": 858,
		"cpu_name": "Intel(R) Xeon(R) CPU E5-2630 v4",
		"cpu_frequency": "2.20GHz",
		"ram": 121190,
		"vcpu_available": 400,
		"disk_available": 858,
		"ram_available": 121190,
		"cpu_type": "Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz",
		"instances_count": 0,
		"ip_addresses_count": 0
	}]
}`

var (
	privateClusterListResponse = fmt.Sprintf(`{"private_clusters": [%s]}`, privateClusterResponse)
)

func TestPrivateClusters_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: privateClusterListResponse}
	server := newFakeServer("/api/v1/clusters/private", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateClusters, err := api.PrivateClusters.List(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult privateClustersRoot
	if err = json.Unmarshal([]byte(privateClusterListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.PrivateClusters, privateClusters) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateClusters)
	}
}
