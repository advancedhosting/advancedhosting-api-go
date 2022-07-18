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
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const forwardingRuleResponse = `{
	"request_protocol": "tcp",
	"request_port": 1,
	"communication_protocol": "tcp",
	"communication_port": 1,
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"state": "defined"
}`

const lbPrivateNetworkResponse = `{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"addresses": [{
		"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		"server_id": "820efca4-4a15-4ab7-82fc-9e76f6d61325",
		"address": "192.168.0.1"
	}],
	"state": "defined"
}`

const lbBackendNodeResponse = `{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"cloud_server_id": "296c68d1-bbc6-488b-a29f-f0578806b001",
	"state": "defined"
}`

const lbHealthCheckResponse = `{
	"type": "tcp",
	"url": "string",
	"interval": 10,
	"timeout": 5,
	"unhealthy_threshold": 2,
	"healthy_threshold": 2,
	"port": 0,
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"state": "defined"
}`

const lbIPAddressResponse = `{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"type": "string",
	"address": "192.168.0.1",
	"state": "defined"
}`

const loadBalancerResponse = `{
	"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
	"name": "string",
	"meta": {
		"kubernetes": {
			"cluster": {
				"id": "id",
				"number": "number"
			}
		}
	},
	"datacenter_id": "5839cebe-c7a5-4a27-8253-7bd619ca430d",
	"state": "defined",
	"balancing_algorithm": "round_robin",
	"ip_addresses": [{
		"wcs_ip_address_id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		"type": "string",
		"address": "192.168.0.1",
		"state": "defined"
	}],
	"private_networks": [{
		"wcs_private_network_id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		"addresses": [{
			"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
			"server_id": "820efca4-4a15-4ab7-82fc-9e76f6d61325",
			"address": "192.168.0.1"
		}],
		"state": "defined"
	}],
	"forwarding_rules": [{
		"request_protocol": "tcp",
		"request_port": 1,
		"communication_protocol": "tcp",
		"communication_port": 1,
		"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		"state": "defined"
	}],
	"backend_nodes": [{
		"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		"cloud_server_id": "296c68d1-bbc6-488b-a29f-f0578806b001",
		"state": "defined"
	}],
	"health_check": {
		"type": "tcp",
		"url": "string",
		"interval": 10,
		"timeout": 5,
		"unhealthy_threshold": 2,
		"healthy_threshold": 2,
		"port": 0,
		"id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
		"state": "defined"
	}
}`

var (
	loadBalancerListResponse               = fmt.Sprintf(`{"load_balancers": [%s]}`, loadBalancerResponse)
	loadBalancerGetResponse                = fmt.Sprintf(`{"load_balancer": %s}`, loadBalancerResponse)
	loadBalancerForwardingRuleListResponse = fmt.Sprintf(`{"forwarding_rules": [%s]}`, forwardingRuleResponse)
	loadBalancerForwardingRuleGetResponse  = fmt.Sprintf(`{"forwarding_rule": %s}`, forwardingRuleResponse)
	loadBalancerPrivateNetworkListResponse = fmt.Sprintf(`{"private_networks": [%s]}`, lbPrivateNetworkResponse)
	loadBalancerPrivateNetworkGetResponse  = fmt.Sprintf(`{"private_network": %s}`, lbPrivateNetworkResponse)
	loadBalancerBackendNodeListResponse    = fmt.Sprintf(`{"backend_nodes": [%s]}`, lbBackendNodeResponse)
	loadBalancerBackendNodeGetResponse     = fmt.Sprintf(`{"backend_node": %s}`, lbBackendNodeResponse)
	loadBalancerHealthCheckListResponse    = fmt.Sprintf(`{"health_checks": [%s]}`, lbHealthCheckResponse)
	loadBalancerHealthCheckGetResponse     = fmt.Sprintf(`{"health_check": %s}`, lbHealthCheckResponse)
	loadBalancerIPAddressListResponse      = fmt.Sprintf(`{"ip_addresses": [%s]}`, lbIPAddressResponse)
	loadBalancerIPAddressGetResponse       = fmt.Sprintf(`{"ip_address": %s}`, lbIPAddressResponse)
)

