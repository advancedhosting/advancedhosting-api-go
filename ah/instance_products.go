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

// InstanceProductTariff object
type InstanceProductTariff struct {
	Component              string  `json:"component,omitempty"`
	Measure                string  `json:"measure,omitempty"`
	TariffPrice            float64 `json:"tariff_price,omitempty"`
	AdditionalServicePrice int     `json:"additional_service_price,omitempty"`
	IncludedValue          int     `json:"included_value,omitempty"`
}

// InstanceProduct object
type InstanceProduct struct {
	Category *struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"category,omitempty"`
	ID               string                  `json:"id,omitempty"`
	CreatedAt        string                  `json:"created_at,omitempty"`
	UpdatedAt        string                  `json:"updated_at,omitempty"`
	Name             string                  `json:"name,omitempty"`
	Type             string                  `json:"type,omitempty"`
	Price            string                  `json:"price,omitempty"`
	Currency         string                  `json:"currency,omitempty"`
	Vcpu             string                  `json:"vcpu,omitempty"`
	RAM              string                  `json:"ram,omitempty"`
	Disk             string                  `json:"disk,omitempty"`
	Traffic          string                  `json:"traffic,omitempty"`
	Slug             string                  `json:"slug,omitempty"`
	Tariff           []InstanceProductTariff `json:"tariff,omitempty"`
	Hot              bool                    `json:"hot,omitempty"`
	AvailableOnTrial bool                    `json:"available_on_trial,omitempty"`
}

// InstanceProductsAPI is an interface for instance products.
type InstanceProductsAPI interface {
	List(context.Context, *ListOptions) ([]InstanceProduct, *Meta, error)
}

// InstanceProductsService implements InstanceProductsAPI interface.
type InstanceProductsService struct {
	client *APIClient
}

type instanceProductsRoot struct {
	Meta     *Meta             `json:"meta,omitempty"`
	Products []InstanceProduct `json:"products"`
}

// List returns all available volume products
func (ips *InstanceProductsService) List(ctx context.Context, options *ListOptions) ([]InstanceProduct, *Meta, error) {

	path := "api/v1/products/instances"

	var ipRoot instanceProductsRoot

	if err := ips.client.list(ctx, path, options, &ipRoot); err != nil {
		return nil, nil, err
	}
	return ipRoot.Products, ipRoot.Meta, nil
}
