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
	"fmt"
	"net/http"
)

// Backup object
type Backup struct {
	ID                         string `json:"id,omitempty"`
	CreatedAt                  string `json:"created_at,omitempty"`
	UpdatedAt                  string `json:"updated_at,omitempty"`
	Name                       string `json:"name,omitempty"`
	Status                     string `json:"status,omitempty"`
	Type                       string `json:"type,omitempty"`
	Note                       string `json:"note,omitempty"`
	InstanceID                 string `json:"instance_id,omitempty"`
	InstanceName               string `json:"instance_name,omitempty"`
	Size                       int    `json:"size,omitempty"`
	MinDiskSize                int    `json:"min_disk_size,omitempty"`
	InstanceRemoved            bool   `json:"instance_removed,omitempty"`
	InstanceSnapshotBySchedule bool   `json:"instance_snapshot_by_schedule,omitempty"`
	Public                     bool   `json:"public,omitempty"`
}

// InstanceBackups object
type InstanceBackups struct {
	InstanceID                 string   `json:"instance_id,omitempty"`
	InstanceName               string   `json:"instance_name,omitempty"`
	Backups                    []Backup `json:"backups,omitempty"`
	InstanceRemoved            bool     `json:"instance_removed,omitempty"`
	InstanceSnapshotBySchedule bool     `json:"instance_snapshot_by_schedule,omitempty"`
}

// BackupsAPI is an interface for backups.
type BackupsAPI interface {
	List(context.Context, *ListOptions) ([]InstanceBackups, error)
	Get(context.Context, string) (*Backup, error)
	Update(context.Context, string, *BackUpUpdateRequest) (*Backup, error)
	Delete(context.Context, string) (*Action, error)
}

// BackupsService implements BackupsAPI interface.
type BackupsService struct {
	client *APIClient
}

type BackupListRoot struct {
	Backups []BackupWithEmbeddedInstance `json:"backups"`
}

type InstanceInfo struct {
	Name                   string `json:"name"`
	SnapshotBySchedule     bool   `json:"snapshot_by_schedule"`
	InstancePrivateCluster bool   `json:"instance_private_cluster"`
	InstanceRemoved        bool   `json:"instance_removed"`
}

type BackupWithEmbeddedInstance struct {
	Instance InstanceInfo `json:"instance"`
	Backup
}

// List returns all available private networks
func (bs *BackupsService) List(ctx context.Context, options *ListOptions) ([]InstanceBackups, error) {
	path := "api/v1/backups"

	var bRoot BackupListRoot

	if err := bs.client.list(ctx, path, options, &bRoot); err != nil {
		return nil, err
	}

	// Group backups by instance ID
	grouped := map[string]*InstanceBackups{}

	for _, b := range bRoot.Backups {
		instID := b.InstanceID

		if _, exists := grouped[instID]; !exists {
			grouped[instID] = &InstanceBackups{
				InstanceID:                 instID,
				InstanceName:               b.Instance.Name,
				InstanceRemoved:            b.Instance.InstanceRemoved,
				InstanceSnapshotBySchedule: b.Instance.SnapshotBySchedule,
				Backups:                    []Backup{},
			}
		}

		// Include the backup (without embedded instance)
		grouped[instID].Backups = append(grouped[instID].Backups, b.Backup)
	}

	// Convert map to slice
	var result []InstanceBackups
	for _, ib := range grouped {
		result = append(result, *ib)
	}

	return result, nil
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
