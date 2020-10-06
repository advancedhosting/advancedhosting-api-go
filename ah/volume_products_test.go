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

const volumeProductResponse = `{
	"id": "ff4ae08e-d510-4e85-8440-9fdfd0f2308a",
	"created_at": "2019-05-27T15:05:58.292Z",
	"updated_at": "2019-09-16T13:43:23.542Z",
	"name": "HDD. Level 2 AMS1",
	"type": "VolumeProduct",
	"price": "0.02",
	"currency": "USD",
	"hot": false,
	"tariff": [
		{
			"component": "volume",
			"tariff_price": 0.02,
			"monthly_hours": 672,
			"measure": "GB"
		}
	],
	"min_size": 10,
	"max_size": 10000,
	"datacenter_ids": [
		"1b1ae192-d44e-451b-8d39-a8670c58e97d"
	],
	"predefined_sizes": [],
	"category": {
		"id": "2463236e-b47a-4f43-b7a8-6cffbb25f79b",
		"name": "Standard"
	},
	"volume_type": {
		"id": "f7d7f049-f767-49bd-93b0-2f3702ca4270",
		"name": "HDD 2",
		"description": "",
		"disk_type": "HDD",
		"replication_level": 2
	}
}`

var (
	volumeProductListResponse = fmt.Sprintf(`{"products": [%s], "meta":{"page": 1,"per_page": 25,"total": 4}}`, volumeProductResponse)
)

func TestVolumeProducts_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: volumeProductListResponse}
	server := newFakeServer("/api/v1/products/volumes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	options := &ListOptions{
		Filters: []FilterInterface{
			&EqFilter{
				Keys:  []string{"id"},
				Value: "ggg",
			},
		},
	}
	volumeProducts, meta, err := api.VolumeProducts.List(ctx, options)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult productsRoot
	json.Unmarshal([]byte(volumeProductListResponse), &expectedResult)

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	if !reflect.DeepEqual(expectedResult.Products, volumeProducts) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volumeProducts)
	}

}

func TestVolumeProducts_ListWithoutOptions(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: volumeProductListResponse}
	server := newFakeServer("/api/v1/products/volumes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	volumeProducts, meta, err := api.VolumeProducts.List(ctx, nil)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult productsRoot
	json.Unmarshal([]byte(volumeProductListResponse), &expectedResult)

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	if !reflect.DeepEqual(expectedResult.Products, volumeProducts) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, volumeProducts)
	}

}