func TestLoadBalancers_List(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerListResponse}
	server := newFakeServer("/api/v1/load_balancers", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	loadBalancers, err := api.LoadBalancers.List(ctx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult loadBalancersRoot
	if err = json.Unmarshal([]byte(loadBalancerListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.LoadBalancers, loadBalancers) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, loadBalancers)
	}
}

func TestLoadBalancers_Get(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerGetResponse}
	server := newFakeServer("/api/v1/load_balancers/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult loadBalancerRoot
	if err := json.Unmarshal([]byte(loadBalancerGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	loadBalancer, err := api.LoadBalancers.Get(ctx, "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if loadBalancer == nil || loadBalancer.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", loadBalancer)
	}

	if !reflect.DeepEqual(expectedResult.LoadBalancer, loadBalancer) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, loadBalancer)
	}
}

func TestLoadBalancers_Create(t *testing.T) {

	request := &LoadBalancerCreateRequest{
		Name:                  "test-name",
		DatacenterID:          "test_dc_id",
		CreatePublicIPAddress: true,
		BalancingAlgorithm:    "round_robin",
		IPAddressIDs:          []string{"ip_address_id1", "ip_address_id2"},
		PrivateNetworkIDs:     []string{"pn_id1", "pn_id2"},
		ForwardingRules: []LBForwardingRuleCreateRequest{
			{
				RequestProtocol:       "tcp",
				RequestPort:           80,
				CommunicationProtocol: "http",
				CommunicationPort:     8080,
			},
		},
		HealthCheck: LBHealthCheckCreateRequest{
			Type:     "tcp",
			Interval: 10,
			Port:     9090,
		},
		BackendNodes: []LBBackendNodeCreateRequest{
			{
				CloudServerID: "cs_id1",
			},
		},
		Meta: map[string]interface{}{
			"kubernetes": map[string]map[string]string{
				"cluster": {
					"id":     "id",
					"number": "number",
				},
			},
		},
	}
	fakeResponse := &fakeServerResponse{
		responseBody: loadBalancerGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	loadBalancer, err := api.LoadBalancers.Create(ctx, request)

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if loadBalancer == nil {
		t.Errorf("Empty response")
	}

	var expectedResult loadBalancerRoot
	if err = json.Unmarshal([]byte(loadBalancerGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error %s", err)
	}

	if !reflect.DeepEqual(expectedResult.LoadBalancer, loadBalancer) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, loadBalancer)
	}

}

