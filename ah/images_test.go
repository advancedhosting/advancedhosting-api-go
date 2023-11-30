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

const imageResponse = `{
	"id": "b62166ec-4271-411e-920a-2d1f98c09567",
	"created_at": "2019-08-08T08:59:15.112Z",
	"updated_at": "2020-02-28T10:35:14.302Z",
	"name": "Debian 10 x64",
	"distribution": "Debian",
	"version": "10",
	"architecture": "amd64",
	"public": true,
	"image_slug": null
}`

var (
	imageListResponse = fmt.Sprintf(`{"products": [%s], "meta":{"page": 1,"per_page": 25,"total": 4}}`, imageResponse)
)

func TestImages_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: imageListResponse}
	server := newFakeServer("/api/v1/images", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	images, meta, err := api.Images.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult imagesRoot
	if err = json.Unmarshal([]byte(imageListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	if !reflect.DeepEqual(expectedResult.Images, images) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, images)
	}
}
