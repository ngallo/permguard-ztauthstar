// Copyright 2024 Nitro Agility S.r.l.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package sanitizers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestStringify tests the Stringify function.
func TestSanitizeString(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{ Input: "ab c", Output: "ab c" },
		{ Input: "AB C", Output: "ab c" },
		{ Input: "", Output: "" },
	}
	for _, test := range tests {
		output := SanitizeString(test.Input)
		assert.Equal(t, test.Output, output)
	}
}

// TestSanitizeWilcardString tests the SanitizeWilcardString function.
func TestSanitizeWilcardString(t *testing.T) {
	tests := []struct {
		Input  string
		Output string
	}{
		{ Input: "ab c", Output: "ab c" },
		{ Input: "AB C", Output: "ab c" },
		{ Input: "", Output: "*" },
	}
	for _, test := range tests {
		output := SanitizeWilcardString(test.Input)
		assert.Equal(t, test.Output, output)
	}
}
