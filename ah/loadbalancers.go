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
	"fmt"
	"net/http"
)

// LoadBalancer object
type LoadBalancer struct {
	Meta               map[string]interface{} `json:"meta,omitempty"`
	ID                 string                 `json:"id,omitempty"`
	Name               string                 `json:"name,omitempty"`
	DatacenterID       string                 `json:"datacenter_id,omitempty"`
	State              string                 `json:"state,omitempty"`
	BalancingAlgorithm string                 `json:"balancing_algorithm,omitempty"`
	IPAddresses        []LBIPAddress          `json:"ip_addresses,omitempty"`
	PrivateNetworks    []LBPrivateNetwork     `json:"private_networks,omitempty"`
	ForwardingRules    []LBForwardingRule     `json:"forwarding_rules,omitempty"`
	BackendNodes       []LBBackendNode        `json:"backend_nodes,omitempty"`
	HealthChecks       []LBHealthCheck        `json:"health_checks,omitempty"`
}

// LBIPAddress object
type LBIPAddress struct {
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Address string `json:"address,omitempty"`
	State   string `json:"state,omitempty"`
}

// LBPrivateNetworkAddress object
type LBPrivateNetworkAddress struct {
	ID       string `json:"id,omitempty"`
	ServerID string `json:"server_id,omitempty"`
	Address  string `json:"address,omitempty"`
}

// LBPrivateNetwork object
type LBPrivateNetwork struct {
	ID        string                    `json:"id,omitempty"`
	State     string                    `json:"state,omitempty"`
	Addresses []LBPrivateNetworkAddress `json:"addresses,omitempty"`
}

// LBForwardingRule object
type LBForwardingRule struct {
	ID                    string `json:"id,omitempty"`
	State                 string `json:"state,omitempty"`
	RequestProtocol       string `json:"request_protocol,omitempty"`
	CommunicationProtocol string `json:"communication_protocol,omitempty"`
	RequestPort           int    `json:"request_port,omitempty"`
	CommunicationPort     int    `json:"communication_port,omitempty"`
}

// LBBackendNode object
type LBBackendNode struct {
	ID            string `json:"id,omitempty"`
	State         string `json:"state,omitempty"`
	CloudServerID string `json:"cloud_server_id,omitempty"`
}

