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

package policies

import (
	"strings"
)

// sanitizeString sanitizes a string.
func (pm *policyManager) sanitizeString(value string) string {
	return strings.ToLower(value)
}

// sanitizeWilcardString sanitizes a wildcard string.
func (pm *policyManager) sanitizeWilcardString(value string) string {
	sanitizedValue := pm.sanitizeString(value)
	if len(value) == 0 {
		sanitizedValue = "*"
	}
	return sanitizedValue
}

// sanitizePolicy sanitizes a policy.
func (pm *policyManager) sanitizePolicy(policy *Policy) (*Policy, error) {
	policy.SyntaxVersion = pm.sanitizeString(policy.SyntaxVersion)
	policy.Type = pm.sanitizeString(policy.Type)
	policy.Name = pm.sanitizeString(policy.Name)
	policy.Resource = UURString(pm.sanitizeWilcardString(string(policy.Resource)))
	for i, action := range policy.Actions {
		policy.Actions[i] = ARString(pm.sanitizeWilcardString(string(action)))
	}
	return policy, nil
}
