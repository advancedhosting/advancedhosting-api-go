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
	"encoding/json"
	"fmt"
	"net/http"
)

// Volume object
type Volume struct {
	Instance *struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"instance,omitempty"`
	VolumePool *struct {
		Name             string   `json:"name,omitempty"`
		DatacenterIDs    []string `json:"datacenter_ids,omitempty"`
		ReplicationLevel int      `json:"replication_level,omitempty"`
	} `json:"volume_pool,omitempty"`
	ID         string          `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	FileSystem string          `json:"file_system,omitempty"`
	State      string          `json:"state,omitempty"`
	Number     string          `json:"number,omitempty"`
	OriginalID string          `json:"original_id,omitempty"`
	CreatedAt  string          `json:"created_at,omitempty"`
	AttachedAt string          `json:"attached_at,omitempty"`
	ProductID  string          `json:"product_id,omitempty"`
	Meta       json.RawMessage `json:"meta,omitempty"`
	Size       int             `json:"size,omitempty"`
	Port       int             `json:"port,omitempty"`
	PlanID     int             `json:"plan_id,omitempty"`
}

// VolumeAction object
type VolumeAction struct {
	*Action
	ResultParams *struct {
		CopiedVolumeID string `json:"copied_volume_id,omitempty"`
	} `json:"result_params,omitempty"`
}

type volumeActionRoot struct {
	Action *VolumeAction `json:"action"`
}

type volumeActionsRoot struct {
	Actions []VolumeAction `json:"actions"`
}

// VolumesAPI is an interface for volumes.
type VolumesAPI interface {
	List(context.Context, *ListOptions) ([]Volume, *Meta, error)
	Get(context.Context, string) (*Volume, error)
	Create(context.Context, *VolumeCreateRequest) (*Volume, error)
	Update(context.Context, string, *VolumeUpdateRequest) (*Volume, error)
	Copy(context.Context, string, *VolumeCopyActionRequest) (*VolumeAction, error)
	Resize(context.Context, string, int) (*Action, error)
	ActionInfo(context.Context, string, string) (*VolumeAction, error)
	Actions(context.Context, string) ([]VolumeAction, error)
	Delete(context.Context, string) error
}

// VolumesService implements VolumesAPI interface.
type VolumesService struct {
	client *APIClient
}

type volumesRoot struct {
	Meta    *Meta    `json:"meta"`
	Volumes []Volume `json:"volumes"`
}

// List returns all available private networks
func (vs *VolumesService) List(ctx context.Context, options *ListOptions) ([]Volume, *Meta, error) {
	path := "api/v1/volumes"

	var vsRoot volumesRoot

	if err := vs.client.list(ctx, path, options, &vsRoot); err != nil {
		return nil, nil, err
	}

	return vsRoot.Volumes, vsRoot.Meta, nil

}

type volumeRoot struct {
	Volume *Volume `json:"volume"`
}

// Get volume
func (vs *VolumesService) Get(ctx context.Context, volumeID string) (*Volume, error) {
	path := fmt.Sprintf("api/v1/volumes/%s", volumeID)
	req, err := vs.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var vRoot volumeRoot
	_, err = vs.client.Do(ctx, req, &vRoot)

	if err != nil {
		return nil, err
	}

	return vRoot.Volume, nil
}

// VolumeCreateRequest object
type VolumeCreateRequest struct {
	Name string `json:"name"`
	Meta string `json:"meta,omitempty"`
	// Deprecated: Please use PlanID instead.
	ProductID string `json:"product_id,omitempty"`
	// Deprecated: Please use PlanSlug instead.
	ProductSlug string `json:"product_slug,omitempty"`
	PlanSlug    string `json:"plan_slug,omitempty"`
	FileSystem  string `json:"file_system,omitempty"`
	InstanceID  string `json:"instance_id,omitempty"`
	Size        int    `json:"size"`
	PlanID      int    `json:"plan_id,omitempty"`
}

// Create volume
func (vs *VolumesService) Create(ctx context.Context, createRequest *VolumeCreateRequest) (*Volume, error) {

	type request struct {
		Volume *VolumeCreateRequest `json:"volume"`
	}
	req, err := vs.client.newRequest(http.MethodPost, "api/v1/volumes", &request{createRequest})
	if err != nil {
		return nil, err
	}

	var vRoot volumeRoot
	if _, err := vs.client.Do(ctx, req, &vRoot); err != nil {
		return nil, err
	}

	return vRoot.Volume, nil
}

