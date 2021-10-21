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

// VolumeProductTariff object
type VolumeProductTariff struct {
	Component    string  `json:"component,omitempty"`
	Measure      string  `json:"measure,omitempty"`
	TariffPrice  float64 `json:"tariff_price,omitempty"`
	MonthlyHours int     `json:"monthly_hours,omitempty"`
}

// VolumeProduct object
type VolumeProduct struct {
	Category *struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"category,omitempty"`
	VolumeType *struct {
		ID               string `json:"id,omitempty"`
		Name             string `json:"name,omitempty"`
		Description      string `json:"description,omitempty"`
		DiskType         string `json:"disk_type,omitempty"`
		ReplicationLevel int    `json:"replication_level,omitempty"`
	}
	ID            string                `json:"id,omitempty"`
	CreatedAt     string                `json:"created_at,omitempty"`
	UpdatedAt     string                `json:"updated_at,omitempty"`
	Name          string                `json:"name,omitempty"`
	Type          string                `json:"type,omitempty"`
	Price         string                `json:"price,omitempty"`
	Currency      string                `json:"currency,omitempty"`
	Slug          string                `json:"slug,omitempty"`
	DatacenterIDs []string              `json:"datacenter_ids,omitempty"`
	Tariff        []VolumeProductTariff `json:"tariff,omitempty"`
	Hot           bool                  `json:"hot,omitempty"`
	MinSize       int                   `json:"min_size,omitempty"`
	MaxSize       int                   `json:"max_size,omitempty"`
}

// VolumeProductsAPI is an interface for volume products.
type VolumeProductsAPI interface {
	List(context.Context, *ListOptions) ([]VolumeProduct, *Meta, error)
}

// VolumeProductsService implements VolumeProductsAPI interface.
type VolumeProductsService struct {
	client *APIClient
}

type productsRoot struct {
	Meta     *Meta           `json:"meta,omitempty"`
	Products []VolumeProduct `json:"products"`
}

// List returns all available volume products
func (vps *VolumeProductsService) List(ctx context.Context, options *ListOptions) ([]VolumeProduct, *Meta, error) {

	path := "api/v1/products/volumes"

	var pRoot productsRoot

	if err := vps.client.list(ctx, path, options, &pRoot); err != nil {
		return nil, nil, err
	}
	return pRoot.Products, pRoot.Meta, nil
}
