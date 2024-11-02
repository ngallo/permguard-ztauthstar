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

package types

import (
	"fmt"
)

const (
	// PermCodeLanguage is the permcode language.
	PermCodeLanguage = "permcode"
	// PermCodeLanguageID is the permcode language ID.
	PermCodeLanguageID = uint32(1)

	// PermCodeSyntaxLatest is the latest permcode syntax.
	PermCodeSyntaxLatest = "permcode1"
	// PermCodeSyntaxIDLatest is the latest permcode syntax ID.
	PermCodeSyntaxIDLatest = uint32(1)

	// ClassTypeSchema is the object type for domain schemas.
	ClassTypeSchema = "schema"
	// ClassTypeObject is the object type for domain objects.
	ClassTypeIDSchema = uint32(1)

	// ClassTypeACPermission is the class type for an access control permission.
	ClassTypeACPermission = "acpermission"
	// ClassTypeIDACPermission is the object type for an access control permission.
	ClassTypeIDACPermission = uint32(2)

	// ClassTypeACPolicy is the object type for an access control policy.
	ClassTypeACPolicy = "acpolicy"
	// ClassTypeIDACPolicy is the object type for an access control policy.
	ClassTypeIDACPolicy = uint32(3)
)

// GetClassTypeID returns the language id, syntax id, and class type id.
func GetClassTypeID(classType string) (uint32, uint32,  uint32, error) {
	var classTypeID uint32
	switch classType {
	case ClassTypeSchema:
		classTypeID = ClassTypeIDSchema
	case ClassTypeACPermission:
		classTypeID = ClassTypeIDACPermission
	case ClassTypeACPolicy:
		classTypeID = ClassTypeIDACPolicy
	default:
		return 0, 0, 0, fmt.Errorf("permcode: invalid class type %s", classType)
	}
	return PermCodeLanguageID, PermCodeSyntaxIDLatest, classTypeID, nil
}

// GetClassType returns the language, syntax, and class type.
func GetClassType(langID, syntaxID, classTypeID uint32) (string, string,  string, error) {
	if langID != PermCodeLanguageID {
		return "", "", "", fmt.Errorf("permcode: invalid language ID %d", langID)
	}
	if syntaxID != PermCodeSyntaxIDLatest {
		return "", "", "", fmt.Errorf("permcode: invalid syntax ID %d", syntaxID)
	}
	var classType string
	switch classTypeID {
	case ClassTypeIDSchema:
		classType = ClassTypeSchema
	case ClassTypeIDACPermission:
		classType = ClassTypeACPermission
	case ClassTypeIDACPolicy:
		classType = ClassTypeACPolicy
	default:
		return "", "", "", fmt.Errorf("permcode: invalid class type %d", classTypeID)
	}
	return PermCodeLanguage, PermCodeSyntaxLatest, classType, nil
}

// Class is the base class.
type Class struct {
	SyntaxVersion string `json:"syntax"`
	Type          string `json:"type"`
}

// ClassInfo is the class info.
type ClassInfo struct {
	SID      string
	Type     string
	Instance any
}