func TestLoadBalancers_Update(t *testing.T) {

	request := &LoadBalancerUpdateRequest{
		Name:               "test-name",
		BalancingAlgorithm: "round_robin",
	}

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.Update(ctx, "test_lb_id", request)

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_Delete(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.Delete(ctx, "test_lb_id")

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_ListForwardingRules(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerForwardingRuleListResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/forwarding_rules", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	forwardingRules, err := api.LoadBalancers.ListForwardingRules(ctx, "test_lb_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult lbForwardingRulesRoot
	if err := json.Unmarshal([]byte(loadBalancerForwardingRuleListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.ForwardingRules, forwardingRules) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, forwardingRules)
	}
}

func TestLoadBalancers_GetForwardingRule(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerForwardingRuleGetResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/forwarding_rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult lbForwardingRuleRoot
	if err := json.Unmarshal([]byte(loadBalancerForwardingRuleGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	forwardingRule, err := api.LoadBalancers.GetForwardingRule(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if forwardingRule == nil || forwardingRule.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", forwardingRule)
	}

	if !reflect.DeepEqual(expectedResult.ForwardingRule, forwardingRule) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, forwardingRule)
	}
}

func TestLoadBalancers_CreateForwardingRule(t *testing.T) {

	request := &LBForwardingRuleCreateRequest{
		RequestProtocol:       "tcp",
		RequestPort:           80,
		CommunicationProtocol: "http",
		CommunicationPort:     8080,
	}

	fakeResponse := &fakeServerResponse{
		responseBody: loadBalancerForwardingRuleGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/forwarding_rules", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	forwardingRule, err := api.LoadBalancers.CreateForwardingRule(ctx, "test_lb_id", request)

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	if forwardingRule == nil {
		t.Errorf("Empty response")
	}

	var expectedResult lbForwardingRuleRoot
	if err = json.Unmarshal([]byte(loadBalancerForwardingRuleGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.ForwardingRule, forwardingRule) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, forwardingRule)
	}

}

func TestLoadBalancers_DeleteForwardingRule(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/forwarding_rules/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.DeleteForwardingRule(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_ListPrivateNetworks(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerPrivateNetworkListResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/private_networks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetworks, err := api.LoadBalancers.ListPrivateNetworks(ctx, "test_lb_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult lbPrivateNetworksRoot
	if err := json.Unmarshal([]byte(loadBalancerPrivateNetworkListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.PrivateNetworks, privateNetworks) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateNetworks)
	}
}

func TestLoadBalancers_GetPrivateNetwork(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerPrivateNetworkGetResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/private_networks/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult lbPrivateNetworkRoot
	if err := json.Unmarshal([]byte(loadBalancerPrivateNetworkGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	privateNetwork, err := api.LoadBalancers.GetPrivateNetwork(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if privateNetwork == nil || privateNetwork.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", privateNetwork)
	}

	if !reflect.DeepEqual(expectedResult.PrivateNetwork, privateNetwork) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateNetwork)
	}
}

func TestLoadBalancers_ConnectPrivateNetwork(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: loadBalancerPrivateNetworkListResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/private_networks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	privateNetworks, err := api.LoadBalancers.ConnectPrivateNetworks(ctx, "test_lb_id", []string{"test_pn_id"})

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	var expectedResult lbPrivateNetworksRoot
	if err = json.Unmarshal([]byte(loadBalancerPrivateNetworkListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.PrivateNetworks, privateNetworks) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, privateNetworks)
	}

}

func TestLoadBalancers_DisconnectPrivateNetwork(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/private_networks/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.DisconnectPrivateNetwork(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_ListBackendNodes(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerBackendNodeListResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/backend_nodes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	backendNodes, err := api.LoadBalancers.ListBackendNodes(ctx, "test_lb_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult lbBackendNodesRoot
	if err := json.Unmarshal([]byte(loadBalancerBackendNodeListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.BackendNodes, backendNodes) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, backendNodes)
	}
}

func TestLoadBalancers_GetBackendNode(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerBackendNodeGetResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/backend_nodes/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult lbBackendNodeRoot
	if err := json.Unmarshal([]byte(loadBalancerBackendNodeGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	backendNode, err := api.LoadBalancers.GetBackendNode(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if backendNode == nil || backendNode.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", backendNode)
	}

	if !reflect.DeepEqual(expectedResult.BackendNode, backendNode) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, backendNode)
	}
}

func TestLoadBalancers_AddBackendNodes(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: loadBalancerBackendNodeListResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/backend_nodes", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	backendNodes, err := api.LoadBalancers.AddBackendNodes(ctx, "test_lb_id", []string{"test_bn_id_1", "test_bn_id_2"})

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	var expectedResult lbBackendNodesRoot
	if err = json.Unmarshal([]byte(loadBalancerBackendNodeListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.BackendNodes, backendNodes) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, backendNodes)
	}

}

func TestLoadBalancers_DeleteBackendNode(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/backend_nodes/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.DeleteBackendNode(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_ListHealthChecks(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerHealthCheckListResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/health_checks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	healthChecks, err := api.LoadBalancers.ListHealthChecks(ctx, "test_lb_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult lbHealthChecksRoot
	if err := json.Unmarshal([]byte(loadBalancerHealthCheckListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.HealthChecks, healthChecks) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, healthChecks)
	}
}

func TestLoadBalancers_GetHealthCheck(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerHealthCheckGetResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/health_checks/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult lbHealthCheckRoot
	if err := json.Unmarshal([]byte(loadBalancerHealthCheckGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	healthCheck, err := api.LoadBalancers.GetHealthCheck(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if healthCheck == nil || healthCheck.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", healthCheck)
	}

	if !reflect.DeepEqual(expectedResult.HealthCheck, healthCheck) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, healthCheck)
	}
}

func TestLoadBalancers_CreateHealthCheck(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: loadBalancerHealthCheckGetResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/health_checks", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	request := &LBHealthCheckCreateRequest{
		Type:               "http",
		URL:                "/",
		Interval:           10,
		Timeout:            2,
		UnhealthyThreshold: 3,
		HealthyThreshold:   3,
		Port:               8080,
	}

	healthCheck, err := api.LoadBalancers.CreateHealthCheck(ctx, "test_lb_id", request)

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	var expectedResult lbHealthCheckRoot
	if err = json.Unmarshal([]byte(loadBalancerHealthCheckGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.HealthCheck, healthCheck) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, healthCheck)
	}
}

func TestLoadBalancers_UpdateHealthCheck(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/health_checks/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	request := &LBHealthCheckUpdateRequest{
		UnhealthyThreshold: 3,
		HealthyThreshold:   3,
	}

	err := api.LoadBalancers.UpdateHealthCheck(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08", request)

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_DeleteHealthCheck(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/health_checks/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.DeleteHealthCheck(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}

func TestLoadBalancers_ListIPAddresses(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerIPAddressListResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	ipAddresses, err := api.LoadBalancers.ListIPAddresses(ctx, "test_lb_id")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	var expectedResult lbIPAddressesRoot
	if err := json.Unmarshal([]byte(loadBalancerIPAddressListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.IPAddresses, ipAddresses) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, ipAddresses)
	}
}

func TestLoadBalancers_GetIPAddress(t *testing.T) {
	fakeResponse := &fakeServerResponse{responseBody: loadBalancerIPAddressGetResponse}
	server := newFakeServer("/api/v1/load_balancers/test_lb_id/ip_addresses/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	var expectedResult lbIPAddressRoot
	if err := json.Unmarshal([]byte(loadBalancerIPAddressGetResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	ipAddress, err := api.LoadBalancers.GetIPAddress(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ipAddress == nil || ipAddress.ID != "497f6eca-6276-4993-bfeb-53cbbbba6f08" {
		t.Errorf("Invalid response: %v", ipAddress)
	}

	if !reflect.DeepEqual(expectedResult.IPAddress, ipAddress) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, ipAddress)
	}
}

func TestLoadBalancers_AssignIPAddresses(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: loadBalancerIPAddressListResponse,
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/ip_addresses", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()

	ipAddresses, err := api.LoadBalancers.AssignIPAddresses(ctx, "test_lb_id", []string{"test_ip_id_1", "test_ip_id_2"})

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	var expectedResult lbIPAddressesRoot
	if err = json.Unmarshal([]byte(loadBalancerIPAddressListResponse), &expectedResult); err != nil {
		t.Errorf("Unexpected Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(expectedResult.IPAddresses, ipAddresses) {
		t.Errorf("unexpected result, expected %v. got: %v", expectedResult, ipAddresses)
	}
}

func TestLoadBalancers_ReleaseIPAddress(t *testing.T) {

	fakeResponse := &fakeServerResponse{
		responseBody: "",
		statusCode:   202,
	}

	server := newFakeServer("/api/v1/load_balancers/test_lb_id/ip_addresses/497f6eca-6276-4993-bfeb-53cbbbba6f08", fakeResponse)

	fakeClientOptions := &ClientOptions{
		Token:      "test_token",
		BaseURL:    server.URL,
		HTTPClient: server.Client(),
	}
	api, _ := NewAPIClient(fakeClientOptions)

	ctx := context.Background()
	err := api.LoadBalancers.ReleaseIPAddress(ctx, "test_lb_id", "497f6eca-6276-4993-bfeb-53cbbbba6f08")

	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

}
