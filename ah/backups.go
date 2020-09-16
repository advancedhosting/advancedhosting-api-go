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
