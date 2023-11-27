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
	"net/http"
	"net/http/httptest"
	"testing"
)

func newFakeEmptyServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, _ = rw.Write([]byte("test"))
	}))

	return server
}

func TestNewAPIClientWithValidURL(t *testing.T) {
	server := newFakeEmptyServer()

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	_, err := NewAPIClient(fakeClientOptions)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

}

func TestNewAPIClientWithInvalidURL(t *testing.T) {
	server := newFakeEmptyServer()

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    "htp/invalid",
		HTTPClient: server.Client(),
	}

	_, err := NewAPIClient(fakeClientOptions)
	if err == nil {
		t.Errorf("Unexpected api client")
	}

}
