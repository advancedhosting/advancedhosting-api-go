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

// InstancePrivateNetworkInfo object
type InstancePrivateNetworkInfo struct {
	ID          string `json:"id,omitempty"`
	IP          string `json:"ip"`
	MACAddress  string `json:"mac_address"`
	State       string `json:"state,omitempty"`
	ConnectedAt string `json:"connected_at,omitempty"`
	Instance    *struct {
		ID      string `json:"id,omitempty"`
		ImageID string `json:"image_id"`
		Name    string `json:"name"`
		Number  string `json:"number,omitempty"`
	} `json:"instance,omitempty"`
}

// InstancePrivateNetwork object
type InstancePrivateNetwork struct {
	InstancePrivateNetworkInfo
	PrivateNetwork *PrivateNetwork `json:"private_network,omitempty"`
}

// InstancePrivateNetworksAPI is an interface for instance connection to private network.
type InstancePrivateNetworksAPI interface {
	Create(context.Context, *InstancePrivateNetworkCreateRequest) (*InstancePrivateNetwork, error)
	Get(context.Context, string) (*InstancePrivateNetwork, error)
	Update(context.Context, string, *InstancePrivateNetworkUpdateRequest) (*InstancePrivateNetwork, error)
	Delete(context.Context, string) (*InstancePrivateNetwork, error)
}

// InstancePrivateNetworksService implements InstancePrivateNetworkConnectionsAPI interface.
type InstancePrivateNetworksService struct {
	client *APIClient
}

// InstancePrivateNetworkCreateRequest object
type InstancePrivateNetworkCreateRequest struct {
	PrivateNetworkID string `json:"private_network_id"`
	InstanceID       string `json:"instance_id"`
	IP               string `json:"ip,omitempty"`
}

type instancePrivateNetworkInfoRoot struct {
	InstancePrivateNetwork *InstancePrivateNetwork `json:"instance_private_network"`
}

// Get instance private network info
func (ipns *InstancePrivateNetworksService) Get(ctx context.Context, instancePrivateNetworkID string) (*InstancePrivateNetwork, error) {

	path := fmt.Sprintf("api/v1/instance_private_networks/%s", instancePrivateNetworkID)

	req, err := ipns.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var ipnInfo instancePrivateNetworkInfoRoot
	if _, err := ipns.client.Do(ctx, req, &ipnInfo); err != nil {
		return nil, err
	}

	return ipnInfo.InstancePrivateNetwork, nil
}

// Create instance connection to the private network
func (ipns *InstancePrivateNetworksService) Create(
	ctx context.Context,
	addRequest *InstancePrivateNetworkCreateRequest) (*InstancePrivateNetwork, error) {

	type request struct {
		PrivateNetwork *InstancePrivateNetworkCreateRequest `json:"instance_private_network"`
	}
	req, err := ipns.client.newRequest(http.MethodPost, "api/v1/instance_private_networks", &request{addRequest})
	if err != nil {
		return nil, err
	}

	var ipnInfo instancePrivateNetworkInfoRoot
	if _, err := ipns.client.Do(ctx, req, &ipnInfo); err != nil {
		return nil, err
	}

	return ipnInfo.InstancePrivateNetwork, nil
}

// InstancePrivateNetworkUpdateRequest object
type InstancePrivateNetworkUpdateRequest struct {
	IP string `json:"ip"`
}

// Update instance connection to private network
func (ipns *InstancePrivateNetworksService) Update(
	ctx context.Context,
	instancePrivateNetworkID string,
	updateRequest *InstancePrivateNetworkUpdateRequest) (*InstancePrivateNetwork, error) {

	type request struct {
		InstancePrivateNetwork *InstancePrivateNetworkUpdateRequest `json:"instance_private_network"`
	}

	path := fmt.Sprintf("api/v1/instance_private_networks/%s", instancePrivateNetworkID)
	req, err := ipns.client.newRequest(http.MethodPatch, path, &request{updateRequest})

	if err != nil {
		return nil, err
	}

	var ipnInfo instancePrivateNetworkInfoRoot
	if _, err := ipns.client.Do(ctx, req, &ipnInfo); err != nil {
		return nil, err
	}

	return ipnInfo.InstancePrivateNetwork, nil
}

// Delete disconnects instance from the private network
func (ipns *InstancePrivateNetworksService) Delete(ctx context.Context, instancePrivateNetworkID string) (*InstancePrivateNetwork, error) {
	path := fmt.Sprintf("api/v1/instance_private_networks/%s", instancePrivateNetworkID)
	req, err := ipns.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	var ipnInfo instancePrivateNetworkInfoRoot
	if _, err := ipns.client.Do(ctx, req, &ipnInfo); err != nil {
		return nil, err
	}

	return ipnInfo.InstancePrivateNetwork, nil
}
