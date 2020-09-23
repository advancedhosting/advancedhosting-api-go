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
