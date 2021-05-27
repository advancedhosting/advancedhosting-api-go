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
	"strings"

	"github.com/google/go-querystring/query"
)

// ListOptions represents options to get list of resources.
type ListOptions struct {
	Meta     *ListMetaOptions
	Filters  []FilterInterface
	Sortings []*Sorting
}

// ListMetaOptions represents meta options.
type ListMetaOptions struct {
	Page int `url:"page,omitempty"`
}

func buildListQuery(options *ListOptions) string {

	var queryString []string

	if options.Meta != nil {
		metaParams, _ := query.Values(options.Meta)
		queryString = append(queryString, metaParams.Encode())
	}

	if len(options.Filters) > 0 {
		queryString = append(queryString, BuildFilterQuery(options.Filters))
	}

	if len(options.Sortings) > 0 {
		queryString = append(queryString, BuildSortingQuery(options.Sortings))
	}

	return strings.Join(queryString, "&")
}
