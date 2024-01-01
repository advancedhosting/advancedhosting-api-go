package ah

import (
	"context"
	"testing"
)

func TestWorkerDelete(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: "", statusCode: 202}
	api, _ := newFakeAPIClient(
		"/api/v2/kubernetes/clusters/497f6eca-6276-4993-bfeb-53cbbbba6f08/worker_pools/e312aa01-d123-4a90-9c6d-f7641d2e4cc7/workers/339e3dd3-9734-40ec-8e5c-aa7c4c3be319",
		fakeResponse,
	)

	ctx := context.Background()

	req := &ClusterDeleteNodeRequest{
		Replace: false,
	}

	err := api.KubernetesClusters.DeleteWorker(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08", "e312aa01-d123-4a90-9c6d-f7641d2e4cc7", "339e3dd3-9734-40ec-8e5c-aa7c4c3be319", req)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
