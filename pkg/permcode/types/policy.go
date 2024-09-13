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
	"strings"

	aztext "github.com/permguard/permguard-core/pkg/extensions/text"
)

const (
	// UUR format string: {parition}:{account}:{tenant}:{domain}:{resource}:{resource-filter}.
	uurFormatString = "uur:%s:%s:%s:%s:%s%s"
	// AR format string: {resource}:{action}.
	arFormatString = "ar:%s:%s"
	// resourceFilterSeparator is the separator for the resource filter.
	resourceFilterSeparator = "/"
	// KeywordPartition is the account keyword.
	KeywordPartition = "$partition"
	// KeywordAccount is the account keyword.
	KeywordAccount = "$account"
	// KeywordTenant is the tenant keyword.
	KeywordTenant = "$tenant"
)

// UURString is the UUR wildcard string.
type UURString aztext.WildcardString

// FormatUURString formats the UUR string.
func FormatUURString(partition, account string, tenant, domain, resource aztext.WildcardString, resourceFileter []aztext.WildcardString) UURString {
	resFilter := ""
	for _, f := range resourceFileter {
		resFilter = fmt.Sprintf("%s%s%s", resFilter, resourceFilterSeparator, f)
	}
	return UURString(fmt.Sprintf(uurFormatString, partition, account, tenant, domain, resource, resFilter))
}

// UUR is the Universally Unique Resource.
type UUR struct {
	Partition      string
	Account        string
	Tenant         aztext.WildcardString
	Domain         aztext.WildcardString
	Resource       aztext.WildcardString
	ResourceFilter []aztext.WildcardString
}

// Prase parses the UUR string.
func (s *UURString) Prase() (*UUR, error) {
	uurStr := string(*s)
	parts := strings.Split(uurStr, ":")
	if len(parts) != 6 || parts[0] != "uur" {
		return nil, fmt.Errorf("permcode: invalid uur string (%s)", uurStr)
	}
	partition := parts[1]
	account := parts[2]
	tenant := parts[3]
	domain := parts[4]
	resParts := strings.Split(parts[5], resourceFilterSeparator)
	resource := resParts[0]
	resourceFilter := []aztext.WildcardString{}
	if len(resParts) > 1 {
		for _, filter := range resParts[1:] {
			resourceFilter = append(resourceFilter, aztext.WildcardString(filter))
		}
	}
	return &UUR{
		Partition:      partition,
		Account:        account,
		Tenant:         aztext.WildcardString(tenant),
		Domain:         aztext.WildcardString(domain),
		Resource:       aztext.WildcardString(resource),
		ResourceFilter: resourceFilter,
	}, nil
}

// ARString is the AR wildcard string.
type ARString aztext.WildcardString

// FormatARString formats the AR string.
func FormatARString(resource, action aztext.WildcardString) ARString {
	return ARString(fmt.Sprintf(arFormatString, resource, action))
}

// Prase parses the UUR string.
func (s *ARString) Prase() (*AR, error) {
	arStr := string(*s)
	parts := strings.Split(arStr, ":")
	if len(parts) != 3 || parts[0] != "ar" {
		return nil, fmt.Errorf(`permcode: invalid ar string "(%s)"`, arStr)
	}
	resource := parts[1]
	action := parts[2]
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
	Class
	Name     string     `json:"name"`
	Actions  []ARString `json:"actions"`
	Resource UURString  `json:"resource"`
}

// PolicyInfo is the policy info.
type PolicyInfo struct {
	SID    string
	Policy *Policy
}
