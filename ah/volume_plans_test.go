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

const volumePlanResponse = `{
	"id": 380171553,
	"type": "volume",
	"currency": "usd",
	"name": "HDD. Level 2",
	"data": [],
	"custom_attributes": {
		"hot": true,
		"slug": "vsl-hddl2",
		"max_size": 1000000,
		"min_size": 1,
		"volume_type": {
			"id": "987768d4-9049-4d1e-ad70-c73ffa50bf58",
			"disk_type": "HDD",
			"replication_level": 2
		},
		"datacenter_ids": [
			"62893e3a-84e7-46a4-9cc2-ad892ef37a48"
		],
		"websaProductId": "a391bb0a-32ff-4702-9060-cbbf59b247f1",
		"predefined_sizes": [
			{
				"name": "min",
				"size": 2
			},
			{
				"name": "max",
				"size": 3
			}
		]
	},
	"prices": {
		"380171554": {
			"id": 380171554,
			"plan_id": 380171553,
			"type": "overuse,volume_du",
			"unit": "gb",
			"currency": "usd",
			"quantity": "0",
			"price": "5.00",
			"object_id": null
		}
	}
}`

var (
	volumePlansListResponse = fmt.Sprintf(`{"data": [%s]}`, volumePlanResponse)
)

func TestVolumePlans_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: volumePlansListResponse}
	server := newFakeServer("/api/v1/plans/public", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	volumePlans, err := api.VolumePlans.List(ctx)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult volumePlansRoot
	if err = json.Unmarshal([]byte(volumePlansListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.Plans, volumePlans) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volumePlans)
	}

}
