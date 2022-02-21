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
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const clusterResponse = `{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"name": "string",
	"datacenter_id": "5839cebe-c7a5-4a27-8253-7bd619ca430d",
	"state": "defined",
	"count": 0,
	"created_at": "2019-08-24T14:15:22Z",
	"plan_id": 0,
	"number": "string"
}`

var (
	clusterGetResponse = fmt.Sprintf(`{"cluster": %s}`, clusterResponse)
)

func TestClusters_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: clusterGetResponse}
	server := newFakeServer("/api/v1/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult clusterRoot
	if err := json.Unmarshal([]byte(clusterGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	cluster, err := api.Clusters.Get(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if cluster == nil || cluster.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", cluster)
	}

	if !reflect.DeepEqual(expectedResult.Cluster, cluster) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, cluster)
	}
}