// LBHealthCheck object
type LBHealthCheck struct {
	ID                 string `json:"id,omitempty"`
	State              string `json:"state,omitempty"`
	Type               string `json:"type,omitempty"`
	URL                string `json:"url,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   int    `json:"Healthy_threshold,omitempty"`
	Port               int    `json:"port,omitempty"`
}

// LoadBalancersAPI is an interface for load balancers.
type LoadBalancersAPI interface {
	List(context.Context) ([]LoadBalancer, error)
	Get(context.Context, string) (*LoadBalancer, error)
	Create(context.Context, *LoadBalancerCreateRequest) (*LoadBalancer, error)
	Update(context.Context, string, *LoadBalancerUpdateRequest) error
	Delete(context.Context, string) error

	ListForwardingRules(context.Context, string) ([]LBForwardingRule, error)
	GetForwardingRule(context.Context, string, string) (*LBForwardingRule, error)
	CreateForwardingRule(context.Context, string, *LBForwardingRuleCreateRequest) (*LBForwardingRule, error)
	DeleteForwardingRule(context.Context, string, string) error

	ListPrivateNetworks(context.Context, string) ([]LBPrivateNetwork, error)
	GetPrivateNetwork(context.Context, string, string) (*LBPrivateNetwork, error)
	ConnectPrivateNetworks(context.Context, string, []string) ([]LBPrivateNetwork, error)
	DisconnectPrivateNetwork(context.Context, string, string) error

	ListBackendNodes(context.Context, string) ([]LBBackendNode, error)
	GetBackendNode(context.Context, string, string) (*LBBackendNode, error)
	AddBackendNodes(context.Context, string, []string) ([]LBBackendNode, error)
	DeleteBackendNode(context.Context, string, string) error

	ListHealthChecks(context.Context, string) ([]LBHealthCheck, error)
	GetHealthCheck(context.Context, string, string) (*LBHealthCheck, error)
	CreateHealthCheck(context.Context, string, *LBHealthCheckCreateRequest) (*LBHealthCheck, error)
	UpdateHealthCheck(context.Context, string, string, *LBHealthCheckUpdateRequest) error
	DeleteHealthCheck(context.Context, string, string) error

	ListIPAddresses(context.Context, string) ([]LBIPAddress, error)
	GetIPAddress(context.Context, string, string) (*LBIPAddress, error)
	AssignIPAddresses(context.Context, string, []string) ([]LBIPAddress, error)
	ReleaseIPAddress(context.Context, string, string) error
}

// LoadBalancersService implements LoadBalancersAPI interface.
type LoadBalancersService struct {
	client *APIClient
}

type loadBalancersRoot struct {
	LoadBalancers []LoadBalancer `json:"load_balancers,omitempty"`
}

// List returns all available load balancers
func (lb *LoadBalancersService) List(ctx context.Context) ([]LoadBalancer, error) {
	path := "api/v1/load_balancers"

	var lbsRoot loadBalancersRoot

	if err := lb.client.list(ctx, path, nil, &lbsRoot); err != nil {
		return nil, err
	}

	return lbsRoot.LoadBalancers, nil
}

// Get load balancer
func (lb *LoadBalancersService) Get(ctx context.Context, lbID string) (*LoadBalancer, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s", lbID)

	req, err := lb.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var lbRoot loadBalancerRoot
	_, err = lb.client.Do(ctx, req, &lbRoot)

	if err != nil {
		return nil, err
	}

	return lbRoot.LoadBalancer, nil
}

// LoadBalancerCreateRequest object
type LoadBalancerCreateRequest struct {
	Meta                  map[string]interface{}          `json:"meta,omitempty"`
	Name                  string                          `json:"name"`
	DatacenterID          string                          `json:"datacenter_id"`
	BalancingAlgorithm    string                          `json:"balancing_algorithm,omitempty"`
	IPAddressIDs          []string                        `json:"ip_address_ids,omitempty"`
	PrivateNetworkIDs     []string                        `json:"private_network_ids,omitempty"`
	ForwardingRules       []LBForwardingRuleCreateRequest `json:"forwarding_rules,omitempty"`
	HealthChecks          []LBHealthCheckCreateRequest    `json:"health_checks,omitempty"`
	BackendNodes          []LBBackendNodeCreateRequest    `json:"backend_nodes,omitempty"`
	CreatePublicIPAddress bool                            `json:"create_public_ip_address"`
}

// LBForwardingRuleCreateRequest object
type LBForwardingRuleCreateRequest struct {
	RequestProtocol       string `json:"request_protocol"`
	CommunicationProtocol string `json:"communication_protocol"`
	RequestPort           int    `json:"request_port"`
	CommunicationPort     int    `json:"communication_port"`
}

// LBHealthCheckCreateRequest object
type LBHealthCheckCreateRequest struct {
	Type               string `json:"type"`
	URL                string `json:"url,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   int    `json:"Healthy_threshold,omitempty"`
	Port               int    `json:"port"`
}

// LBBackendNodeCreateRequest object
type LBBackendNodeCreateRequest struct {
	CloudServerID string `json:"cloud_server_id"`
}

type loadBalancerRoot struct {
	LoadBalancer *LoadBalancer `json:"load_balancer,omitempty"`
}

// Create a load balancer
func (lb *LoadBalancersService) Create(ctx context.Context, createRequest *LoadBalancerCreateRequest) (*LoadBalancer, error) {

	req, err := lb.client.newRequest(http.MethodPost, "api/v1/load_balancers", createRequest)
	if err != nil {
		return nil, err
	}

	var lbInfo loadBalancerRoot
	if _, err := lb.client.Do(ctx, req, &lbInfo); err != nil {
		return nil, err
	}

	return lbInfo.LoadBalancer, nil
}

// LoadBalancerUpdateRequest represents a request to update a load balancer.
type LoadBalancerUpdateRequest struct {
	Name               string `json:"name,omitempty"`
	BalancingAlgorithm string `json:"balancing_algorithm,omitempty"`
}

