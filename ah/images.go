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
)

// Image object
type Image struct {
	ID           string `json:"id,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	Name         string `json:"name,omitempty"`
	Distribution string `json:"distribution,omitempty"`
	Version      string `json:"version,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Slug         string `json:"slug,omitempty"`
	Public       bool   `json:"public,omitempty"`
}

// ImagesAPI is an interface for images.
type ImagesAPI interface {
	List(context.Context, *ListOptions) ([]Image, *Meta, error)
}

// ImagesService implements ImagesAPI interface.
type ImagesService struct {
	client *APIClient
}

type imagesRoot struct {
	Images []Image `json:"images"`
	Meta   *Meta   `json:"meta"`
}

// List returns all available images
func (is *ImagesService) List(ctx context.Context, options *ListOptions) ([]Image, *Meta, error) {

	path := "api/v1/images"

	var iRoot imagesRoot

	if err := is.client.list(ctx, path, options, &iRoot); err != nil {
		return nil, nil, err
	}
	return iRoot.Images, iRoot.Meta, nil
}
