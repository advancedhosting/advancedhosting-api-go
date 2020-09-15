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

// Action object
type Action struct {
	ID           string `json:"id,omitempty"`
	State        string `json:"state,omitempty"`
	ResourceID   string `json:"resource_id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	ResourceType string `json:"resource_type,omitempty"`
	Type         string `json:"type,omitempty"`
	UserID       string `json:"user_id,omitempty"`
	Note         string `json:"note,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	StartedAt    string `json:"started_at,omitempty"`
	CompletedAt  string `json:"completed_at,omitempty"`
}

type actionRoot struct {
	Action *Action `json:"action"`
}

type actionsRoot struct {
	Actions []Action `json:"actions"`
}
