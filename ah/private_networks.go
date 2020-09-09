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
	"fmt"
	"net/http"
)

// PrivateNetwork object
type PrivateNetwork struct {
	ID             string `json:"id,omitempty"`
	Number         string `json:"number,omitempty"`
	CIDR           string `json:"cidr,omitempty"`
	Name           string `json:"name,omitempty"`
	State          string `json:"state,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
	InstancesCount int    `json:"instances_count,omitempty"`
}

// PrivateNetworksAPI is an interface for private networks.
type PrivateNetworksAPI interface {
	List(context.Context, *ListOptions) ([]PrivateNetwork, error)
	Get(context.Context, string) (*PrivateNetworkInfo, error)
	Create(context.Context, *PrivateNetworkCreateRequest) (*PrivateNetworkInfo, error)
}

// PrivateNetworksService implements PrivateNetworksAPI interface.
type PrivateNetworksService struct {
	client *APIClient
}

type privateNetworksRoot struct {
	PrivateNetworks []PrivateNetwork `json:"private_networks,omitempty"`
}

// List returns all available private networks
func (pns *PrivateNetworksService) List(ctx context.Context, options *ListOptions) ([]PrivateNetwork, error) {
	path := "api/v1/private_networks"

	var pnsRoot privateNetworksRoot

	if err := pns.client.list(ctx, path, options, &pnsRoot); err != nil {
		return nil, err
	}

	return pnsRoot.PrivateNetworks, nil
}

type privateNetworkInfoRoot struct {
	PrivateNetwork *PrivateNetworkInfo `json:"private_network,omitempty"`
}

// Get private network
func (pns *PrivateNetworksService) Get(ctx context.Context, privateNetworkID string) (*PrivateNetworkInfo, error) {
	path := fmt.Sprintf("api/v1/private_networks/%s", privateNetworkID)
	req, err := pns.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var pnInfo privateNetworkInfoRoot
	_, err = pns.client.Do(ctx, req, &pnInfo)

	if err != nil {
		return nil, err
	}

	return pnInfo.PrivateNetwork, err
}

// InstancePrivateNetworkAttributes object
type InstancePrivateNetworkAttributes struct {
	InstanceID string `json:"instance_id,omitempty"`
	IP         string `json:"ip,omitempty"`
}

// PrivateNetworkCreateRequest object
type PrivateNetworkCreateRequest struct {
	Name                             string                             `json:"name"`
	CIDR                             string                             `json:"cidr"`
	InstancePrivateNetworkAttributes []InstancePrivateNetworkAttributes `json:"instance_private_networks_attributes,omitempty"`
}

// Create private network
func (pns *PrivateNetworksService) Create(ctx context.Context, createRequest *PrivateNetworkCreateRequest) (*PrivateNetworkInfo, error) {

	type request struct {
		PrivateNetwork *PrivateNetworkCreateRequest `json:"private_network"`
	}
	req, err := pns.client.newRequest(http.MethodPost, "api/v1/private_networks", &request{createRequest})
	if err != nil {
		return nil, err
	}

	var pnInfo privateNetworkInfoRoot
	if _, err := pns.client.Do(ctx, req, &pnInfo); err != nil {
		return nil, err
	}

	return pnInfo.PrivateNetwork, nil

}
