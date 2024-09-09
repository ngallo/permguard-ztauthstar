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

package permcode

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	aztypes "github.com/permguard/permguard-abs-language/pkg/permcode/types"
	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
)

// TestMashalingOfPolicies tests the marshalling of policies.
func TestMashalingOfPolicies(t *testing.T) {
	tests := []struct {
		Path string
	}{
		{
			Path: "./testdata/policy-mashaling",
		},
	}
	for _, test := range tests {
		testCases, _ := os.ReadDir(test.Path)
		for _, testCase := range testCases {
			testCaseFile := filepath.Join(test.Path, testCase.Name())
			testCaseData, _ := os.ReadFile(testCaseFile)
			var data map[string]any
			json.Unmarshal(testCaseData, &data)
			t.Run(data["testcase"].(string), func(t *testing.T) {
				assert := assert.New(t)
				pm, _ := NewPermCodeManager()

				sanitize := data["sanitize"].(bool)
				validate := data["validate"].(bool)
				optimize := data["optimize"].(bool)

				inputData, _ := json.Marshal(data["input"])
				policyInInfo, err := pm.UnmarshalClass(inputData, aztypes.ClassTypeACPolicy, sanitize, validate, optimize)
				assert.Nil(err, "UnmarshalClass should not return an error")
				policyInData, err := pm.MarshalClass(policyInInfo.Instance, sanitize, validate, optimize)
				assert.Nil(err, "MarshalClass should not return an error")
				policyInDataSha := azcrypto.ComputeSHA256(policyInData)

				outputData, err := json.Marshal(data["output"])
				policyOutInfo, _ := pm.UnmarshalClass(outputData, aztypes.ClassTypeACPolicy, false, false, false)
				policyOutData, _ := pm.MarshalClass(policyOutInfo.Instance, false, false, false)
				policyOutDataSha := azcrypto.ComputeSHA256(policyOutData)

				assert.Nil(err)
				assert.Equal(policyInDataSha, policyOutDataSha, "Input and output SHA256 hashes should match")
			})
		}
	}
}
