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

// AccessToken object
type AccessToken struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Token     string `json:"token,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

// AccessTokenCreateRequest object
type AccessTokenCreateRequest struct {
	Name string `json:"name"`
}

// AccessTokensService implements AccessTokensAPI interface.
type AccessTokensService struct {
	client *APIClient
}

// AccessTokensAPI is an interface for access tokens API.
type AccessTokensAPI interface {
	List(context.Context, string, *ListOptions) ([]AccessToken, error)
	Create(context.Context, string, *AccessTokenCreateRequest) (*AccessToken, error)
	Delete(context.Context, string, string) error
}

type AccessTokenRoot struct {
	AccessToken *AccessToken `json:"access_token,omitempty"`
}

// List access tokens
func (ats *AccessTokensService) List(ctx context.Context, userID string, options *ListOptions) ([]AccessToken, error) {
	path := fmt.Sprintf("api/internal/users/%s/access_tokens", userID)

	var accessTokensRoot []AccessToken
	if err := ats.client.list(ctx, path, options, &accessTokensRoot); err != nil {
		return nil, err
	}

	return accessTokensRoot, nil
}

// Create access token
func (ats *AccessTokensService) Create(ctx context.Context, userID string, createRequest *AccessTokenCreateRequest) (*AccessToken, error) {
	path := fmt.Sprintf("api/internal/users/%s/access_tokens", userID)

	req, err := ats.client.newRequest(http.MethodPost, path, createRequest)
	if err != nil {
		return nil, err
	}

	var accessTokenRoot AccessTokenRoot
	if _, err := ats.client.Do(ctx, req, &accessTokenRoot); err != nil {
		return nil, err
	}

	return accessTokenRoot.AccessToken, nil
}

// Delete access token
func (ats *AccessTokensService) Delete(ctx context.Context, userID string, accessTokenID string) error {
	path := fmt.Sprintf("api/internal/users/%s/access_tokens/%s", userID, accessTokenID)
	req, err := ats.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	_, err = ats.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}
