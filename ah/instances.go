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
	"errors"
	"fmt"
	"net/http"
)

var (
	// ErrPrimaryIPNotFound is returned when primary is not found
	ErrPrimaryIPNotFound = errors.New("primary ip is not found")
)

// InstanceShutDownStatus represents instance's shutdown statuse
const InstanceShutDownStatus = "stopped"

// InstanceRegion object
type InstanceRegion struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Slug         string   `json:"slug,omitempty"`
	CountryCode  string   `json:"country_code,omitempty"`
	Country      string   `json:"country,omitempty"`
	City         string   `json:"city,omitempty"`
	ParentID     string   `json:"parent_id,omitempty"`
	Group        bool     `json:"group,omitempty"`
	RegionsCount int      `json:"regions_count,omitempty"`
	Services     []string `json:"services,omitempty"`
}

// InstanceDatacenterRegion object
type InstanceDatacenterRegion struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// InstanceDatacenter object
type InstanceDatacenter struct {
	ID                string                    `json:"id,omitempty"`
	Name              string                    `json:"name,omitempty"`
	FullName          string                    `json:"full_name,omitempty"`
	InstancesRunning  int                       `json:"instances_running,omitempty"`
	PrivateNodesCount int                       `json:"private_nodes_count,omitempty"`
	Region            *InstanceDatacenterRegion `json:"region,omitempty"`
}

// InstanceV4Network object
type InstanceV4Network struct {
	Type      string `json:"type,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
}

// InstanceNetworks object
type InstanceNetworks struct {
	V4 []InstanceV4Network `json:"v4,omitempty"`
}

// Action object
type Action struct {
	ID           string `json:"id,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
	State        string `json:"state,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Type         string `json:"type,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	Note         string `json:"note,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	StartedAt    string `json:"started_at,omitempty"`
	CompletedAt  string `json:"completed_at,omitempty"`
}

// InstanceVolume object
type InstanceVolume struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Size       int    `json:"size,omitempty"`
	Number     string `json:"number,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
	Port       int    `json:"port,omitempty"`
	FileSystem string `json:"file_system,omitempty"`
	State      string `json:"state,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	AttachedAt string `json:"attached_at,omitempty"`
}

// InstancePrivateNetwork object
type InstancePrivateNetwork struct {
	ID             string `json:"id,omitempty"`
	IP             string `json:"ip,omitempty"`
	MacAddress     string `json:"mac_address,omitempty"`
	State          string `json:"state,omitempty"`
	ConnectedAt    string `json:"connected_at,omitempty"`
	PrivateNetwork struct {
		ID string `json:"id,omitempty"`
	} `json:"private_network,omitempty"`
}

// InstanceImage object
type InstanceImage struct {
	ID           string `json:"id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	Type         string `json:"type,omitempty"`
	Name         string `json:"name,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Public       bool   `json:"public,omitempty"`
}

// InstanceSSHKey object
type InstanceSSHKey struct {
	CreatedAt   string `json:"created_at,omitempty"`
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
}

// Instance object
type Instance struct {
	ID                         string                   `json:"id,omitempty"`
	CreatedAt                  string                   `json:"created_at,omitempty"`
	UpdatedAt                  string                   `json:"updated_at,omitempty"`
	Number                     string                   `json:"number,omitempty"`
	Name                       string                   `json:"name,omitempty"`
	State                      string                   `json:"state,omitempty"`
	Disk                       int                      `json:"disk,omitempty"`
	StateDescription           string                   `json:"state_description,omitempty"`
	Locked                     bool                     `json:"locked,omitempty"`
	UseSSHPassword             bool                     `json:"use_ssh_password,omitempty"`
	SSHKeys                    []InstanceSSHKey         `json:"ssh_keys,omitempty"`
	ProductID                  string                   `json:"product_id,omitempty"`
	Vcpu                       int                      `json:"vcpu,omitempty"`
	RAM                        int                      `json:"ram,omitempty"`
	Traffic                    int                      `json:"traffic,omitempty"`
	Tags                       []string                 `json:"tags,omitempty"`
	PrimaryInstanceIPAddressID string                   `json:"primary_instance_ip_address_id,omitempty"`
	IPScheme                   string                   `json:"ip_scheme,omitempty"`
	Region                     *InstanceRegion          `json:"region,omitempty"`
	Datacenter                 *InstanceDatacenter      `json:"datacenter,omitempty"`
	Features                   []string                 `json:"features,omitempty"`
	Networks                   *InstanceNetworks        `json:"networks,omitempty"`
	CurrentAction              *Action                  `json:"current_action,omitempty"`
	LastAction                 *Action                  `json:"last_action,omitempty"`
	Reason                     string                   `json:"reason,omitempty"`
	SnapshotBySchedule         bool                     `json:"snapshot_by_schedule,omitempty"`
	SnapshotPeriod             string                   `json:"snapshot_period,omitempty"`
	MaxVolumesNumber           int                      `json:"max_volumes_number,omitempty"`
	PrivateNetworks            []InstancePrivateNetwork `json:"instance_private_networks,omitempty"`
	IPAddresses                []InstanceIPAddress      `json:"instance_ip_addresses,omitempty"`
	Image                      *InstanceImage           `json:"image,omitempty"`
	Volumes                    []InstanceVolume         `json:"volumes,omitempty"`
}

