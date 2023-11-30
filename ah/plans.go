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

// PlanPrice object
type PlanPrice struct {
	Type     string `json:"type,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Currency string `json:"currency,omitempty"`
	Quantity string `json:"quantity,omitempty"`
	Price    string `json:"price,omitempty"`
	ID       int    `json:"id,omitempty"`
	PlanID   int    `json:"plan_id,omitempty"`
}

// Plan object
type Plan struct {
	Prices   map[int]PlanPrice `json:"prices,omitempty"`
	Type     string            `json:"type,omitempty"`
	Currency string            `json:"currency,omitempty"`
	Name     string            `json:"name,omitempty"`
	ID       int               `json:"id"`
}
