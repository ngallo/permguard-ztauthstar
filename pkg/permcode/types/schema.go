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

package types

// DomainAction represents the domain action.
type DomainAction struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// DomainResource represents the domain resource.
type DomainResource struct {
	Name    string   		`json:"name"`
	Actions []DomainAction	`json:"actions"`
}

// Domain represents the domain.
type Domain struct {
	Name        string				`json:"name"`
	Description string     			`json:"description"`
	Resources   []DomainResource 	`json:"resources"`
}

// Schema represents the schema for the domains.
type Schema struct {
	Class
	Domains []Domain `json:"domains"`
}

// SchemaInfo is the schema info.
type SchemaInfo struct {
	ID		string
	Schema	*Schema
}
