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
        "ram": 2048,
        "disk": 30,
        "slug": "dc-1-2",
        "vcpu": 1,
        "default": false,
        "traffic": 4000,
        "position": 100,
        "optimized": "cpu",
        "dedicated_cpu": true,
        "fork_on_purchase": false,
        "available_on_trial": false
	},
	"prices": {
		"391445307": {
			"id": 391445307,
			"plan_id": 391445272,
			"type": "monthly,vps",
			"unit": "items",
			"currency": "usd",
			"quantity": "0",
			"price": "21.00",
			"object_id": null,
			"class": "SinglePrice",
			"formula": "cap.monthly(\u002728 days\u0027)"
		},
		"391445308": {
			"id": 391445308,
			"plan_id": 391445272,
			"type": "feature,autobackup",
			"unit": "items",
			"currency": "usd",
			"quantity": "0",
			"price": "2.10",
			"object_id": null,
			"class": "SinglePrice",
			"formula": "cap.monthly(\u002728 days\u0027)"
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
