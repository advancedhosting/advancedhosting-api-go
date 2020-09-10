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

// IPAddressAssignment object
type IPAddressAssignment struct {
	ID          string `json:"id,omitempty"`
	InstanceID  string `json:"instance_id,omitempty"`
	IPAddressID string `json:"ip_address_id,omitempty"`
	State       string `json:"state,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// IPAddressAssignmentsAPI is an interface for ip address assignments.
type IPAddressAssignmentsAPI interface {
	Create(context.Context, *IPAddressAssignmentCreateRequest) (*IPAddressAssignment, error)
	Delete(context.Context, string) error
}

// IPAddressAssignmentsService implements IPAddressAssignmentsAPI interface.
type IPAddressAssignmentsService struct {
	client *APIClient
}

// IPAddressAssignmentCreateRequest represents a request to assign an ip address to isntance.
type IPAddressAssignmentCreateRequest struct {
	IPAddressID string `json:"ip_address_id"`
	InstanceID  string `json:"instance_id"`
}

type ipAddressAssignmentRoot struct {
	InstanceIPAddress *IPAddressAssignment `json:"instance_ip_address"`
}

// Create ip address assignment
func (ips *IPAddressAssignmentsService) Create(ctx context.Context, createRequest *IPAddressAssignmentCreateRequest) (*IPAddressAssignment, error) {

	type request struct {
		InstanceIPAddress *IPAddressAssignmentCreateRequest `json:"instance_ip_address"`
	}

	path := "api/v1/instance_ip_addresses"
	req, err := ips.client.newRequest(http.MethodPost, path, &request{createRequest})

	if err != nil {
		return nil, err
	}

	var ipRoot ipAddressAssignmentRoot
	if _, err := ips.client.Do(ctx, req, &ipRoot); err != nil {
		return nil, err
	}

	return ipRoot.InstanceIPAddress, nil
}

// Get ip address assignment
// TODO Implement after WCS-3500
func (ips *IPAddressAssignmentsService) Get(ctx context.Context) {
}

// List ip address assignments
// TODO Implement after WCS-3500
func (ips *IPAddressAssignmentsService) List(ctx context.Context) {
}

// Delete assignment
func (ips *IPAddressAssignmentsService) Delete(ctx context.Context, isntanceIPAssignmentID string) error {
	path := fmt.Sprintf("api/v1/instance_ip_addresses/%s", isntanceIPAssignmentID)
	req, err := ips.client.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return err
	}

	if _, err := ips.client.Do(ctx, req, nil); err != nil {
		return err
	}

	return nil
}
