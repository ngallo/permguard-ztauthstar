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
	"encoding/json"
	"errors"
	"fmt"

	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
)

const (
	// UUR format string: {account}:{tenant}:{domain}:{resource}:{resource-filter}.
	uurFormatString = "uur:%s:%s:%s:%s:%s"
	// AR format string: {resource}:{action}.
	arFormatString = "%s:%s"
)

// policyManager is the manager for policies.
type policyManager struct {
}

// newPolicyManager creates a new PolicyManager.
func newPolicyManager() *policyManager {
	return &policyManager{}
}

// sanitizeValidateOptimize sanitizes, validates and optimize the input policy.
func (pm *policyManager) sanitizeValidateOptimize(policy *Policy, sanitize bool, validate bool, optimize bool) (*Policy, error) {
	var err error
	targetPolicy := policy
	if sanitize {
		targetPolicy, err = pm.sanitizePolicy(policy)
		if err != nil {
			return nil, err
		}
	}
	if validate {
		valid, err := pm.validate(targetPolicy)
		if !valid {
			return nil, errors.New("authz: policy is invalid")
		}
		if err != nil {
			return nil, err
		}
	}
	return targetPolicy, nil
}

// UnmarshalPolicy unmarshals a policy from the given data, and optionally sanitized, validates and optimizes it based on the provided parameters.
func (pm *policyManager) UnmarshalPolicy(data []byte, sanitize bool, validate bool, optimize bool) (*PolicyInfo, error) {
	if data == nil {
		return nil, errors.New("authz: policy cannot be unmarshaled from nil data")
	}
	var policy Policy
	if err := json.Unmarshal(data, &policy); err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal policy - %w", err)
	}
	snzPolicy, err := pm.sanitizeValidateOptimize(&policy, sanitize, validate, optimize)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal policy - %w", err)
	}
	strfy, err := aztext.Stringify(snzPolicy, nil)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal policy - %w", err)
	}
	return &PolicyInfo{
		PolicyHash: azcrypto.ComputeStringSHA1(strfy),
		Policy:     snzPolicy,
	}, nil
}

// MarshalPolicy marshals a policy to a byte array, and optionally sanitized, validates and optimizes it based on the provided parameters.
func (pm *policyManager) MarshalPolicy(policy *Policy, sanitize bool, validate bool, optimize bool) ([]byte, error) {
	if policy == nil {
		return nil, errors.New("authz: policy cannot be marshaled from nil policy")
	}
	snzPolicy, err := pm.sanitizeValidateOptimize(policy, sanitize, validate, optimize)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal policy - %w", err)
	}
	data, err := json.Marshal(snzPolicy)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to marshal policy - %w", err)
	}
	return data, nil
}