// VolumeUpdateRequest represents a request to update a volume.
type VolumeUpdateRequest struct {
	Meta map[string]string `json:"meta,omitempty"`
	Name string            `json:"name,omitempty"`
}

// Update volume
func (vs *VolumesService) Update(ctx context.Context, volumeID string, request *VolumeUpdateRequest) (*Volume, error) {
	path := fmt.Sprintf("api/v1/volumes/%s", volumeID)
	req, err := vs.client.newRequest(http.MethodPut, path, request)

	if err != nil {
		return nil, err
	}

	var vRoot volumeRoot
	if _, err := vs.client.Do(ctx, req, &vRoot); err != nil {
		return nil, err
	}

	return vRoot.Volume, nil
}

// VolumeCopyActionRequest represents a request to create new volume from origin.
type VolumeCopyActionRequest struct {
	Name string
	// Deprecated: Please use PlanID instead.
	ProductID string
	// Deprecated: Please use PlanSlug instead.
	ProductSlug string
	PlanSlug    string
	PlanID      int
}

type volumeCopyActionRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// Deprecated: Please use PlanID instead.
	ProductID string `json:"product_id,omitempty"`
	// Deprecated: Please use PlanSlug instead.
	ProductSlug string `json:"product_slug,omitempty"`
	PlanSlug    string `json:"plan_slug,omitempty"`
	PlanID      int    `json:"plan_id,omitempty"`
}

// Copy volume
func (vs *VolumesService) Copy(ctx context.Context, volumeID string, request *VolumeCopyActionRequest) (*VolumeAction, error) {
	path := fmt.Sprintf("api/v1/volumes/%s/actions", volumeID)

	copyRequest := &volumeCopyActionRequest{
		Name: request.Name,
		Type: "copy",
	}

	if request.PlanSlug != "" {
		copyRequest.PlanSlug = request.PlanSlug
	} else {
		copyRequest.PlanID = request.PlanID
	}

	if request.PlanSlug == "" && request.PlanID == 0 {
		if request.ProductSlug != "" {
			copyRequest.ProductSlug = request.ProductSlug
		} else {
			copyRequest.ProductID = request.ProductID
		}
	}

	req, err := vs.client.newRequest(http.MethodPost, path, copyRequest)

	if err != nil {
		return nil, err
	}

	var aRoot volumeActionRoot
	if _, err := vs.client.Do(ctx, req, &aRoot); err != nil {
		return nil, err
	}

	return aRoot.Action, nil
}

type volumeResizeActionRequest struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

// Resize volume
func (vs *VolumesService) Resize(ctx context.Context, volumeID string, size int) (*Action, error) {
	path := fmt.Sprintf("api/v1/volumes/%s/actions", volumeID)

	request := &volumeResizeActionRequest{
		Size: size,
		Type: "resize",
	}

	req, err := vs.client.newRequest(http.MethodPost, path, request)

	if err != nil {
		return nil, err
	}

	var aRoot actionRoot
	if _, err := vs.client.Do(ctx, req, &aRoot); err != nil {
		return nil, err
	}

	return aRoot.Action, nil
}

// ActionInfo returns volume's action info by action ID
func (vs *VolumesService) ActionInfo(ctx context.Context, volumeID, actionID string) (*VolumeAction, error) {
	path := fmt.Sprintf("api/v1/volumes/%s/actions/%s", volumeID, actionID)
	req, err := vs.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var aRoot volumeActionRoot
	resp, err := vs.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error getting volume action: %v", resp.StatusCode)
	}

	return aRoot.Action, nil
}

// Actions returns volume's actions list
func (vs *VolumesService) Actions(ctx context.Context, volumeID string) ([]VolumeAction, error) {
	path := fmt.Sprintf("api/v1/volumes/%s/actions", volumeID)
	req, err := vs.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var asRoot volumeActionsRoot

	if _, err = vs.client.Do(ctx, req, &asRoot); err != nil {
		return nil, err
	}
	return asRoot.Actions, nil
}

// Delete volume
func (vs *VolumesService) Delete(ctx context.Context, volumeID string) error {
	path := fmt.Sprintf("api/v1/volumes/%s", volumeID)
	req, err := vs.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}
	_, err = vs.client.Do(ctx, req, nil)
	if err != nil {
		return err
	}
	return nil
}
