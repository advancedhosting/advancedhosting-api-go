package ah

import (
	"context"
)

// VolumeProductTariff object
type VolumeProductTariff struct {
	Component    string  `json:"component,omitempty"`
	TariffPrice  float64 `json:"tariff_price,omitempty"`
	MonthlyHours int     `json:"monthly_hours,omitempty"`
	Measure      string  `json:"measure,omitempty"`
}

// VolumeProduct object
type VolumeProduct struct {
	ID            string                `json:"id,omitempty"`
	CreatedAt     string                `json:"created_at,omitempty"`
	UpdatedAt     string                `json:"updated_at,omitempty"`
	Name          string                `json:"name,omitempty"`
	Type          string                `json:"type,omitempty"`
	Price         string                `json:"price,omitempty"`
	Currency      string                `json:"currency,omitempty"`
	Hot           bool                  `json:"hot,omitempty"`
	Tariff        []VolumeProductTariff `json:"tariff,omitempty"`
	MinSize       int                   `json:"min_size,omitempty"`
	MaxSize       int                   `json:"max_size,omitempty"`
	DatacenterIDs []string              `json:"datacenter_ids,omitempty"`
	Slug          string                `json:"slug,omitempty"`
	Category      *struct {
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
	Products []VolumeProduct `json:"products"`
	Meta     *Meta           `json:"meta,omitempty"`
}

// List returns all available volume products
func (vps *VolumeProductsService) List(ctx context.Context, options *ListOptions) ([]VolumeProduct, *Meta, error) {

	eqTypeFilter := &EqFilter{
		Keys:  []string{"type"},
		Value: "VolumeProduct",
	}

	if options == nil || options.Filters == nil {
		options = &ListOptions{
			Filters: []FilterInterface{eqTypeFilter},
		}
	} else {
		options.Filters = append(options.Filters, eqTypeFilter)
	}

	path := "api/v1/products"

	var pRoot productsRoot

	if err := vps.client.list(ctx, path, options, &pRoot); err != nil {
		return nil, nil, err
	}
	return pRoot.Products, pRoot.Meta, nil
}
