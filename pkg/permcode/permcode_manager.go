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
	"errors"
	"fmt"

	azsanitizers "github.com/permguard/permguard-abs-language/pkg/permcode/sanitizers"
	aztypes "github.com/permguard/permguard-abs-language/pkg/permcode/types"
	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
)

const (
	// errMessageUnmarshalClass is the error message for unmarshaling a class.
	errMessageUnmarshalClass = "permcode: failed to unmarshal class - %w"
	// errMessageMarshalClass is the error message for marshaling a class.
	errMessageMarshalClass = "permcode: failed to marshal class - %w"
)

// PermCodeManager is the manager for policies.
type PermCodeManager struct {
}

// NewPermCodeManager creates a new PermCodeManager.
func NewPermCodeManager() *PermCodeManager {
	return &PermCodeManager{}
}

// sanitizeValidateOptimize sanitizes, validates and optimize the input policy.
func (pm *PermCodeManager) sanitizeValidateOptimize(instance any, sanitize bool, validate bool, optimize bool) (any, error) {
	switch v := instance.(type) {
	case *aztypes.Policy:
		return pm.sanitizeValidateOptimizePolicy(v, sanitize, validate, optimize)
	case *aztypes.Permission:
		return pm.sanitizeValidateOptimizePermission(v, sanitize, validate, optimize)
	}
	return nil, errors.New("permcode: not implemented")
}

// UnmarshalClass unmarshals a input byte array to a class instance.
func (pm *PermCodeManager) UnmarshalClass(data []byte, classType string, sanitize bool, validate bool, optimize bool) (*aztypes.ClassInfo, error) {
	if data == nil {
		return nil, errors.New("permcode: type cannot be unmarshaled from nil data")
	}
	var class aztypes.Class
	if err := json.Unmarshal(data, &class); err != nil {
		return nil, fmt.Errorf(errMessageUnmarshalClass, err)
	}
	class.SyntaxVersion = azsanitizers.SanitizeString(class.SyntaxVersion)
	class.Type = azsanitizers.SanitizeString(class.Type)
	if class.SyntaxVersion != aztypes.PolicySyntax {
		return nil, fmt.Errorf("permcode: failed to unmarshal type - invalid syntax version %s", class.SyntaxVersion)
	}
	if class.Type != classType {
		return nil, fmt.Errorf("permcode: failed to unmarshal type - invalid type %s", class.Type)
	}
	var classInstance any
	switch class.Type {
	case aztypes.ClassTypeACPolicy:
		var policy aztypes.Policy
		if err := json.Unmarshal(data, &policy); err != nil {
			return nil, fmt.Errorf(errMessageUnmarshalClass, err)
		}
		snzPolicy, err := pm.sanitizeValidateOptimize(&policy, sanitize, validate, optimize)
		classInstance = snzPolicy
		classType = aztypes.ClassTypeACPolicy
		if err != nil {
			return nil, fmt.Errorf(errMessageUnmarshalClass, err)
		}
	case aztypes.ClassTypeACPermission:
		var permission aztypes.Permission
		if err := json.Unmarshal(data, &permission); err != nil {
			return nil, fmt.Errorf(errMessageUnmarshalClass, err)
		}
		snzPermission, err := pm.sanitizeValidateOptimize(&permission, sanitize, validate, optimize)
		classInstance = snzPermission
		classType = aztypes.ClassTypeACPermission
		if err != nil {
			return nil, fmt.Errorf(errMessageUnmarshalClass, err)
		}
	default:
		return nil, fmt.Errorf("permcode: failed to unmarshal class - invalid type %s", class.Type)
	}
	strfy, err := aztext.Stringify(classInstance, nil)
	if err != nil {
		return nil, fmt.Errorf(errMessageUnmarshalClass, err)
	}
	return &aztypes.ClassInfo{
		SID:      azcrypto.ComputeStringSHA256(strfy),
		Type:     classType,
		Instance: classInstance,
	}, nil
}

// MarshalClass marshals a input class instance to a byte array.
func (pm *PermCodeManager) MarshalClass(instance any, sanitize bool, validate bool, optimize bool) ([]byte, error) {
	if instance == nil {
		return nil, errors.New("permcode: class cannot be marshaled from nil instance")
	}
	snzPolicy, err := pm.sanitizeValidateOptimize(instance, sanitize, validate, optimize)
	if err != nil {
		return nil, fmt.Errorf(errMessageMarshalClass, err)
	}
	data, err := json.Marshal(snzPolicy)
	if err != nil {
		return nil, fmt.Errorf("permcode: failed to marshal class - %w", err)
	}
	return data, nil
}
