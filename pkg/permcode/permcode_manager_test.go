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
			Path: "./testdata/mashaling",
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
				pm := NewPermCodeManager()

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

// TestMashalingOfPoliciesWithArgumentsErrors tests argument error cases for marshalling and unmarshalling.
func TestMashalingOfPoliciesWithArgumentsErrors(t *testing.T) {
	t.Run("nil instance in MarshalClass", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		result, err := pm.MarshalClass(nil, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})

	t.Run("invalid JSON structure in UnmarshalClass", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		jsonStr := `{"id":"1", "color":"red"}`
		jsonBytes := []byte(jsonStr)
		result, err := pm.UnmarshalClass(jsonBytes, aztypes.ClassTypeACPolicy, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})

	t.Run("unmarshal class with incorrect syntax version", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		jsonStr := `{"syntax":"invalidSyntax", "type":"acpolicy"}`
		jsonBytes := []byte(jsonStr)
		result, err := pm.UnmarshalClass(jsonBytes, aztypes.ClassTypeACPolicy, false, false, false)
		assert.NotNil(err, "Expected error for invalid syntax version")
		assert.Nil(result)
	})

	t.Run("unmarshal class with missing class type", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		jsonStr := `{"syntax":"permguard1"}`
		jsonBytes := []byte(jsonStr)
		result, err := pm.UnmarshalClass(jsonBytes, aztypes.ClassTypeACPolicy, false, false, false)
		assert.NotNil(err, "Expected error for missing class type")
		assert.Nil(result)
	})

	t.Run("marshal invalid object type", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		obj := "invalid object"
		result, err := pm.MarshalClass(obj, false, false, false)
		assert.NotNil(err)
		assert.Nil(result)
	})
}

// TestMashalingOfPoliciesWithErrors tests the error cases during policy marshaling/unmarshaling.
func TestMashalingOfPoliciesWithErrors(t *testing.T) {
	tests := []struct {
		Path string
	}{
		{
			Path: "./testdata/mashaling-with-errors",
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
				pm := NewPermCodeManager()

				sanitize := data["sanitize"].(bool)
				validate := data["validate"].(bool)
				optimize := data["optimize"].(bool)

				inputData, _ := json.Marshal(data["input"])
				policyInInfo, err := pm.UnmarshalClass(inputData, aztypes.ClassTypeACPolicy, sanitize, validate, optimize)

				assert.NotNil(err, "Expected error during unmarshaling")
				assert.Nil(policyInInfo, "Policy info should be nil on error")
			})
		}
	}
}

// TestPermCodeManagerWithOptimize tests the optimization behavior in the PermCodeManager.
func TestPermCodeManagerWithOptimize(t *testing.T) {
	t.Run("Test optimization on invalid policy 1", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		validPolicy := aztypes.Policy{Name: "valid-policy"}
		_, err := pm.MarshalClass(&validPolicy, true, true, true)

		assert.NotNil(err, "Optimization should not cause an error")
	})

	t.Run("Test optimization with invalid policy 2", func(t *testing.T) {
		assert := assert.New(t)
		pm := NewPermCodeManager()

		invalidPolicy := aztypes.Policy{Name: "invalid policy name @#"}
		result, err := pm.MarshalClass(&invalidPolicy, true, true, true)

		assert.NotNil(err, "Expected error during optimization of invalid policy")
		assert.Nil(result)
	})
}
