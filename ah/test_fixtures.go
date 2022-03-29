/*
Copyright 2021 Advanced Hosting

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
)

type fakeServerResponse struct {
	responseBody string
	statusCode   int
}

func newFakeServer(url string, response *fakeServerResponse) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(url, func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("content-type", "application/json")
		if response.statusCode != 0 {
			rw.WriteHeader(response.statusCode)
		}
		_, _ = rw.Write([]byte(response.responseBody))
	})
	return httptest.NewServer(mux)
}

func newFakeClientOptions(server *httptest.Server) *ClientOptions {
	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	return fakeClientOptions
}

func newFakeAPIClient(url string, response *fakeServerResponse) (*APIClient, error) {
	server := newFakeServer(url, response)
	options := newFakeClientOptions(server)
	api, err := NewAPIClient(options)
	return api, err
}
