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
	"sort"
	"strconv"
	"strings"

	azsanitizers "github.com/permguard/permguard-abs-language/pkg/language/sanitizers"
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
			return nil, errors.New("language: policy is invalid")
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
	resource.Account = strings.ReplaceAll(resource.Account, " ", "")
	if resource.Account == "" {
		resource.Account = aztypes.KeywordAccount
	}
	resource.Tenant = aztext.WildcardString(strings.ReplaceAll(string(resource.Tenant), " ", ""))
	if resource.Tenant == "" {
		resource.Tenant = aztypes.KeywordTenant
	}
	resource.Domain = aztext.WildcardString(azsanitizers.SanitizeWilcardString(string(resource.Domain)))
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
		return false, errors.New("language: invalid name")
	}
	for _, action := range policy.Actions {
		ar, err := action.Prase()
		if err != nil {
			return false, err
		}
		if !azvalidators.ValidateWildcardName(string(ar.Resource)) {
			return false, errors.New("language: invalid resource")
		}
		if !azvalidators.ValidateWildcardName(string(ar.Action)) {
			return false, errors.New("language: invalid action")
		}
	}
	uur, err := policy.Resource.Prase()
	if err != nil {
		return false, err
	}
	if uur.Account != aztypes.KeywordAccount {
		account, err := strconv.ParseInt(uur.Account, 10, 64)
		if err != nil {
			return false, errors.New("language: invalid account number, must be an integer")
		}
		if !azvalidators.ValidateAccountID(account) {
			return false, errors.New("language: invalid account id")
		}
	}
	if uur.Tenant != aztypes.KeywordTenant {
		if !azvalidators.ValidateWildcardName(string(uur.Tenant)) {
			return false, errors.New("language: invalid tenant")
		}
	}
	if !azvalidators.ValidateWildcardName(string(uur.Domain)) {
		return false, errors.New("language: invalid domain")
	}
	if !azvalidators.ValidateWildcardName(string(uur.Resource)) {
		return false, errors.New("language: invalid resource")
	}
	for _, filter := range uur.ResourceFilter {
		filterStr := string(filter)
		if filterStr == "" || strings.Contains(filterStr, " ") {
			return false, errors.New("language: invalid resource filter")
		}
	}
	return true, nil
}

// removeDuplicates removes duplicate actions.
func (pm *LanguageManager) removeARStringsDuplicates(actions []aztypes.ARString, compare func(a, b aztypes.ARString) bool) []aztypes.ARString {
	for i := 0; i < len(actions); i++ {
		for j := 0; j < len(actions); j++ {
			if i != j && compare(actions[i], actions[j]) {
				actions = append(actions[:j], actions[j+1:]...)
				if j < i {
					i--
				}
				j--
			}
		}
	}
	stringActions := make([]string, len(actions))
	for i, action := range actions {
		stringActions[i] = string(action)
	}
	sort.Strings(stringActions)
	sortedActions := make([]aztypes.ARString, len(stringActions))
	for i, action := range stringActions {
		sortedActions[i] = aztypes.ARString(action)
	}
	return sortedActions
}

// optimizePolicy optimizes a policy.
func (pm *LanguageManager) optimizePolicy(policy *aztypes.Policy) (*aztypes.Policy, error) {
	uur, err := policy.Resource.Prase()
	if err != nil {
		return nil, err
	}
	seen := make(map[aztypes.ARString]bool)
	uniqueActions := []aztypes.ARString{}
	for _, action := range policy.Actions {
		rn, err := action.Prase()
		if err != nil {
			return nil, err
		}
		if rn.Resource != uur.Resource {
			continue
		}
		if !seen[action] {
			seen[action] = true
			uniqueActions = append(uniqueActions, action)
		}
	}
	uniqueActions = pm.removeARStringsDuplicates(uniqueActions, func(a, b aztypes.ARString) bool {
		x := aztext.WildcardString(a)
		y := aztext.WildcardString(b)
		return x.WildcardInclude(string(y)) && !y.WildcardInclude(string(x))
	})
	policy.Actions = uniqueActions
	return policy, nil
}
