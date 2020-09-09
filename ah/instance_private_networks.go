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

// InstancePrivateNetwork object
type InstancePrivateNetwork struct {
	ID          string                      `json:"id,omitempty"`
	IP          string                      `json:"ip"`
	MACAddress  string                      `json:"mac_address"`
	State       string                      `json:"state,omitempty"`
	ConnectedAt string                      `json:"connected_at,omitempty"`
	Instance    *PrivateNetworkInstanceInfo `json:"instance,omitempty"`
}

// PrivateNetworkInstanceInfo object
type PrivateNetworkInstanceInfo struct {
	ID      string `json:"id,omitempty"`
	ImageID string `json:"image_id"`
	Name    string `json:"name"`
	Number  string `json:"number,omitempty"`
}

// PrivateNetworkInfo object
type PrivateNetworkInfo struct {
	*PrivateNetwork
	InstancePrivateNetworks []struct {
		*InstancePrivateNetwork
	} `json:"instance_private_networks,omitempty"`
}
