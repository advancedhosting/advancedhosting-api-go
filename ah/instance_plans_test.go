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

const instancePlanResponse = `{
	"id": 380171663,
	"type": "vps",
	"currency": "usd",
	"name": "512 - low price",
	"data": [],
	"custom_attributes": {
		"hot": false,
		"ram": "1024",
		"disk": "100",
		"slug": "vps512_low",
		"vcpu": "1",
		"traffic": "1",
		"websaProductId": "643c6170-7c67-425c-8c37-45da658d2065",
		"available_on_trial": false
	},
	"prices": {
		"380171664": {
			"id": 380171664,
			"plan_id": 380171663,
			"type": "monthly,vps",
			"unit": "items",
			"currency": "usd",
			"quantity": "0",
			"price": "0.04",
			"object_id": null
		},
		"380171665": {
			"id": 380171665,
			"plan_id": 380171663,
			"type": "feature,autobackup",
			"unit": "items",
			"currency": "usd",
			"quantity": "0",
			"price": "0.00",
			"object_id": null
		}
	}
}`

var (
	instancePlansListResponse = fmt.Sprintf(`{"data": [%s]}`, instancePlanResponse)
)

func TestInstancePlans_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: instancePlansListResponse}
	server := newFakeServer("/api/v1/plans/public", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	instancePlans, err := api.InstancePlans.List(ctx)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult instancePlansRoot
	if err = json.Unmarshal([]byte(instancePlansListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.Plans, instancePlans) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instancePlans)
	}

}
