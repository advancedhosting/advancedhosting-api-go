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

// IPAddress object
type IPAddress struct {
	Address                      string   `json:"address,omitempty"`
	Type                         string   `json:"address_type,omitempty"`
	CreatedAt                    string   `json:"created_at,omitempty"`
	DatacenterFullName           string   `json:"datacenter_full_name,omitempty"`
	ID                           string   `json:"id,omitempty"`
	ReverseDNS                   string   `json:"reverse_dns,omitempty"`
	UpdatedAt                    string   `json:"updated_at,omitempty"`
	InstanceIDs                  []string `json:"instance_ids,omitempty"`
	DeleteProtection             bool     `json:"delete_protection,omitempty"`
	NetworkUsedForPrivateCluster bool     `json:"network_used_for_private_cluster,omitempty"`
}

type ipAddressesRoot struct {
	IPAddresses []IPAddress `json:"ip_addresses"`
}

type ipAddressRoot struct {
	IPAddress *IPAddress `json:"ip_address"`
}

// IPAddressesAPI is an interface for ip addresses.
type IPAddressesAPI interface {
	List(context.Context, *ListOptions) ([]IPAddress, error)
	Create(context.Context, *IPAddressCreateRequest) (*IPAddress, error)
	Get(context.Context, string) (*IPAddress, error)
	Delete(context.Context, string) error
	Update(context.Context, string, *IPAddressUpdateRequest) (*IPAddress, error)
}

// IPAddressesService implements IPAddressesAPI interface.
type IPAddressesService struct {
	client *APIClient
}

// List returns all available ip addresses
func (ips *IPAddressesService) List(ctx context.Context, options *ListOptions) ([]IPAddress, error) {
	path := "api/v1/ip_addresses"
	var ipsRoot ipAddressesRoot

	if err := ips.client.list(ctx, path, options, &ipsRoot); err != nil {
		return nil, err
	}
	return ipsRoot.IPAddresses, nil
}

// IPAddressCreateRequest represents a request to create an ip address.
type IPAddressCreateRequest struct {
	Type             string   `json:"address_type"`
	DatacenterID     string   `json:"datacenter_id,omitempty"`
	DatacenterSlug   string   `json:"datacenter_slug,omitempty"`
	ReverseDNS       string   `json:"reverse_dns,omitempty"`
	InstanceIDs      []string `json:"instance_ids,omitempty"`
	DeleteProtection bool     `json:"delete_protection,omitempty"`
}

// Create ip address
func (ips *IPAddressesService) Create(ctx context.Context, createRequest *IPAddressCreateRequest) (*IPAddress, error) {

	type request struct {
		IPAddress *IPAddressCreateRequest `json:"ip_address"`
	}
	req, err := ips.client.newRequest(http.MethodPost, "api/v1/ip_addresses", &request{createRequest})
	if err != nil {
		return nil, err
	}

	var ipRoot ipAddressRoot
	if _, err := ips.client.Do(ctx, req, &ipRoot); err != nil {
		return nil, err
	}

	return ipRoot.IPAddress, nil

}

// Get ip address
func (ips *IPAddressesService) Get(ctx context.Context, ipAddressID string) (*IPAddress, error) {
	options := &ListOptions{
		Filters: []FilterInterface{
			&EqFilter{
				Keys:  []string{"id"},
				Value: ipAddressID,
			},
		},
	}

	ipAddresses, err := ips.List(ctx, options)
	if err != nil {
		return nil, err
	}

	if len(ipAddresses) != 1 {
		return nil, ErrResourceNotFound
	}

	return &ipAddresses[0], nil
}

// IPAddressUpdateRequest represents a request to update an ip address resource.
type IPAddressUpdateRequest struct {
	ReverseDNS       string `json:"reverse_dns,omitempty"`
	DeleteProtection bool   `json:"delete_protection,omitempty"`
}

// Update ip address resource
func (ips *IPAddressesService) Update(ctx context.Context, ipAddressID string, request *IPAddressUpdateRequest) (*IPAddress, error) {
	path := fmt.Sprintf("api/v1/ip_addresses/%s", ipAddressID)
	req, err := ips.client.newRequest(http.MethodPatch, path, request)

	if err != nil {
		return nil, err
	}

	var ipRoot ipAddressRoot
	if _, err := ips.client.Do(ctx, req, &ipRoot); err != nil {
		return nil, err
	}

	return ipRoot.IPAddress, nil
}

// Delete ip address
func (ips *IPAddressesService) Delete(ctx context.Context, ipAddressID string) error {
	path := fmt.Sprintf("api/v1/ip_addresses/%s", ipAddressID)
	req, err := ips.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	_, err = ips.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}
