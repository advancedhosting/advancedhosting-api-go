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
	"fmt"
	"strings"
)

// FilterInterface is an interface for Filter.
type FilterInterface interface {
	Encode() []string
}

// InFilter represents Ransack "*_in" filter .
type InFilter struct {
	Keys   []string
	Values []string
}

// Encode returns Ransack "*_in" filter expression
func (f *InFilter) Encode() []string {
	var query []string
	keys := strings.Join(f.Keys, "_or_")
	for _, value := range f.Values {
		query = append(query, fmt.Sprintf("q[%s_in][]=%s", keys, value))
	}
	return query
}

// EqFilter represents Ransack "*_eq" filter .
type EqFilter struct {
	Value string
	Keys  []string
}

// Encode returns Ransack "*_eq" filter expression
func (f *EqFilter) Encode() []string {
	keys := strings.Join(f.Keys, "_or_")
	return []string{fmt.Sprintf("q[%s_eq]=%s", keys, f.Value)}
}

// ContFilter represents Ransack "*_cont" filter .
type ContFilter struct {
	Value string
	Keys  []string
}

// Encode returns Ransack "*_eq" filter expression
func (f *ContFilter) Encode() []string {
	keys := strings.Join(f.Keys, "_or_")
	return []string{fmt.Sprintf("q[%s_cont]=%s", keys, f.Value)}
}

// BuildFilterQuery returns Ransack filter expression
func BuildFilterQuery(filters []FilterInterface) string {
	var query []string
	for _, filter := range filters {
		query = append(query, filter.Encode()...)
	}
	return strings.Join(query, "&")
}

// Sorting represents Ransack sorting expression
type Sorting struct {
	Key   string
	Order string
}

// Encode returns Ransack sorting expression
func (s *Sorting) Encode() []string {
	return []string{fmt.Sprintf("q[s]=%s+%s", s.Key, s.Order)}
}

// BuildSortingQuery returns Ransack sorting expression
func BuildSortingQuery(sortings []*Sorting) string {
	var query []string
	for _, sorting := range sortings {
		query = append(query, sorting.Encode()...)
	}
	return strings.Join(query, "&")
}
