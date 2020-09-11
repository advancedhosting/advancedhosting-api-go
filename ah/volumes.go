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
	"fmt"
	"net/http"
)

// Volume object
type Volume struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	FileSystem string `json:"file_system,omitempty"`
	Size       int    `json:"size,omitempty"`
	Port       int    `json:"port,omitempty"`
	State      string `json:"state,omitempty"`
	OriginalID string `json:"original_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	AttachedAt string `json:"attached_at,omitempty"`
	ProductID  string `json:"product_id,omitempty"`
	Instance   *struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"instance,omitempty"`
	Product *struct {
		ID            string `json:"id,omitempty"`
		Name          string `json:"name,omitempty"`
		MinVolumeSize int    `json:"min_volume_size,omitempty"`
		MaxVolumeSize int    `json:"max_volume_size,omitempty"`
	} `json:"product,omitempty"`
	VolumePool *struct {
		Name             string   `json:"name,omitempty"`
		DatacenterIDs    []string `json:"datacenter_ids,omitempty"`
		ReplicationLevel int      `json:"replication_level,omitempty"`
	} `json:"volume_pool,omitempty"`
}

// VolumesAPI is an interface for volumes.
type VolumesAPI interface {
	List(context.Context, *ListOptions) ([]Volume, *Meta, error)
	Get(context.Context, string) (*Volume, error)
	Create(context.Context, *VolumeCreateRequest) (*Volume, error)
	Update(context.Context, string, *VolumeUpdateRequest) (*Volume, error)
	Copy(context.Context, string, string, string) (*Action, error)
	Resize(context.Context, string, int) (*Action, error)
	Delete(context.Context, string) error
}

// VolumesService implements VolumesAPI interface.
type VolumesService struct {
	client *APIClient
}

type volumesRoot struct {
	Volumes []Volume `json:"volumes"`
	Meta    *Meta    `json:"meta"`
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
	Name       string `json:"name"`
	Size       int    `json:"size"`
	ProductID  string `json:"product_id"`
	FileSystem string `json:"file_system,omitempty"`
	InstanceID string `json:"instance_id,omitempty"`
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
	Name string `json:"name,omitempty"`
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

type volumeCopyActionRequest struct {
	Name      string `json:"name"`
	ProductID string `json:"product_id"`
	Type      string `json:"type"`
}

// Copy volume
func (vs *VolumesService) Copy(ctx context.Context, volumeID, name, productID string) (*Action, error) {
	path := fmt.Sprintf("api/v1/volumes/%s/actions", volumeID)

	request := &volumeCopyActionRequest{
		Name:      name,
		ProductID: productID,
		Type:      "copy",
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

type volumeResizeActionRequest struct {
	Size int    `json:"size"`
	Type string `json:"type"`
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
