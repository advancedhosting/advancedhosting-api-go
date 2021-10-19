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

// PlanPrice object
type PlanPrice struct {
	Type     string `json:"type"`
	Unit     string `json:"unit"`
	Currency string `json:"currency"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Id       int    `json:"id"`
	PlanId   int    `json:"plan_id"`
}

// Plan object
type Plan struct {
	Prices           map[int]PlanPrice `json:"prices"`
	Type             string            `json:"type"`
	Currency         string            `json:"currency"`
	Name             string            `json:"name"`
	CustomAttributes struct {
		Ram              string `json:"ram"`
		Disk             string `json:"disk"`
		Slug             string `json:"slug"`
		Vcpu             string `json:"vcpu"`
		Traffic          string `json:"traffic"`
		WebsaProductId   string `json:"websaProductId"`
		Hot              bool   `json:"hot"`
		AvailableOnTrial bool   `json:"available_on_trial"`
	} `json:"custom_attributes"`
	Id int `json:"id"`
}
