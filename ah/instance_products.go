package ah

import (
	"context"
)

// InstanceProductTariff object
type InstanceProductTariff struct {
	Component              string  `json:"component,omitempty"`
	TariffPrice            float64 `json:"tariff_price,omitempty"`
	AdditionalServicePrice int     `json:"additional_service_price,omitempty"`
	IncludedValue          int     `json:"included_value,omitempty"`
	Measure                string  `json:"measure,omitempty"`
}

// InstanceProduct object
type InstanceProduct struct {
	ID               string                  `json:"id,omitempty"`
	CreatedAt        string                  `json:"created_at,omitempty"`
	UpdatedAt        string                  `json:"updated_at,omitempty"`
	Name             string                  `json:"name,omitempty"`
	Type             string                  `json:"type,omitempty"`
	Price            string                  `json:"price,omitempty"`
	Currency         string                  `json:"currency,omitempty"`
	Hot              bool                    `json:"hot,omitempty"`
	Tariff           []InstanceProductTariff `json:"tariff,omitempty"`
	Vcpu             string                  `json:"vcpu,omitempty"`
	RAM              string                  `json:"ram,omitempty"`
	Disk             string                  `json:"disk,omitempty"`
	Traffic          string                  `json:"traffic,omitempty"`
	AvailableOnTrial bool                    `json:"available_on_trial,omitempty"`
	Slug             string                  `json:"slug,omitempty"`
	Category         *struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"category,omitempty"`
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
	Products []InstanceProduct `json:"products"`
	Meta     *Meta             `json:"meta,omitempty"`
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
