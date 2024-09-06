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
	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
)

const (
	// ACPolicyType is the AC policy type.
	ACPolicyType = "acpolicy"
)

// UURString is the UUR wildcard string.
type UURString aztext.WildcardString

// UUR is the Universally Unique Resource.
type UUR struct {
	account        aztext.WildcardString
	tenant         aztext.WildcardString
	domain         aztext.WildcardString
	resource       aztext.WildcardString
	resourceFilter aztext.WildcardString
}

// ARString is the AR wildcard string.
type ARString aztext.WildcardString

// AR is the Action Resource.
type AR struct {
	Resource aztext.WildcardString
	Action   aztext.WildcardString
}

// Policy is the policy.
type Policy struct {
	BaseType
	Name     string     `json:"name"`
	Actions  []ARString `json:"actions"`
	Resource UURString  `json:"resource"`
}
