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
)

type InstancePlanAttributes struct {
	RAM              string `json:"ram,omitempty"`
	Disk             string `json:"disk,omitempty"`
	Slug             string `json:"slug,omitempty"`
	Vcpu             string `json:"vcpu,omitempty"`
	Traffic          string `json:"traffic,omitempty"`
	WebsaProductId   string `json:"websaProductId,omitempty"`
	Hot              bool   `json:"hot,omitempty"`
	AvailableOnTrial bool   `json:"available_on_trial,omitempty"`
}

type InstancePlan struct {
	CustomAttributes *InstancePlanAttributes `json:"custom_attributes,omitempty"`
	Plan
}

// InstancePlansAPI is an interface for instance plans.
type InstancePlansAPI interface {
	List(context.Context) ([]InstancePlan, error)
}

// InstancePlansService implements InstancePlansAPI interface.
type InstancePlansService struct {
	client *APIClient
}

type instancePlansRoot struct {
	Plans []InstancePlan `json:"data"`
}

// List returns all available instance plans
func (ips *InstancePlansService) List(ctx context.Context) ([]InstancePlan, error) {

	path := "api/v1/plans/public?type=vps"

	var ipRoot instancePlansRoot

	if err := ips.client.list(ctx, path, nil, &ipRoot); err != nil {
		return nil, err
	}
	return ipRoot.Plans, nil
}