// Update load balancer
func (lb *LoadBalancersService) Update(ctx context.Context, lbID string, request *LoadBalancerUpdateRequest) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s", lbID)
	req, err := lb.client.newRequest(http.MethodPatch, path, request)

	if err != nil {
		return err
	}

	_, err = lb.client.Do(ctx, req, nil)

	return err
}

// Delete load balancer
func (lb *LoadBalancersService) Delete(ctx context.Context, lbID string) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s", lbID)
	req, err := lb.client.newRequest(http.MethodDelete, path, nil)

	if err != nil {
		return err
	}

	_, err = lb.client.Do(ctx, req, nil)

	return err
}

type lbForwardingRulesRoot struct {
	ForwardingRules []LBForwardingRule `json:"forwarding_rules,omitempty"`
}

type lbForwardingRuleRoot struct {
	ForwardingRule *LBForwardingRule `json:"forwarding_rule,omitempty"`
}

// ListForwardingRules returns all available forwarding rules
func (lb *LoadBalancersService) ListForwardingRules(ctx context.Context, lbID string) ([]LBForwardingRule, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/forwarding_rules", lbID)

	var frsRoot lbForwardingRulesRoot

	if err := lb.client.list(ctx, path, nil, &frsRoot); err != nil {
		return nil, err
	}

	return frsRoot.ForwardingRules, nil
}

// GetForwardingRule returns forwarding rule info
func (lb *LoadBalancersService) GetForwardingRule(ctx context.Context, lbID, frID string) (*LBForwardingRule, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/forwarding_rules/%s", lbID, frID)

	req, err := lb.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var frRoot lbForwardingRuleRoot
	_, err = lb.client.Do(ctx, req, &frRoot)

	if err != nil {
		return nil, err
	}

	return frRoot.ForwardingRule, nil
}

// CreateForwardingRule creates a forwarding rule
func (lb *LoadBalancersService) CreateForwardingRule(ctx context.Context, lbID string, request *LBForwardingRuleCreateRequest) (*LBForwardingRule, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/forwarding_rules", lbID)

	req, err := lb.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var frRoot lbForwardingRuleRoot
	_, err = lb.client.Do(ctx, req, &frRoot)

	if err != nil {
		return nil, err
	}

	return frRoot.ForwardingRule, nil
}

// DeleteForwardingRule remove forwarding rule
func (lb *LoadBalancersService) DeleteForwardingRule(ctx context.Context, lbID, frID string) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s/forwarding_rules/%s", lbID, frID)

	req, err := lb.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	resp, err := lb.client.Do(ctx, req, nil)

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error delete forwarding rule: %v", resp.StatusCode)
	}

	return err
}

type lbPrivateNetworksRoot struct {
	PrivateNetworks []LBPrivateNetwork `json:"private_networks,omitempty"`
}

type lbPrivateNetworkRoot struct {
	PrivateNetwork *LBPrivateNetwork `json:"private_network,omitempty"`
}

// ListPrivateNetworks returns all connected private networks
func (lb *LoadBalancersService) ListPrivateNetworks(ctx context.Context, lbID string) ([]LBPrivateNetwork, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/private_networks", lbID)

	var pnRoot lbPrivateNetworksRoot

	if err := lb.client.list(ctx, path, nil, &pnRoot); err != nil {
		return nil, err
	}

	return pnRoot.PrivateNetworks, nil
}

// GetPrivateNetwork returns private network info
func (lb *LoadBalancersService) GetPrivateNetwork(ctx context.Context, lbID, pnID string) (*LBPrivateNetwork, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/private_networks/%s", lbID, pnID)

	req, err := lb.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var pnRoot lbPrivateNetworkRoot
	_, err = lb.client.Do(ctx, req, &pnRoot)

	if err != nil {
		return nil, err
	}

	return pnRoot.PrivateNetwork, nil
}

type connectLBPrivateNetworksRequest struct {
	PrivateNetworkIDs []string `json:"private_network_ids,omitempty"`
}