// PrimaryIPAddr returns primary IP object of the instance
func (i *Instance) PrimaryIPAddr() (*InstanceIPAddress, error) {
	for _, ipAddress := range i.IPAddresses {
		if ipAddress.ID == i.PrimaryInstanceIPAddressID {
			return &ipAddress, nil
		}
	}
	return nil, ErrPrimaryIPNotFound
}

// InstanceIPAddress object
type InstanceIPAddress struct {
	ID          string `json:"id,omitempty"`
	InstanceID  string `json:"instance_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	IPAddressID string `json:"ip_address_id,omitempty"`
	Address     string `json:"address,omitempty"`
}

type instancesRoot struct {
	Instances []Instance `json:"instances"`
	Meta      *Meta      `json:"meta"`
}

type instanceRoot struct {
	Instance *Instance `json:"instance"`
	Meta     *Meta     `json:"meta"`
}

// InstanceCreateRequest represents a request to create a instance.
type InstanceCreateRequest struct {
	Name                  string   `json:"name"`
	DatacenterID          string   `json:"datacenter_id"`
	ImageID               string   `json:"image_id"`
	ProductID             string   `json:"product_id"`
	UseSSHPassword        bool     `json:"use_ssh_password"`
	Tags                  []string `json:"tags"`
	SSHKeyIDs             []string `json:"ssh_key_ids"`
	CreatePublicIPAddress bool     `json:"create_public_ip_address"`
	SnapshotPeriod        string   `json:"snapshot_period"`
	SnapshotBySchedule    bool     `json:"snapshot_by_schedule"`
}

// InstanceRenameRequest represents a request to rename the instance.
type InstanceRenameRequest struct {
	Name string `json:"name"`
}

// InstanceActionRequest represents an action request.
type InstanceActionRequest struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// InstanceUpgradeRequest represents a request to upgrade the instance.
type InstanceUpgradeRequest struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	ProductID string `json:"product_id"`
}

// InstancesAPI is an interface for instances.
type InstancesAPI interface {
	List(context.Context, *ListOptions) ([]Instance, *Meta, error)
	Get(context.Context, string) (*Instance, error)
	Create(context.Context, *InstanceCreateRequest) (*Instance, error)
	Rename(context.Context, string, string) (*Instance, error)
	Upgrade(context.Context, string, string) error
	Shutdown(context.Context, string) error
	PowerOff(context.Context, string) error
	Destroy(context.Context, string) error
	SetPrimaryIP(context.Context, string, string) (*InstanceAction, error)
	ActionInfo(context.Context, string, string) (*InstanceAction, error)
	Actions(context.Context, string) ([]InstanceAction, error)
}

// InstancesService implements InstancesApi interface.
type InstancesService struct {
	client *APIClient
}

// Create new instance.
func (is *InstancesService) Create(ctx context.Context, createRequest *InstanceCreateRequest) (*Instance, error) {

	type request struct {
		Instance *InstanceCreateRequest `json:"instance"`
	}
	req, err := is.client.newRequest(http.MethodPost, "api/v1/instances", &request{createRequest})
	if err != nil {
		return nil, err
	}

	var instanceRoot instanceRoot
	resp, err := is.client.Do(ctx, req, &instanceRoot)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 202 {
		return nil, fmt.Errorf("Error creating instance: %v", resp.StatusCode)
	}
	return instanceRoot.Instance, err

}

// Rename instance.
func (is *InstancesService) Rename(ctx context.Context, instanceID, name string) (*Instance, error) {
	createRequest := &InstanceRenameRequest{
		Name: name,
	}
	path := fmt.Sprintf("api/v1/instances/%s", instanceID)
	req, err := is.client.newRequest(http.MethodPatch, path, createRequest)
	if err != nil {
		return nil, err
	}

	var instanceRoot instanceRoot
	resp, err := is.client.Do(ctx, req, &instanceRoot)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error renaming instance: %v", resp.StatusCode)
	}
	return instanceRoot.Instance, err
}

// Upgrade instance.
func (is *InstancesService) Upgrade(ctx context.Context, instanceID, productID string) error {
	upgradeRequest := &InstanceUpgradeRequest{
		ID:        instanceID,
		Type:      "upgrade",
		ProductID: productID,
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, upgradeRequest)
	if err != nil {
		return err
	}

	resp, err := is.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error upgrade instance: %v", resp.StatusCode)
	}
	return err
}

// Shutdown instance.
func (is *InstancesService) Shutdown(ctx context.Context, instanceID string) error {
	actionRequest := &InstanceActionRequest{
		ID:   instanceID,
		Type: "shutdown",
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, actionRequest)
	if err != nil {
		return err
	}

	resp, err := is.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error shutdown instance: %v", resp.StatusCode)
	}
	return err
}

// PowerOff instance.
func (is *InstancesService) PowerOff(ctx context.Context, instanceID string) error {
	actionRequest := &InstanceActionRequest{
		ID:   instanceID,
		Type: "power_off",
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, actionRequest)
	if err != nil {
		return err
	}

	resp, err := is.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error power_off instance: %v", resp.StatusCode)
	}
	return err
}

// InstanceDestroyRequest represents a request to destroy the instance.
type InstanceDestroyRequest struct {
	BackupsStrategy string `json:"backups_strategy"`
}

// Destroy isntance.
func (is *InstancesService) Destroy(ctx context.Context, instanceID string) error {
	destroyRequest := &InstanceDestroyRequest{
		BackupsStrategy: "destroy",
	}
	path := fmt.Sprintf("api/v1/instances/%s", instanceID)
	req, err := is.client.newRequest(http.MethodDelete, path, destroyRequest)
	if err != nil {
		return err
	}

	var a interface{}
	resp, err := is.client.Do(ctx, req, a)
	if err != nil {
		return err
	}

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error shutdown instance: %v", resp.StatusCode)
	}
	return err
}

// List returns all available instances
func (is *InstancesService) List(ctx context.Context, options *ListOptions) ([]Instance, *Meta, error) {
	path := "api/v1/instances"

	var instRoot instancesRoot

	if err := is.client.list(ctx, path, options, &instRoot); err != nil {
		return nil, nil, err
	}

	return instRoot.Instances, instRoot.Meta, nil

}

// Get returns all instance by instanceID
func (is *InstancesService) Get(ctx context.Context, instanceID string) (*Instance, error) {
	path := fmt.Sprintf("api/v1/instances/%s", instanceID)
	req, err := is.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var instanceRoot instanceRoot
	_, err = is.client.Do(ctx, req, &instanceRoot)

	if err != nil {
		return nil, err
	}

	return instanceRoot.Instance, err

}

type instanceSetPrimaryIPRequest struct {
	Type                string `json:"type"`
	InstanceIPAddressID string `json:"instance_ip_address_id"`
}

// SetPrimaryIP makes ip primary for instance
func (is *InstancesService) SetPrimaryIP(ctx context.Context, instanceID, ipAssignmentID string) (*InstanceAction, error) {
	request := &instanceSetPrimaryIPRequest{
		InstanceIPAddressID: ipAssignmentID,
		Type:                "set_primary_ip",
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var actionRoot instanceActionRoot
	_, err = is.client.Do(ctx, req, &actionRoot)
	if err != nil {
		return nil, err
	}

	return actionRoot.Action, nil
}

// InstanceAction object
type InstanceAction struct {
	ID           string `json:"id,omitempty"`
	State        string `json:"state,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Type         string `json:"type,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	Note         string `json:"note,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	StartedAt    string `json:"started_at,omitempty"`
	CompletedAt  string `json:"completed_at,omitempty"`
}

type instanceActionRoot struct {
	Action *InstanceAction `json:"action"`
}

type instanceActionsRoot struct {
	Actions []InstanceAction `json:"actions"`
}

// ActionInfo returns instance's action info by action ID
func (is *InstancesService) ActionInfo(ctx context.Context, instanceID, actionID string) (*InstanceAction, error) {
	path := fmt.Sprintf("api/v1/instances/%s/actions/%s", instanceID, actionID)
	req, err := is.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var actionRoot instanceActionRoot
	resp, err := is.client.Do(ctx, req, &actionRoot)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting instance action: %v", resp.StatusCode)
	}

	return actionRoot.Action, nil
}

// Actions returns instance's actions list
func (is *InstancesService) Actions(ctx context.Context, instanceID string) ([]InstanceAction, error) {
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var actionsRoot instanceActionsRoot

	if _, err = is.client.Do(ctx, req, &actionsRoot); err != nil {
		return nil, err
	}
	return actionsRoot.Actions, nil
}
