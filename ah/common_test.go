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
	"testing"
)

func TestListOptions_Basic(t *testing.T) {
	options := &ListOptions{
		Meta: &ListMetaOptions{
			Page: 1,
		},
		Filters: []FilterInterface{
			&InFilter{
				Keys:   []string{"test"},
				Values: []string{"1", "3"},
			},
		},
		Sortings: []*Sorting{
			{
				Key:   "test2",
				Order: "asc",
			},
		},
	}
	expectedResult := "page=1&q[test_in][]=1&q[test_in][]=3&q[s]=test2+asc"
	result := buildListQuery(options)

	if result != expectedResult {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedResult, result)
	}

}

func TestListOptions_OnlyMeta(t *testing.T) {
	options := &ListOptions{
		Meta: &ListMetaOptions{
			Page: 1,
		},
	}
	expectedResult := "page=1"
	result := buildListQuery(options)

	if result != expectedResult {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedResult, result)
	}
}

func TestListOptions_Empty(t *testing.T) {
	options := &ListOptions{}
	expectedResult := ""
	result := buildListQuery(options)

	if result != expectedResult {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedResult, result)
	}
}

func TestListOptions_OnlyFilters(t *testing.T) {
	options := &ListOptions{
		Filters: []FilterInterface{
			&InFilter{
				Keys:   []string{"test"},
				Values: []string{"1", "3"},
			},
		},
	}
	expectedResult := "q[test_in][]=1&q[test_in][]=3"
	result := buildListQuery(options)

	if result != expectedResult {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedResult, result)
	}
}

func TestListOptions_OnlySorting(t *testing.T) {
	options := &ListOptions{
		Sortings: []*Sorting{
			{
				Key:   "test2",
				Order: "asc",
			},
		},
	}
	expectedResult := "q[s]=test2+asc"
	result := buildListQuery(options)

	if result != expectedResult {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedResult, result)
	}
}