// ConnectPrivateNetworks connects LB to a private networks
func (lb *LoadBalancersService) ConnectPrivateNetworks(ctx context.Context, lbID string, pnIDs []string) ([]LBPrivateNetwork, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/private_networks", lbID)

	req, err := lb.client.newRequest(http.MethodPost, path, &connectLBPrivateNetworksRequest{PrivateNetworkIDs: pnIDs})
	if err != nil {
		return nil, err
	}

	var pnRoot lbPrivateNetworksRoot
	_, err = lb.client.Do(ctx, req, &pnRoot)

	if err != nil {
		return nil, err
	}

	return pnRoot.PrivateNetworks, nil
}

// DisconnectPrivateNetwork removes lb from the private network
func (lb *LoadBalancersService) DisconnectPrivateNetwork(ctx context.Context, lbID, pnID string) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s/private_networks/%s", lbID, pnID)

	req, err := lb.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	resp, err := lb.client.Do(ctx, req, nil)

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error disconnect from private network: %v", resp.StatusCode)
	}

	return err
}

type lbBackendNodesRoot struct {
	BackendNodes []LBBackendNode `json:"backend_nodes,omitempty"`
}

type lbBackendNodeRoot struct {
	BackendNode *LBBackendNode `json:"backend_node,omitempty"`
}

// ListBackendNodes returns all connected backend nodes
func (lb *LoadBalancersService) ListBackendNodes(ctx context.Context, lbID string) ([]LBBackendNode, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/backend_nodes", lbID)

	var bnRoot lbBackendNodesRoot

	if err := lb.client.list(ctx, path, nil, &bnRoot); err != nil {
		return nil, err
	}

	return bnRoot.BackendNodes, nil
}

// GetBackendNode returns backend node info
func (lb *LoadBalancersService) GetBackendNode(ctx context.Context, lbID, bnID string) (*LBBackendNode, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/backend_nodes/%s", lbID, bnID)

	req, err := lb.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var bnRoot lbBackendNodeRoot
	_, err = lb.client.Do(ctx, req, &bnRoot)

	if err != nil {
		return nil, err
	}

	return bnRoot.BackendNode, nil
}

type addLBBackendNodesRequest struct {
	BackendNodes []LBBackendNodeCreateRequest `json:"backend_nodes,omitempty"`
}

// AddBackendNodes connects backend nodes to the LB
func (lb *LoadBalancersService) AddBackendNodes(ctx context.Context, lbID string, bnIDs []string) ([]LBBackendNode, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/backend_nodes", lbID)

	var lbCloudServers []LBBackendNodeCreateRequest

	for _, bnID := range bnIDs {
		lbCloudServers = append(lbCloudServers, LBBackendNodeCreateRequest{bnID})
	}

	req, err := lb.client.newRequest(http.MethodPost, path, &addLBBackendNodesRequest{BackendNodes: lbCloudServers})
	if err != nil {
		return nil, err
	}

	var bnRoot lbBackendNodesRoot
	_, err = lb.client.Do(ctx, req, &bnRoot)

	if err != nil {
		return nil, err
	}

	return bnRoot.BackendNodes, nil
}

// DeleteBackendNode removes backend node from the LB
func (lb *LoadBalancersService) DeleteBackendNode(ctx context.Context, lbID, bnID string) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s/backend_nodes/%s", lbID, bnID)

	req, err := lb.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	resp, err := lb.client.Do(ctx, req, nil)

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error deleting backend node: %v", resp.StatusCode)
	}

	return err
}

type lbHealthChecksRoot struct {
	HealthChecks []LBHealthCheck `json:"health_checks,omitempty"`
}

type lbHealthCheckRoot struct {
	HealthCheck *LBHealthCheck `json:"health_check,omitempty"`
}

// ListHealthChecks returns all health checks
func (lb *LoadBalancersService) ListHealthChecks(ctx context.Context, lbID string) ([]LBHealthCheck, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/health_checks", lbID)

	var hcRoot lbHealthChecksRoot

	if err := lb.client.list(ctx, path, nil, &hcRoot); err != nil {
		return nil, err
	}

	return hcRoot.HealthChecks, nil
}

// GetHealthCheck returns health check info
func (lb *LoadBalancersService) GetHealthCheck(ctx context.Context, lbID, hcID string) (*LBHealthCheck, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/health_checks/%s", lbID, hcID)

	req, err := lb.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var hcRoot lbHealthCheckRoot
	_, err = lb.client.Do(ctx, req, &hcRoot)

	if err != nil {
		return nil, err
	}

	return hcRoot.HealthCheck, nil
}

