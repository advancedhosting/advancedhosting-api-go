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

const instanceProductResponse = `{
	"id": "df42a96b-b381-412c-a605-d66d7bf081af",
	"created_at": "2020-04-02T15:58:08.150Z",
	"updated_at": "2020-04-02T15:58:48.490Z",
	"name": "Start XS",
	"type": "VpsProduct",
	"price": "3.0",
	"currency": "USD",
	"hot": false,
	"tariff": [
		{
			"component": "basic_resources",
			"tariff_price": 3,
			"additional_service_price": 0,
			"included_value": 0,
			"measure": "hour"
		},
		{
			"component": "scheduled_backup",
			"tariff_price": 0.3,
			"additional_service_price": 0,
			"included_value": 0,
			"measure": "hour"
		}
	],
	"vcpu": "1",
	"ram": "1024",
	"disk": "12",
	"traffic": "5000",
	"available_on_trial": true,
	"slug": null,
	"category": {
		"id": "2463236e-b47a-4f43-b7a8-6cffbb25f79b",
		"name": "Standard"
	}
}`

var (
	instanceProductListResponse = fmt.Sprintf(`{"products": [%s], "meta":{"page": 1,"per_page": 25,"total": 4}}`, instanceProductResponse)
)

func TestInstanceProducts_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: instanceProductListResponse}
	server := newFakeServer("/api/v1/products/instances", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	instanceProducts, meta, err := api.InstanceProducts.List(ctx, nil)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult instanceProductsRoot
	if err = json.Unmarshal([]byte(instanceProductListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	if !reflect.DeepEqual(expectedResult.Products, instanceProducts) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, instanceProducts)
	}

}
