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
// See the License for the specific language governing schemas and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package permcode

import (
	"errors"
	"fmt"
	"sort"

	azsanitizers "github.com/permguard/permguard-abs-language/pkg/permcode/sanitizers"
	aztypes "github.com/permguard/permguard-abs-language/pkg/permcode/types"
	azvalidators "github.com/permguard/permguard-core/pkg/extensions/validators"
)

// UnmarshalSchema unmarshals a input byte array to the schema instance.
func (pm *PermCodeManager) UnmarshalSchema(data []byte, sanitize bool, validate bool, optimize bool) (*aztypes.SchemaInfo, error) {
	clasInfo, err := pm.UnmarshalClass(data, aztypes.ClassTypeSchema, sanitize, validate, optimize)
	if err != nil {
		return nil, err
	}
	return &aztypes.SchemaInfo{
		SID:    clasInfo.SID,
		Schema: clasInfo.Instance.(*aztypes.Schema),
	}, nil
}

// sanitizeValidateOptimize sanitizes, validates and optimize the input schema.
func (pm *PermCodeManager) sanitizeValidateOptimizeSchema(schema *aztypes.Schema, sanitize bool, validate bool, optimize bool) (*aztypes.Schema, error) {
	var err error
	targetSchema := schema
	if sanitize {
		targetSchema, err = pm.sanitizeSchema(schema)
		if err != nil {
			return nil, err
		}
	}
	if validate {
		valid, err := pm.validateSchema(targetSchema)
		if err != nil {
			return nil, err
		}
		if !valid {
			return nil, errors.New("permcode: schema is invalid")
		}
	}
	if optimize {
		targetSchema, err = pm.optimizeSchema(targetSchema)
		if err != nil {
			return nil, err
		}
	}
	return targetSchema, nil
}

// sanitizeSchema sanitizes the schema.
func (pm *PermCodeManager) sanitizeSchema(schema *aztypes.Schema) (*aztypes.Schema, error) {
	schema.SyntaxVersion = azsanitizers.SanitizeString(schema.SyntaxVersion)
	schema.Type = azsanitizers.SanitizeString(schema.Type)
	for _, domain := range schema.Domains {
		domain.Name = azsanitizers.SanitizeString(domain.Name)
		if domain.Resources == nil {
			domain.Resources = []aztypes.DomainResource{}
		}
		for _, resource := range domain.Resources {
			resource.Name = azsanitizers.SanitizeString(resource.Name)
			if resource.Actions == nil {
				resource.Actions = []aztypes.DomainAction{}
			}
			for _, action := range resource.Actions {
				action.Name = azsanitizers.SanitizeString(action.Name)
			}
		}
	}
	return schema, nil
}

// validateSchema validates the schema.
func (pm *PermCodeManager) validateSchema(schema *aztypes.Schema) (bool, error) {
	if schema.SyntaxVersion != aztypes.PermCodeSyntaxLatest {
		return false, fmt.Errorf(`permcode: unsupported policy syntax version '%s'`, schema.SyntaxVersion)
	}
	if schema.Type != aztypes.ClassTypeSchema {
		return false, fmt.Errorf(`permcode: invalid schema type '%s'`, schema.Type)
	}
	domainNames := make(map[string]bool)
	for _, domain := range schema.Domains {
		if !azvalidators.ValidateName(domain.Name) {
			return false, fmt.Errorf(`permcode: invalid domain name '%s'`, domain.Name)
		}
		if _, exists := domainNames[domain.Name]; exists {
			return false, fmt.Errorf(`permcode: duplicate domain name '%s'`, domain.Name)
		}
		domainNames[domain.Name] = true
		resourceNames := make(map[string]bool)
		for _, resource := range domain.Resources {
			if !azvalidators.ValidateName(resource.Name) {
				return false, fmt.Errorf(`permcode: invalid resource name '%s' in domain '%s'`, resource.Name, domain.Name)
			}
			if _, exists := resourceNames[resource.Name]; exists {
				return false, fmt.Errorf(`permcode: duplicate resource name '%s' in domain '%s'`, resource.Name, domain.Name)
			}
			resourceNames[resource.Name] = true
			actionNames := make(map[string]bool)
			for _, action := range resource.Actions {
				if !azvalidators.ValidateName(action.Name) {
					return false, fmt.Errorf(`permcode: invalid action name '%s' in resource '%s' of domain '%s'`, action.Name, resource.Name, domain.Name)
				}
				if _, exists := actionNames[action.Name]; exists {
					return false, fmt.Errorf(`permcode: duplicate action name '%s' in resource '%s' of domain '%s'`, action.Name, resource.Name, domain.Name)
				}
				actionNames[action.Name] = true
			}
		}
	}
	return true, nil
}

// optimizeSchema optimizes the schema.
func (pm *PermCodeManager) optimizeSchema(schema *aztypes.Schema) (*aztypes.Schema, error) {
	uniqueDomains := make(map[string]aztypes.Domain)
	for _, domain := range schema.Domains {
		if _, exists := uniqueDomains[domain.Name]; !exists {
			uniqueResources := make(map[string]aztypes.DomainResource)
			for _, resource := range domain.Resources {
				if _, exists := uniqueResources[resource.Name]; !exists {
					uniqueActions := make(map[string]aztypes.DomainAction)
					for _, action := range resource.Actions {
						if _, exists := uniqueActions[action.Name]; !exists {
							uniqueActions[action.Name] = action
						}
					}
					resource.Actions = make([]aztypes.DomainAction, 0, len(uniqueActions))
					for _, action := range uniqueActions {
						resource.Actions = append(resource.Actions, action)
					}
					sort.Slice(resource.Actions, func(i, j int) bool {
						return resource.Actions[i].Name < resource.Actions[j].Name
					})
					uniqueResources[resource.Name] = resource
				}
			}
			domain.Resources = make([]aztypes.DomainResource, 0, len(uniqueResources))
			for _, resource := range uniqueResources {
				domain.Resources = append(domain.Resources, resource)
			}
			sort.Slice(domain.Resources, func(i, j int) bool {
				return domain.Resources[i].Name < domain.Resources[j].Name
			})
			uniqueDomains[domain.Name] = domain
		}
	}
	schema.Domains = make([]aztypes.Domain, 0, len(uniqueDomains))
	for _, domain := range uniqueDomains {
		schema.Domains = append(schema.Domains, domain)
	}
	sort.Slice(schema.Domains, func(i, j int) bool {
		return schema.Domains[i].Name < schema.Domains[j].Name
	})
	return schema, nil
}
