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
	"fmt"
	"net/http"
)

// SSHKey object
type SSHKey struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

// SSHKeysAPI is an interface for ssh keys.
type SSHKeysAPI interface {
	List(context.Context, *ListOptions) ([]SSHKey, *Meta, error)
	Get(context.Context, string) (*SSHKey, error)
	Create(context.Context, *SSHKeyCreateRequest) (*SSHKey, error)
	Update(context.Context, string, *SSHKeyUpdateRequest) (*SSHKey, error)
	Delete(context.Context, string) error
}

// SSHKeysService implements SSHKeysAPI interface.
type SSHKeysService struct {
	client *APIClient
}

type sshKeysRoot struct {
	Meta    *Meta    `json:"meta"`
	SSHKeys []SSHKey `json:"ssh_keys"`
}

// List returns all available ssh keys
func (sk *SSHKeysService) List(ctx context.Context, options *ListOptions) ([]SSHKey, *Meta, error) {
	path := "api/v1/ssh_keys"

	var sshRoot sshKeysRoot

	if err := sk.client.list(ctx, path, options, &sshRoot); err != nil {
		return nil, nil, err
	}

	return sshRoot.SSHKeys, sshRoot.Meta, nil
}

type sshKeyRoot struct {
	SSHKey *SSHKey `json:"ssh_key"`
	Meta   *Meta   `json:"meta"`
}

// Get ssh key info
func (sk *SSHKeysService) Get(ctx context.Context, sshKeyID string) (*SSHKey, error) {
	path := fmt.Sprintf("api/v1/ssh_keys/%s", sshKeyID)
	req, err := sk.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var sshRoot sshKeyRoot
	_, err = sk.client.Do(ctx, req, &sshRoot)

	if err != nil {
		return nil, err
	}

	return sshRoot.SSHKey, nil
}

// SSHKeyCreateRequest object
type SSHKeyCreateRequest struct {
	Name      string `json:"name"`
	PublicKey string `json:"public_key,omitempty"`
}

// Create ssh key
func (sk *SSHKeysService) Create(ctx context.Context, createRequest *SSHKeyCreateRequest) (*SSHKey, error) {

	type request struct {
		SSHKey *SSHKeyCreateRequest `json:"ssh_key"`
	}
	req, err := sk.client.newRequest(http.MethodPost, "api/v1/ssh_keys", &request{createRequest})
	if err != nil {
		return nil, err
	}

	var sshRoot sshKeyRoot
	if _, err := sk.client.Do(ctx, req, &sshRoot); err != nil {
		return nil, err
	}

	return sshRoot.SSHKey, nil
}

// SSHKeyUpdateRequest object
type SSHKeyUpdateRequest struct {
	Name      string `json:"name,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
}

// Update ssh key
func (sk *SSHKeysService) Update(ctx context.Context, sshKeyID string, updateRequest *SSHKeyUpdateRequest) (*SSHKey, error) {
	path := fmt.Sprintf("api/v1/ssh_keys/%s", sshKeyID)
	req, err := sk.client.newRequest(http.MethodPut, path, updateRequest)
	if err != nil {
		return nil, err
	}

	var sshRoot sshKeyRoot
	if _, err := sk.client.Do(ctx, req, &sshRoot); err != nil {
		return nil, err
	}

	return sshRoot.SSHKey, nil
}

// Delete ssh key
func (sk *SSHKeysService) Delete(ctx context.Context, sshKeyID string) error {
	path := fmt.Sprintf("api/v1/ssh_keys/%s", sshKeyID)
	req, err := sk.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	if _, err = sk.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
