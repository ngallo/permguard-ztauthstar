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

// TestMashalingOfPermissions tests the marshalling of permissions.
func TestMashalingOfPermissions(t *testing.T) {
	tests := []struct {
		Path string
	}{
		{
			Path: "./testdata/permission-mashaling",
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
				permissionInInfo, err := pm.UnmarshalClass(inputData, aztypes.ClassTypeACPermission, sanitize, validate, optimize)
				assert.Nil(err, "UnmarshalClass should not return an error")
				permissionInData, err := pm.MarshalClass(permissionInInfo.Instance, sanitize, validate, optimize)
				assert.Nil(err, "MarshalClass should not return an error")
				permissionInDataSha := azcrypto.ComputeSHA256(permissionInData)

				outputData, err := json.Marshal(data["output"])
				permissionOutInfo, _ := pm.UnmarshalClass(outputData, aztypes.ClassTypeACPermission, false, false, false)
				permissionOutData, _ := pm.MarshalClass(permissionOutInfo.Instance, false, false, false)
				permissionOutDataSha := azcrypto.ComputeSHA256(permissionOutData)

				assert.Nil(err)
				assert.Equal(permissionInDataSha, permissionOutDataSha, "Input and output SHA256 hashes should match")
			})
		}
	}
}
