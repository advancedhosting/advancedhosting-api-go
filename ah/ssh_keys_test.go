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

const sshKeyResponse = `{
	"id": "c1965765-c1b5-4c1e-a55a-aa0a208cd89e",
	"name": "test@test.com",
	"fingerprint": "f0:c0:00:55:6b:86:46:c7:a0:34:ff:ff:ff:ff:ff:f",
	"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQA test@test.com",
	"created_at": "2020-09-18T07:30:03.057Z"
}`

var (
	sshKeyListResponse = fmt.Sprintf(`{"ssh_keys": [%s], "meta":{"page": 1,"per_page": 25,"total": 4}}`, sshKeyResponse)
	sshKeyGetResponse  = fmt.Sprintf(`{"ssh_key": %s}`, sshKeyResponse)
)

func TestSSHKeys_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: sshKeyListResponse}
	server := newFakeServer("/api/v1/ssh_keys", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	sshKeys, meta, err := api.SSHKeys.List(ctx, nil)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult sshKeysRoot
	if err = json.Unmarshal([]byte(sshKeyListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if meta == nil {
		t.Errorf("unexpected meta: %v", meta)
	}

	if !reflect.DeepEqual(expectedResult.SSHKeys, sshKeys) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, sshKeys)
	}

}

func TestSSHKeys_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: sshKeyGetResponse}
	server := newFakeServer("/api/v1/ssh_keys/c1965765-c1b5-4c1e-a55a-aa0a208cd89e", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult sshKeyRoot
	if err := json.Unmarshal([]byte(sshKeyGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	sshKey, err := api.SSHKeys.Get(ctx, "c1965765-c1b5-4c1e-a55a-aa0a208cd89e")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sshKey == nil || sshKey.ID != "c1965765-c1b5-4c1e-a55a-aa0a208cd89e" {
		t.Errorf("Invalid response: %v", sshKey)
	}

	if !reflect.DeepEqual(expectedResult.SSHKey, sshKey) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, sshKey)
	}
}

func TestSSHKeys_Create(t *testing.T) {
	request := &SSHKeyCreateRequest{
		Name:      "test@test.com",
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQA test@test.com",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: sshKeyGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/ssh_keys", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	sshKey, err := api.SSHKeys.Create(ctx, request)

	if sshKey == nil {
		t.Errorf("Empty response")
	}

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	var expectedResult sshKeyRoot
	if err = json.Unmarshal([]byte(sshKeyGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.SSHKey, sshKey) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, sshKey)
	}

}

func TestSSHKeys_Update(t *testing.T) {
	request := &SSHKeyUpdateRequest{
		Name:      "test@test.com",
		PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQA test@test.com",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: sshKeyGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/ssh_keys/c1965765-c1b5-4c1e-a55a-aa0a208cd89e", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	var expectedResult sshKeyRoot
	if err := json.Unmarshal([]byte(sshKeyGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected unmarshal error: %v", err)
	}

	ctx := context.Background()
	sshKey, err := api.SSHKeys.Update(ctx, "c1965765-c1b5-4c1e-a55a-aa0a208cd89e", request)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if sshKey == nil || sshKey.ID != "c1965765-c1b5-4c1e-a55a-aa0a208cd89e" {
		t.Errorf("Invalid response: %v", sshKey)
	}

	if !reflect.DeepEqual(expectedResult.SSHKey, sshKey) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, sshKey)
	}

}

func TestSSHKeys_Delete(t *testing.T) {
	fakeResponse := &fakeServerResponse{}
	server := newFakeServer("/api/v1/ssh_keys/test_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.SSHKeys.Delete(ctx, "test_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
