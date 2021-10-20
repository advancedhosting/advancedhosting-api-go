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
)

type VolumePlanAttributes struct {
	VolumeType *struct {
		ID               string `json:"id,omitempty"`
		DiskType         string `json:"disk_type,omitempty"`
		ReplicationLevel int    `json:"replication_level,omitempty"`
	} `json:"volume_type,omitempty"`

	WebsaProductId  string   `json:"websaProductId,omitempty"`
	Slug            string   `json:"slug,omitempty"`
	DatacenterIds   []string `json:"datacenter_ids,omitempty"`
	PredefinedSizes []struct {
		Name string `json:"name,omitempty"`
		Size int    `json:"size,omitempty"`
	} `json:"predefined_sizes,omitempty"`
	Hot     bool `json:"hot,omitempty"`
	MaxSize int  `json:"max_size,omitempty"`
	MinSize int  `json:"min_size,omitempty"`
}

type VolumePlan struct {
	CustomAttributes *VolumePlanAttributes `json:"custom_attributes,omitempty"`
	Plan
}

// VolumePlansAPI is an interface for volume plans.
type VolumePlansAPI interface {
	List(context.Context) ([]VolumePlan, error)
}

// VolumePlansService implements VolumePlansAPI interface.
type VolumePlansService struct {
	client *APIClient
}

type volumePlansRoot struct {
	Plans []VolumePlan `json:"data"`
}

// List returns all available volume plans
func (vp *VolumePlansService) List(ctx context.Context) ([]VolumePlan, error) {

	path := "api/v1/plans/public?type=volume"

	var ipRoot volumePlansRoot

	if err := vp.client.list(ctx, path, nil, &ipRoot); err != nil {
		return nil, err
	}
	return ipRoot.Plans, nil
}
