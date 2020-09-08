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

import (
	"testing"
)

func TestFilterIn_Basic(t *testing.T) {
	filters := []FilterInterface{
		&InFilter{
			Keys:   []string{"test"},
			Values: []string{"1", "2"},
		},
	}

	expectedQuery := "q[test_in][]=1&q[test_in][]=2"
	query := BuildFilterQuery(filters)

	if query != expectedQuery {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedQuery, query)
	}
}
func TestFilterIn_MultipleKey(t *testing.T) {
	filters := []FilterInterface{
		&InFilter{
			Keys:   []string{"test", "test2"},
			Values: []string{"1", "2"},
		},
	}

	expectedQuery := "q[test_or_test2_in][]=1&q[test_or_test2_in][]=2"
	query := BuildFilterQuery(filters)

	if query != expectedQuery {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedQuery, query)
	}
}

func TestSorting_Basic(t *testing.T) {
	sortings := []*Sorting{
		{
			Key:   "test",
			Order: "asc",
		},
		{
			Key:   "test2",
			Order: "desc",
		},
	}

	expectedQuery := "q[s]=test+asc&q[s]=test2+desc"
	query := BuildSortingQuery(sortings)

	if query != expectedQuery {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedQuery, query)
	}
}

func TestFilterEq_Basic(t *testing.T) {
	filters := []FilterInterface{
		&EqFilter{
			Keys:  []string{"test"},
			Value: "1",
		},
	}

	expectedQuery := "q[test_eq]=1"
	query := BuildFilterQuery(filters)

	if query != expectedQuery {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedQuery, query)
	}
}

func TestFilterCont_Basic(t *testing.T) {
	filters := []FilterInterface{
		&ContFilter{
			Keys:  []string{"test"},
			Value: "1",
		},
	}

	expectedQuery := "q[test_cont]=1"
	query := BuildFilterQuery(filters)

	if query != expectedQuery {
		t.Fatalf("Wrong query. Expected %s, got %s", expectedQuery, query)
	}
}
