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
	"errors"
	"strconv"
	"strings"

	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
)

const (
	// ACPolicyType is the AC policy type.
	ACPolicyType = "acpolicy"
	// UUR format string: {account}:{tenant}:{domain}:{resource}:{resource-filter}.
	uurFormatString = "uur:%s:%s:%s:%s:%s"
	// AR format string: {resource}:{action}.
	arFormatString = "%s:%s"
)

// UURString is the UUR wildcard string.
type UURString aztext.WildcardString

// UUR is the Universally Unique Resource.
type UUR struct {
	account        int64
	tenant         aztext.WildcardString
	domain         aztext.WildcardString
	resource       aztext.WildcardString
	resourceFilter aztext.WildcardString
}

// Prase parses the UUR string.
func (s *UURString) Prase() (*UUR, error) {
    uurStr := string(*s)
    parts := strings.Split(uurStr, ":")
    if len(parts) != 6 || parts[0] != "uur" {
        return nil, errors.New("language: invalid uur string")
    }
    account, err := strconv.ParseInt(parts[1], 10, 64)
    if err != nil {
        return nil, errors.New("language: invalid account number, must be an integer")
    }
    tenant := parts[2]
    domain := parts[3]
    resource := parts[4]
    resourceFilter := parts[5]
	return &UUR{
		account:        account,
		tenant:         aztext.WildcardString(tenant),
		domain:         aztext.WildcardString(domain),
		resource:       aztext.WildcardString(resource),
		resourceFilter: aztext.WildcardString(resourceFilter),
	}, nil
}

// ARString is the AR wildcard string.
type ARString aztext.WildcardString

// Prase parses the UUR string.
func (s *ARString) Prase() (*AR, error) {
    uurStr := string(*s)
    parts := strings.Split(uurStr, ":")
    if len(parts) != 3 || parts[0] != "uur" {
        return nil, errors.New("language: invalid ar string")
    }
    resource := parts[1]
	action := parts[1]
	return &AR{
		Resource: aztext.WildcardString(resource),
		Action:   aztext.WildcardString(action),
	}, nil
}

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
