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

// InstanceImage object
type InstanceImage struct {
	*Image
	Type string `json:"type,omitempty"`
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
	Datacenter                 *Datacenter              `json:"datacenter,omitempty"`
	Features                   []string                 `json:"features,omitempty"`
	Networks                   *InstanceNetworks        `json:"networks,omitempty"`
	CurrentAction              *InstanceAction          `json:"current_action,omitempty"`
	LastAction                 *InstanceAction          `json:"last_action,omitempty"`
	Reason                     string                   `json:"reason,omitempty"`
	SnapshotBySchedule         bool                     `json:"snapshot_by_schedule,omitempty"`
	SnapshotPeriod             string                   `json:"snapshot_period,omitempty"`
	MaxVolumesNumber           int                      `json:"max_volumes_number,omitempty"`
	PrivateNetworks            []InstancePrivateNetwork `json:"instance_private_networks,omitempty"`
	IPAddresses                []InstanceIPAddress      `json:"instance_ip_addresses,omitempty"`
	Image                      *InstanceImage           `json:"image,omitempty"`
	Volumes                    []Volume                 `json:"volumes,omitempty"`
}

// InstanceAction object
type InstanceAction struct {
	*Action
	ResultParams *struct {
		SnapshotID string `json:"snapshot_id,omitempty"`
	} `json:"result_params,omitempty"`
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
	DatacenterID          string   `json:"datacenter_id,omitempty"`
	DatacenterSlug        string   `json:"datacenter_slug,omitempty"`
	ImageID               string   `json:"image_id,omitempty"`
	ImageSlug             string   `json:"image_slug,omitempty"`
	ProductID             string   `json:"product_id,omitempty"`
	ProductSlug           string   `json:"product_slug,omitempty"`
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
	SetPrimaryIP(context.Context, string, string) (*Action, error)
	AttachVolume(context.Context, string, string) (*Action, error)
	DetachVolume(context.Context, string, string) (*Action, error)
	ActionInfo(context.Context, string, string) (*InstanceAction, error)
	Actions(context.Context, string) ([]InstanceAction, error)
	AvailableVolumes(context.Context, string, *ListOptions) ([]Volume, *Meta, error)
	CreateBackup(context.Context, string, string) (*InstanceAction, error)
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

	resp, err := is.client.Do(ctx, req, nil)
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
func (is *InstancesService) SetPrimaryIP(ctx context.Context, instanceID, ipAssignmentID string) (*Action, error) {
	request := &instanceSetPrimaryIPRequest{
		InstanceIPAddressID: ipAssignmentID,
		Type:                "set_primary_ip",
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var aRoot actionRoot
	_, err = is.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}

	return aRoot.Action, nil
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

	var aRoot instanceActionRoot
	resp, err := is.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting instance action: %v", resp.StatusCode)
	}

	return aRoot.Action, nil
}

// Actions returns instance's actions list
func (is *InstancesService) Actions(ctx context.Context, instanceID string) ([]InstanceAction, error) {
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var asRoot instanceActionsRoot

	if _, err = is.client.Do(ctx, req, &asRoot); err != nil {
		return nil, err
	}
	return asRoot.Actions, nil
}

type instanceAttachVolumeRequest struct {
	VolumeID string `json:"volume_id"`
	Type     string `json:"type"`
}

// AttachVolume connects volume to the instance
func (is *InstancesService) AttachVolume(ctx context.Context, instanceID, volumeID string) (*Action, error) {
	request := &instanceAttachVolumeRequest{
		VolumeID: volumeID,
		Type:     "attach_volume",
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var aRoot actionRoot
	_, err = is.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}

	return aRoot.Action, nil
}

type instanceDetachVolumeRequest struct {
	VolumeID string `json:"volume_id"`
	Type     string `json:"type"`
}

// DetachVolume disconnects volume to the instance
func (is *InstancesService) DetachVolume(ctx context.Context, instanceID, volumeID string) (*Action, error) {
	request := &instanceDetachVolumeRequest{
		VolumeID: volumeID,
		Type:     "detach_volume",
	}
	path := fmt.Sprintf("api/v1/instances/%s/actions", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var aRoot actionRoot
	_, err = is.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}

	return aRoot.Action, nil
}

// AvailableVolumes return all attached volumes to the instance.
func (is *InstancesService) AvailableVolumes(ctx context.Context, instanceID string, options *ListOptions) ([]Volume, *Meta, error) {
	path := fmt.Sprintf("api/v1/instances/%s/available_volumes", instanceID)

	var vsRoot volumesRoot

	if err := is.client.list(ctx, path, options, &vsRoot); err != nil {
		return nil, nil, err
	}
	return vsRoot.Volumes, vsRoot.Meta, nil
}

// CreateBackup creates instance's backups
func (is *InstancesService) CreateBackup(ctx context.Context, instanceID, note string) (*InstanceAction, error) {

	var request = &struct {
		Note string `json:"note"`
	}{note}

	path := fmt.Sprintf("api/v1/instances/%s/backups", instanceID)
	req, err := is.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var aRoot instanceActionRoot
	_, err = is.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}

	return aRoot.Action, nil
}