// CreateHealthCheck creates new health check
func (lb *LoadBalancersService) CreateHealthCheck(ctx context.Context, lbID string, request *LBHealthCheckCreateRequest) (*LBHealthCheck, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/health_checks", lbID)

	req, err := lb.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var hcRoot lbHealthCheckRoot
	_, err = lb.client.Do(ctx, req, &hcRoot)

	if err != nil {
		return nil, err
	}

	return hcRoot.HealthCheck, nil
}

// LBHealthCheckUpdateRequest object
type LBHealthCheckUpdateRequest struct {
	Type               string `json:"type,omitempty"`
	URL                string `json:"url,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   int    `json:"Healthy_threshold,omitempty"`
	Port               int    `json:"port,omitempty"`
}

// UpdateHealthCheck updates the health check
func (lb *LoadBalancersService) UpdateHealthCheck(ctx context.Context, lbID, hcID string, request *LBHealthCheckUpdateRequest) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s/health_checks/%s", lbID, hcID)

	req, err := lb.client.newRequest(http.MethodPatch, path, request)
	if err != nil {
		return err
	}

	_, err = lb.client.Do(ctx, req, nil)

	if err != nil {
		return err
	}

	return nil
}

// DeleteHealthCheck removes health check from the LB
func (lb *LoadBalancersService) DeleteHealthCheck(ctx context.Context, lbID, hcID string) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s/health_checks/%s", lbID, hcID)

	req, err := lb.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	resp, err := lb.client.Do(ctx, req, nil)

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error deleting health check: %v", resp.StatusCode)
	}

	return err
}

type lbIPAddressesRoot struct {
	IPAddresses []LBIPAddress `json:"ip_addresses,omitempty"`
}

type lbIPAddressRoot struct {
	IPAddress *LBIPAddress `json:"ip_address,omitempty"`
}

// ListIPAddresses returns all ip addresses
func (lb *LoadBalancersService) ListIPAddresses(ctx context.Context, lbID string) ([]LBIPAddress, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/ip_addresses", lbID)

	var ipRoot lbIPAddressesRoot

	if err := lb.client.list(ctx, path, nil, &ipRoot); err != nil {
		return nil, err
	}

	return ipRoot.IPAddresses, nil
}

// GetIPAddress returns ip address info
func (lb *LoadBalancersService) GetIPAddress(ctx context.Context, lbID, ipID string) (*LBIPAddress, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/ip_addresses/%s", lbID, ipID)

	req, err := lb.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var ipRoot lbIPAddressRoot
	_, err = lb.client.Do(ctx, req, &ipRoot)

	if err != nil {
		return nil, err
	}

	return ipRoot.IPAddress, nil
}

type assignIPAddressRequest struct {
	IPAddressesIDS []string `json:"ip_address_ids"`
}

// AssignIPAddresses assigns ip addresses to the LB
func (lb *LoadBalancersService) AssignIPAddresses(ctx context.Context, lbID string, ipIDs []string) ([]LBIPAddress, error) {
	path := fmt.Sprintf("api/v1/load_balancers/%s/ip_addresses", lbID)

	request := &assignIPAddressRequest{IPAddressesIDS: ipIDs}

	req, err := lb.client.newRequest(http.MethodPost, path, request)
	if err != nil {
		return nil, err
	}

	var ipRoot lbIPAddressesRoot
	_, err = lb.client.Do(ctx, req, &ipRoot)

	if err != nil {
		return nil, err
	}

	return ipRoot.IPAddresses, nil
}

// ReleaseIPAddress removes ip address from the LB
func (lb *LoadBalancersService) ReleaseIPAddress(ctx context.Context, lbID, ipID string) error {
	path := fmt.Sprintf("api/v1/load_balancers/%s/ip_addresses/%s", lbID, ipID)

	req, err := lb.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	resp, err := lb.client.Do(ctx, req, nil)

	if resp.StatusCode != 202 {
		return fmt.Errorf("Error releasing ip address: %v", resp.StatusCode)
	}

	return err
}
