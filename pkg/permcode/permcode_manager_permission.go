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
	"errors"
	"fmt"
	"sort"

	azsanitizers "github.com/permguard/permguard-abs-language/pkg/permcode/sanitizers"
	aztypes "github.com/permguard/permguard-abs-language/pkg/permcode/types"
	azvalidators "github.com/permguard/permguard-core/pkg/extensions/validators"
)

// UnmarshalPermission unmarshals a input byte array to a permission instance.
func (pm *PermCodeManager) UnmarshalPermission(data []byte, sanitize bool, validate bool, optimize bool) (*aztypes.PermissionInfo, error) {
	clasInfo, err := pm.UnmarshalClass(data, aztypes.ClassTypeACPermission, sanitize, validate, optimize)
	if err != nil {
		return nil, err
	}
	return &aztypes.PermissionInfo{
		SID:		clasInfo.SID,
		Permission: clasInfo.Instance.(*aztypes.Permission),
	}, nil
}

// sanitizeValidateOptimize sanitizes, validates and optimize the input permission.
func (pm *PermCodeManager) sanitizeValidateOptimizePermission(permission *aztypes.Permission, sanitize bool, validate bool, optimize bool) (*aztypes.Permission, error) {
	var err error
	targetPermission := permission
	if sanitize {
		targetPermission, err = pm.sanitizePermission(permission)
		if err != nil {
			return nil, err
		}
	}
	if validate {
		valid, err := pm.validatePermission(targetPermission)
		if err != nil {
			return nil, err
		}
		if !valid {
			return nil, errors.New("permcode: permission is invalid")
		}
	}
	if optimize {
		targetPermission, err = pm.optimizePermission(targetPermission)
		if err != nil {
			return nil, err
		}
	}
	return targetPermission, nil
}

// sanitizePermission sanitizes a permission.
func (pm *PermCodeManager) sanitizePermission(permission *aztypes.Permission) (*aztypes.Permission, error) {
	permission.SyntaxVersion = azsanitizers.SanitizeString(permission.SyntaxVersion)
	permission.Type = azsanitizers.SanitizeString(permission.Type)
	permission.Name = azsanitizers.SanitizeString(permission.Name)
	sanitizeSlice := func(slice *[]string) {
		if *slice == nil {
			*slice = []string{}
		}
		for i, policyName := range *slice {
			(*slice)[i] = azsanitizers.SanitizeString(policyName)
		}
	}
	if permission.Permit == nil {
		permission.Permit = []string{}
	}
	sanitizeSlice(&permission.Permit)
	if permission.Forbid == nil {
		permission.Forbid = []string{}
	}
	sanitizeSlice(&permission.Forbid)
	return permission, nil
}

// validatePermission validates the input permission.
func (pm *PermCodeManager) validatePermission(permission *aztypes.Permission) (bool, error) {
	if permission.SyntaxVersion != aztypes.PolicySyntax {
		return false, fmt.Errorf("permcode: invalid policy syntax (%s)", permission.SyntaxVersion)
	}
	if permission.Type != aztypes.ClassTypeACPermission {
		return false, fmt.Errorf("permcode: invalid type (%s)", permission.Type)

	}
	if !azvalidators.ValidateName(permission.Name) {
		return false, fmt.Errorf("permcode: invalid name (%s)", permission.Name)
	}
	validateSlice := func(slice []string, sliceType string) error {
		for _, policyName := range slice {
			if !azvalidators.ValidateName(policyName) {
				return fmt.Errorf("permcode: invalid %s policy name (%s)", sliceType, policyName)
			}
		}
		return nil
	}
	if permission.Permit == nil {
		permission.Permit = []string{}
	}
	if err := validateSlice(permission.Permit, "permit"); err != nil {
		return false, err
	}
	if permission.Forbid == nil {
		permission.Forbid = []string{}
	}
	if err := validateSlice(permission.Forbid, "forbid"); err != nil {
		return false, err
	}
	return true, nil
}

// optimizePermission optimizes a permission.
func (pm *PermCodeManager) optimizePermission(permission *aztypes.Permission) (*aztypes.Permission, error) {
	policySet := make(map[string]struct{})
	optimizeSlice := func(slice []string) []string {
		uniquePolicies := []string{}
		for _, policy := range slice {
			if _, exists := policySet[policy]; !exists {
				policySet[policy] = struct{}{}
				uniquePolicies = append(uniquePolicies, policy)
			}
		}
		sort.Strings(uniquePolicies)
		return uniquePolicies
	}
	if permission.Permit == nil {
		permission.Permit = []string{}
	}
	permission.Permit = optimizeSlice(permission.Permit)
	if permission.Forbid == nil {
		permission.Forbid = []string{}
	}
	permission.Forbid = optimizeSlice(permission.Forbid)
	return permission, nil
}
