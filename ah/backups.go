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

// Backup object
type Backup struct {
	ID          string `json:"id,omitempty"`
	InstanceID  string `json:"instance_id,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	Name        string `json:"name,omitempty"`
	Size        int    `json:"size,omitempty"`
	Public      bool   `json:"public,omitempty"`
	Status      string `json:"status,omitempty"`
	Type        string `json:"type,omitempty"`
	Note        string `json:"note,omitempty"`
	MinDiskSize int    `json:"min_disk_size,omitempty"`
}

// InstanceBackup object
type InstanceBackup struct {
	InstanceID                 string   `json:"instance_id,omitempty"`
	InstanceName               string   `json:"instance_name,omitempty"`
	InstanceRemoved            bool     `json:"instance_removed,omitempty"`
	InstanceSnapshotBySchedule bool     `json:"instance_snapshot_by_schedule,omitempty"`
	Backups                    []Backup `json:"backups,omitempty"`
}

// BackupsAPI is an interface for backups.
type BackupsAPI interface {
	List(context.Context, *ListOptions) ([]InstanceBackup, error)
	Get(context.Context, string) (*Backup, error)
	Update(context.Context, string, *BackUpUpdateRequest) (*Backup, error)
	Delete(context.Context, string) (*Action, error)
}

// BackupsService implements BackupsAPI interface.
type BackupsService struct {
	client *APIClient
}

type instancesBackupsRoot struct {
	InstancesBackups []InstanceBackup `json:"instances_backups"`
}

// List returns all available private networks
func (bs *BackupsService) List(ctx context.Context, options *ListOptions) ([]InstanceBackup, error) {
	path := "api/v1/backups"

	var ibRoot instancesBackupsRoot

	if err := bs.client.list(ctx, path, options, &ibRoot); err != nil {
		return nil, err
	}

	return ibRoot.InstancesBackups, nil

}

type backupRoot struct {
	Backup *Backup `json:"backup"`
}

// Get backup info
func (bs *BackupsService) Get(ctx context.Context, backupID string) (*Backup, error) {
	path := fmt.Sprintf("api/v1/backups/%s", backupID)
	req, err := bs.client.newRequest(http.MethodGet, path, nil)

	if err != nil {
		return nil, err
	}

	var bRoot backupRoot
	if _, err := bs.client.Do(ctx, req, &bRoot); err != nil {
		return nil, err
	}

	return bRoot.Backup, nil
}

// BackUpUpdateRequest represents a request to update a volume.
type BackUpUpdateRequest struct {
	Note string `json:"note,omitempty"`
	Name string `json:"name,omitempty"`
}

// Update backup
func (bs *BackupsService) Update(ctx context.Context, backupID string, request *BackUpUpdateRequest) (*Backup, error) {
	path := fmt.Sprintf("api/v1/backups/%s", backupID)
	req, err := bs.client.newRequest(http.MethodPut, path, request)

	if err != nil {
		return nil, err
	}

	var bRoot backupRoot
	if _, err := bs.client.Do(ctx, req, &bRoot); err != nil {
		return nil, err
	}

	return bRoot.Backup, nil
}

// Delete backup
func (bs *BackupsService) Delete(ctx context.Context, backupID string) (*Action, error) {
	path := fmt.Sprintf("api/v1/backups/%s", backupID)
	req, err := bs.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	var aRoot actionRoot
	_, err = bs.client.Do(ctx, req, &aRoot)
	if err != nil {
		return nil, err
	}
	return aRoot.Action, nil
}
