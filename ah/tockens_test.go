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
	"net/http"
	"reflect"
	"testing"
)

const tokenResponse = `{
     "id": "ded95980-05d8-44aa-977d-3dfc1e7966ba",
     "name": "k8s-WKUB100000-token-ca96d",
     "token": "test_token",
     "scopes": ["email"],
     "expires_in": "2021-12-15T17:51:29.765Z",
     "created_at": "2021-12-15T17:51:29.765Z"
 }`

var tokenListResponse = fmt.Sprintf(`[%s]`, tokenResponse)

func TestTokensService_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: tokenListResponse, statusCode: http.StatusOK}
	server := newFakeServer("/id/api/v1/access_tokens", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	tokens, err := api.Tokens.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	var expectedResult []Token
	if err = json.Unmarshal([]byte(tokenListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult, tokens) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, tokens)
	}
}

func TestTokensService_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: tokenResponse, statusCode: http.StatusOK}
	server := newFakeServer("/id/api/v1/access_tokens/ded95980-05d8-44aa-977d-3dfc1e7966ba", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	token, err := api.Tokens.Get(ctx, "ded95980-05d8-44aa-977d-3dfc1e7966ba")
	if err != nil {
		t.Errorf("Error getting token: %v", err)
	}
	var expectedResult *Token
	if err = json.Unmarshal([]byte(tokenResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult, token) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult.Token, token)
	}
}

func TestTokensService_Create(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: tokenResponse, statusCode: http.StatusOK}
	server := newFakeServer("/id/api/v1/access_tokens", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	token, err := api.Tokens.Create(ctx, &TokenCreateRequest{Name: "test"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	var expectedResult *Token
	if err = json.Unmarshal([]byte(tokenResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult, token) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, token)
	}
}

func TestTokensService_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: http.StatusNoContent}
	server := newFakeServer("/id/api/v1/access_tokens/ded95980-05d8-44aa-977d-3dfc1e7966ba", fakeResponse)
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.Tokens.Delete(ctx, "ded95980-05d8-44aa-977d-3dfc1e7966ba")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
