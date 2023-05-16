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
	"net/http"
	"reflect"
	"testing"
)

var (
	accessTokenResponse = `{
		"id": "104637da-e90d-44e5-9f39-e0a99ccef14f",
		"name": "test",
		"token": "test",
		"scopes": [
			"email"
		],
		"created_at": "2021-09-30T12:00:00Z"
	}`
	UserID = "ce38d0df-edae-43a0-938e-6efa4f3708f0"
)

var accessTokensResponse = fmt.Sprintf(`[%s]`, accessTokenResponse)

func TestAccessToken_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: accessTokensResponse}
	path := fmt.Sprintf("/api/internal/users/%s/access_tokens", UserID)
	server := newFakeServer(path, fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	accessTokens, err := api.AccessTokens.List(ctx, UserID, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult []AccessToken
	if err := json.Unmarshal([]byte(accessTokensResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult, accessTokens) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, accessTokens)
	}
}

func TestAccessToken_Create(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: accessTokenResponse}
	path := fmt.Sprintf("/api/internal/users/%s/access_tokens", UserID)
	api, _ := newFakeAPIClient(path, fakeResponse)

	ctx := context.Background()

	accessToken, err := api.AccessTokens.Create(ctx, UserID, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult *AccessToken
	if err := json.Unmarshal([]byte(accessTokenResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult, accessToken) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, accessToken)
	}
}

func TestAccessToken_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: http.StatusNoContent}
	path := fmt.Sprintf("/api/internal/users/%s/access_tokens/%s", UserID, "104637da-e90d-44e5-9f39-e0a99ccef14f")
	api, _ := newFakeAPIClient(path, fakeResponse)

	ctx := context.Background()

	err := api.AccessTokens.Delete(ctx, UserID, "104637da-e90d-44e5-9f39-e0a99ccef14f")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
