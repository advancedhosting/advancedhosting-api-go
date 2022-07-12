/*
Copyright 2022 Advanced Hosting

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

type PrivateClusterNetwork struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PrivateClusterNode struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	Cpu              int    `json:"cpu"`
	Vcpu             int    `json:"vcpu"`
	Disk             int    `json:"disk"`
	CpuName          string `json:"cpu_name"`
	CpuFrequency     string `json:"cpu_frequency"`
	Ram              int    `json:"ram"`
	VcpuAvailable    int    `json:"vcpu_available"`
	DiskAvailable    int    `json:"disk_available"`
	RamAvailable     int    `json:"ram_available"`
	CpuType          string `json:"cpu_type"`
	InstancesCount   int    `json:"instances_count"`
	IpAddressesCount int    `json:"ip_addresses_count"`
}

// PrivateCluster object
type PrivateCluster struct {
	Id               string                  `json:"id"`
	Name             string                  `json:"name"`
	IpAddressesCount int                     `json:"ip_addresses_count"`
	DatacenterId     string                  `json:"datacenter_id"`
	Networks         []PrivateClusterNetwork `json:"networks"`
	Nodes            []PrivateClusterNode    `json:"nodes"`
}

// PrivateClustersAPI is an interface for private clusters.
type PrivateClustersAPI interface {
	List(context.Context) ([]PrivateCluster, error)
}

// PrivateClustersService implements PrivateClustersAPI interface.
type PrivateClustersService struct {
	client *APIClient
}

type privateClustersRoot struct {
	PrivateClusters []PrivateCluster `json:"private_clusters,omitempty"`
}

// List returns all available private clusters
func (pns *PrivateClustersService) List(ctx context.Context) ([]PrivateCluster, error) {
	path := "api/v1/clusters/private"

	var pcsRoot privateClustersRoot

	if err := pns.client.list(ctx, path, nil, &pcsRoot); err != nil {
		return nil, err
	}

	return pcsRoot.PrivateClusters, nil
}
