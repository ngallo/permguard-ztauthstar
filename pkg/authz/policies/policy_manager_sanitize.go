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
	azsanitizers "github.com/permguard/permguard-abs-language/pkg/extensions/sanitizers"
)

// sanitizePolicy sanitizes a policy.
func (pm *policyManager) sanitizePolicy(policy *Policy) (*Policy, error) {
	policy.SyntaxVersion = azsanitizers.SanitizeString(policy.SyntaxVersion)
	policy.Type = azsanitizers.SanitizeString(policy.Type)
	policy.Name = azsanitizers.SanitizeString(policy.Name)
	policy.Resource = UURString(azsanitizers.SanitizeWilcardString(string(policy.Resource)))
	for i, action := range policy.Actions {
		policy.Actions[i] = ARString(azsanitizers.SanitizeWilcardString(string(action)))
	}
	return policy, nil
}
