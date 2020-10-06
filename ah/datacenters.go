package ah

import (
	"context"
	"fmt"
	"net/http"
)

// DatacenterRegion object
type DatacenterRegion struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// Datacenter object
type Datacenter struct {
	ID                string            `json:"id,omitempty"`
	Name              string            `json:"name,omitempty"`
	FullName          string            `json:"full_name,omitempty"`
	Slug              string            `json:"slug,omitempty"`
	InstancesRunning  int               `json:"instances_running,omitempty"`
	PrivateNodesCount int               `json:"private_nodes_count,omitempty"`
	Region            *DatacenterRegion `json:"region,omitempty"`
}

// DatacentersAPI is an interface for datacenters.
type DatacentersAPI interface {
	List(context.Context, *ListOptions) ([]Datacenter, error)
	Get(context.Context, string) (*Datacenter, error)
}

// DatacentersService implements DatacentersAPI interface.
type DatacentersService struct {
	client *APIClient
}

type datacentersRoot struct {
	Datacenters []Datacenter `json:"datacenters"`
}

// List returns all available datacenters
func (ds *DatacentersService) List(ctx context.Context, options *ListOptions) ([]Datacenter, error) {

	path := "api/v1/datacenters"

	var dRoot datacentersRoot

	if err := ds.client.list(ctx, path, options, &dRoot); err != nil {
		return nil, err
	}
	return dRoot.Datacenters, nil
}

type datacenterRoot struct {
	Datacenter *Datacenter `json:"datacenter"`
}

// Get datacenter info by ID
func (ds *DatacentersService) Get(ctx context.Context, datacenterID string) (*Datacenter, error) {

	path := fmt.Sprintf("api/v1/datacenters/%s", datacenterID)
	req, err := ds.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var dRoot datacenterRoot
	_, err = ds.client.Do(ctx, req, &dRoot)

	if err != nil {
		return nil, err
	}

	return dRoot.Datacenter, nil
}
