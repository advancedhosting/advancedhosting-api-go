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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
)

var (
	// ErrResourceNotFound is returned when resource is not found
	ErrResourceNotFound = errors.New("resource not found")
)

const defaultAPIURL = "https://api.websa.com"

// APIClient implements communication with AH API
type APIClient struct {
	client                  *http.Client
	apiURL                  *url.URL
	Instances               InstancesAPI
	IPAddresses             IPAddressesAPI
	IPAddressAssignments    IPAddressAssignmentsAPI
	PrivateNetworks         PrivateNetworksAPI
	InstancePrivateNetworks InstancePrivateNetworksAPI
	Volumes                 VolumesAPI
	InstancePlans           InstancePlansAPI
	VolumePlans             VolumePlansAPI
	SSHKeys                 SSHKeysAPI
	Backups                 BackupsAPI
	Datacenters             DatacentersAPI
	Images                  ImagesAPI
	LoadBalancers           LoadBalancersAPI
	KubernetesClusters      KubernetesClustersAPI
	Tokens                  TokensAPI
	// Deprecated: Please use VolumePlans instead.
	VolumeProducts VolumeProductsAPI
	// Deprecated: Please use InstancePlans instead.
	InstanceProducts InstanceProductsAPI
}

// ClientOptions represents options to communicate with AH API
type ClientOptions struct {
	HTTPClient *http.Client
	BaseURL    string
	Token      string
}

func (c *APIClient) newRequest(method string, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	u, err := c.apiURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil

}

func (c *APIClient) list(ctx context.Context, path string, options *ListOptions, v interface{}) error {
	if options != nil {
		params := buildListQuery(options)
		path = fmt.Sprintf("%s?%s", path, params)
	}
	req, err := c.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, v)

	if err != nil {
		return err
	}

	return nil
}

// Do sends an API request
func (c *APIClient) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if c := resp.StatusCode; !(c >= 200 && c <= 299) {
		switch c {
		case http.StatusNotFound:
			err = ErrResourceNotFound
		case http.StatusBadRequest:
			err = fmt.Errorf("bad Request")
		default:
			body, _ := io.ReadAll(resp.Body)
			err = fmt.Errorf(string(body))
		}
		return nil, err
	}

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, err
	}
	return resp, nil

}

// NewAPIClient returns APIClient instance
func NewAPIClient(options *ClientOptions) (*APIClient, error) {

	baseURL := defaultAPIURL
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}

	apiURL, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}
	if options.Token == "" {
		return nil, fmt.Errorf("%s", "invalid token")
	}
	var httpClient *http.Client
	if options.HTTPClient != nil {
		httpClient = options.HTTPClient
	} else {
		token := &oauth2.Token{AccessToken: options.Token}
		httpClient = oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	}

	c := &APIClient{
		client: httpClient,
		apiURL: apiURL,
	}
	c.Instances = &InstancesService{client: c}
	c.IPAddresses = &IPAddressesService{client: c}
	c.IPAddressAssignments = &IPAddressAssignmentsService{client: c}
	c.PrivateNetworks = &PrivateNetworksService{client: c}
	c.InstancePrivateNetworks = &InstancePrivateNetworksService{client: c}
	c.Volumes = &VolumesService{client: c}
	c.SSHKeys = &SSHKeysService{client: c}
	c.Backups = &BackupsService{client: c}
	c.Datacenters = &DatacentersService{client: c}
	c.Images = &ImagesService{client: c}
	c.LoadBalancers = &LoadBalancersService{client: c}
	c.KubernetesClusters = &KubernetesClustersService{client: c}
	c.Tokens = &TokensService{client: c}
	c.InstancePlans = &InstancePlansService{client: c}
	c.VolumePlans = &VolumePlansService{client: c}
	c.VolumeProducts = &VolumeProductsService{client: c}
	c.InstanceProducts = &InstanceProductsService{client: c}
	return c, nil
}
