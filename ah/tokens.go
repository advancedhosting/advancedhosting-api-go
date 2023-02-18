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
	"fmt"
	"net/http"
)

// Token object
type Token struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Token     string     `json:"token,omitempty"`
	Scopes    [88]string `json:"scopes,omitempty"`
	ExpiresIn string     `json:"expires_in,omitempty"`
	CreatedAt string     `json:"created_at,omitempty"`
}

// TokenCreateRequest object
type TokenCreateRequest struct {
	Name string `json:"name,omitempty"`
}

// TokensAPI is an interface for tokens.
type TokensAPI interface {
	List(context.Context, *ListOptions) ([]Token, error)
	Get(context.Context, string) (*Token, error)
	Create(context.Context, *TokenCreateRequest) (*Token, error)
	Delete(context.Context, string) error
}

// TokensService implements TokensAPI interface.
type TokensService struct {
	client *APIClient
}

// tokensRoot is a root object for tokens
type tokensRoot struct {
	Token  *Token  `json:"token,omitempty"`
	Tokens []Token `json:"tokens,omitempty"`
}

// Get returns a token by ID
func (s *TokensService) Get(ctx context.Context, tokenId string) (*Token, error) {
	path := fmt.Sprintf("id/api/v1/access_tokens/%s", tokenId)
	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var tokensRoot tokensRoot
	_, err = s.client.Do(ctx, req, &tokensRoot)
	if err != nil {
		return nil, err
	}

	return tokensRoot.Token, nil
}

// List returns all available tokens
func (s *TokensService) List(ctx context.Context, options *ListOptions) ([]Token, error) {
	path := "id/api/v1/access_tokens"

	var tokensRoot tokensRoot
	if err := s.client.list(ctx, path, options, &tokensRoot); err != nil {
		return nil, err
	}

	return tokensRoot.Tokens, nil
}

// Create creates a new token
func (s *TokensService) Create(ctx context.Context, request *TokenCreateRequest) (*Token, error) {
	path := "id/api/v1/access_tokens"
	req, err := s.client.newRequest(http.MethodPost, path, request)

	if err != nil {
		return nil, err
	}

	var tokensRoot tokensRoot
	if _, err := s.client.Do(ctx, req, &tokensRoot); err != nil {
		return nil, err
	}

	return tokensRoot.Token, nil
}

// Delete deletes a token by ID
func (s *TokensService) Delete(ctx context.Context, tokenId string) error {
	path := fmt.Sprintf("id/api/v1/access_tokens/%s", tokenId)
	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	_, err = s.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}
