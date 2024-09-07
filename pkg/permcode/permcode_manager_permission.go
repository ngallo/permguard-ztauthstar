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
	"sort"

	azsanitizers "github.com/permguard/permguard-abs-language/pkg/permcode/sanitizers"
	aztypes "github.com/permguard/permguard-abs-language/pkg/permcode/types"
	azvalidators "github.com/permguard/permguard-core/pkg/extensions/validators"
)

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
		if !valid {
			return nil, errors.New("permcode: permission is invalid")
		}
		if err != nil {
			return nil, err
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
	for i, reference := range permission.PolicyReferences {
		permission.PolicyReferences[i] = azsanitizers.SanitizeString(reference)
	}
	return permission, nil
}

// validatePermission validates the input permission.
func (pm *PermCodeManager) validatePermission(permission *aztypes.Permission) (bool, error) {
	if permission.SyntaxVersion != aztypes.PolicySyntax || permission.Type != aztypes.ClassTypeACPermission {
		return false, nil
	}
	if !azvalidators.ValidateName(permission.Name) {
		return false, errors.New("permcode: invalid name")
	}
	for _, reference := range permission.PolicyReferences {
		if !azvalidators.ValidateName(reference) {
			return false, errors.New("permcode: invalid policy reference name")
		}
	}
	return true, nil
}

// optimizePermission optimizes a permission.
func (pm *PermCodeManager) optimizePermission(permission *aztypes.Permission) (*aztypes.Permission, error) {
	policySet := make(map[string]struct{})
	uniquePolicies := []string{}
	for _, policy := range permission.PolicyReferences {
		if _, exists := policySet[policy]; !exists {
			policySet[policy] = struct{}{}
			uniquePolicies = append(uniquePolicies, policy)
		}
	}
	sort.Strings(uniquePolicies)
	permission.PolicyReferences = uniquePolicies
	return permission, nil
}
