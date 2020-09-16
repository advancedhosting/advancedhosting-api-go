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
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const instancesBackupsResponse = `{
	"instance_id": "61463ad8-f5a2-493a-80a0-7b0059ccaafb",
	"instance_name": "kube-adm",
	"instance_removed": false,
	"instance_snapshot_by_schedule": false,
	"backups": [
		{
			"id": "437696c6-6b56-466d-92b2-6f5231124fbb",
			"instance_id": "61463ad8-f5a2-493a-80a0-7b0059ccaafb",
			"created_at": "2020-07-03T08:33:27.127Z",
			"updated_at": "2020-07-03T08:34:29.148Z",
			"name": "WVDS113828_2020-07-03T083327",
			"size": 1759379456,
			"public": false,
			"status": "active",
			"type": "backup",
			"note": "Init 03.07.2020 at 11:33",
			"min_disk_size": 40000000000
		}
	]
}`

var (
	instancesBackupsListResponse = fmt.Sprintf(`{"instances_backups": [%s]}`, instancesBackupsResponse)
)

func TestBackups_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: instancesBackupsResponse}
	server := newFakeServer("/api/v1/backups", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	backups, err := api.Backups.List(ctx, nil)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult instancesBackupsRoot
	json.Unmarshal([]byte(instancesBackupsResponse), &expectedResult)

	if !reflect.DeepEqual(expectedResult.InstancesBackups, backups) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, backups)
	}
}
