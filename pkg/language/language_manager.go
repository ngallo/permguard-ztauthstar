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
	"encoding/json"
	"errors"
	"fmt"

	aztypes "github.com/permguard/permguard-abs-language/pkg/language/types"
	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
)

const (
	// UUR format string: {account}:{tenant}:{domain}:{resource}:{resource-filter}.
	uurFormatString = "uur:%s:%s:%s:%s:%s"
	// AR format string: {resource}:{action}.
	arFormatString = "%s:%s"
)

// LanguageManager is the manager for policies.
type LanguageManager struct {
}

// NewLanguageManager creates a new LanguageManager.
func NewLanguageManager() *LanguageManager {
	return &LanguageManager{}
}

// sanitizeValidateOptimize sanitizes, validates and optimize the input policy.
func (pm *LanguageManager) sanitizeValidateOptimize(instance any, sanitize bool, validate bool, optimize bool) (*aztypes.Policy, error) {
	switch v := instance.(type) {
	case aztypes.Policy:
		return pm.sanitizeValidateOptimizePolicy(&v, sanitize, validate, optimize)
	}
	return nil, errors.New("authz: not implemented")
}

// UnmarshalType unmarshals a language type from the given data, and optionally sanitized, validates and optimizes it based on the provided parameters.
func (pm *LanguageManager) UnmarshalType(data []byte, sanitize bool, validate bool, optimize bool) (*aztypes.TypeInfo, error) {
	if data == nil {
		return nil, errors.New("authz: type cannot be unmarshaled from nil data")
	}
	var policy aztypes.Policy
	if err := json.Unmarshal(data, &policy); err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal type - %w", err)
	}
	snzPolicy, err := pm.sanitizeValidateOptimize(&policy, sanitize, validate, optimize)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal type - %w", err)
	}
	strfy, err := aztext.Stringify(snzPolicy, nil)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal type - %w", err)
	}
	return &aztypes.TypeInfo{
		Hash: azcrypto.ComputeStringSHA256(strfy),
		Type: snzPolicy,
	}, nil
}

// MarshalType marshals a type to a byte array, and optionally sanitized, validates and optimizes it based on the provided parameters.
func (pm *LanguageManager) MarshalType(instance any, sanitize bool, validate bool, optimize bool) ([]byte, error) {
	if instance == nil {
		return nil, errors.New("authz: type cannot be marshaled from nil instance")
	}
	snzPolicy, err := pm.sanitizeValidateOptimize(instance, sanitize, validate, optimize)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to unmarshal type - %w", err)
	}
	data, err := json.Marshal(snzPolicy)
	if err != nil {
		return nil, fmt.Errorf("authz: failed to marshal type - %w", err)
	}
	return data, nil
}
