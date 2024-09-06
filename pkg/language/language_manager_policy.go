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

package language

import (
	"errors"
	"strings"

	azsanitizers "github.com/permguard/permguard-abs-language/pkg/extensions/sanitizers"
	aztypes "github.com/permguard/permguard-abs-language/pkg/language/types"
	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
	azvalidators "github.com/permguard/permguard-core/pkg/extensions/validators"
)

// sanitizeValidateOptimize sanitizes, validates and optimize the input policy.
func (pm *LanguageManager) sanitizeValidateOptimizePolicy(policy *aztypes.Policy, sanitize bool, validate bool, optimize bool) (*aztypes.Policy, error) {
	var err error
	targetPolicy := policy
	if sanitize {
		targetPolicy, err = pm.sanitizePolicy(policy)
		if err != nil {
			return nil, err
		}
	}
	if validate {
		valid, err := pm.validatePolicy(targetPolicy)
		if !valid {
			return nil, errors.New("authz: policy is invalid")
		}
		if err != nil {
			return nil, err
		}
	}
	if optimize {
		targetPolicy, err = pm.optimizePolicy(targetPolicy)
		if err != nil {
			return nil, err
		}
	}
	return targetPolicy, nil
}

// sanitizePolicy sanitizes a policy.
func (pm *LanguageManager) sanitizePolicy(policy *aztypes.Policy) (*aztypes.Policy, error) {
	policy.SyntaxVersion = azsanitizers.SanitizeString(policy.SyntaxVersion)
	policy.Type = azsanitizers.SanitizeString(policy.Type)
	policy.Name = azsanitizers.SanitizeString(policy.Name)
	for i, action := range policy.Actions {
		policy.Actions[i] = aztypes.ARString(azsanitizers.SanitizeWilcardString(string(action)))
		ar, err := policy.Actions[i].Prase()
		if err != nil {
			return nil, err
		}
		ar.Resource = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(ar.Resource)))
		ar.Action = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(ar.Action)))
		policy.Actions[i] = aztypes.FormatARString(ar.Resource, ar.Action)
	}
	resource, err := policy.Resource.Prase()
	if err != nil {
		return nil, err
	}
	resource.Domain = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(resource.Domain)))
	resource.Tenant = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(resource.Tenant)))
	resource.Resource = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(resource.Resource)))
	for i := range resource.ResourceFilter {
		resource.ResourceFilter[i] = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(resource.ResourceFilter[i])))
	}
	policy.Resource = aztypes.FormatUURString(resource.Account, resource.Tenant, resource.Domain, resource.Resource, resource.ResourceFilter)
	return policy, nil
}

// validatePolicy validates the input policy.
func (pm *LanguageManager) validatePolicy(policy *aztypes.Policy) (bool, error) {
	if policy.SyntaxVersion != aztypes.PolicySyntax || policy.Type != aztypes.ACPolicyType {
		return false, nil
	}
	if !azvalidators.ValidateName(policy.Name) {
		return false, errors.New("authz: invalid name")
	}
	for _, action := range policy.Actions {
		ar, err := action.Prase()
		if err != nil {
			return false, err
		}
		if !azvalidators.ValidateWildcardName(string(ar.Resource)) {
			return false, errors.New("authz: invalid resource")
		}
		if !azvalidators.ValidateWildcardName(string(ar.Action)) {
			return false, errors.New("authz: invalid action")
		}
	}
	uur, err := policy.Resource.Prase()
	if err != nil {
		return false, err
	}
	if !azvalidators.ValidateAccountID(uur.Account) {
		return false, errors.New("authz: invalid account id")
	}
	if !azvalidators.ValidateWildcardName(string(uur.Tenant)) {
		return false, errors.New("authz: invalid tenant")
	}
	if !azvalidators.ValidateWildcardName(string(uur.Domain)) {
		return false, errors.New("authz: invalid domain")
	}
	if !azvalidators.ValidateWildcardName(string(uur.Resource)) {
		return false, errors.New("authz: invalid resource")
	}
	for _, filter := range uur.ResourceFilter {
		filterStr := string(filter)
		if filterStr == "" || strings.Contains(filterStr, " ") {
			return false, errors.New("authz: invalid resource filter")
		}
	}
	return true, nil
}

// optimizePolicy optimizes a policy.
func (pm *LanguageManager) optimizePolicy(policy *aztypes.Policy) (*aztypes.Policy, error) {
	return policy, nil
}
