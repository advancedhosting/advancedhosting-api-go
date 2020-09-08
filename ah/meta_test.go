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

func TestMeta_IsLastPage(t *testing.T) {
	meta := &Meta{
		Page:    1,
		PerPage: 25,
		Total:   8,
	}

	if !meta.IsLastPage() {
		t.Fatalf("Page should be the last")
	}
}

func TestMeta_IsNotLastPage(t *testing.T) {
	meta := &Meta{
		Page:    1,
		PerPage: 25,
		Total:   30,
	}

	if meta.IsLastPage() {
		t.Fatalf("Page shouldn't be the last")
	}
}
