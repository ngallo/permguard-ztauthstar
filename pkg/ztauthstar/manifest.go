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

package ztauthstar

type AutoGenerated struct {
	Metadata Metadata `json:"metadata"`
	Authz    Authz    `json:"authz"`
}
type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	License     string `json:"license"`
}
type Language struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
type Engine struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	Distribution string `json:"distribution"`
}
type Cedar00 struct {
	Language Language `json:"language"`
	Engine   Engine   `json:"engine"`
}
type Runtimes struct {
	Cedar00 Cedar00 `json:"cedar0.0+"`
}
type Location struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
}
type Root struct {
	Location Location `json:"location"`
	Runtime  string   `json:"runtime"`
	Schema   bool     `json:"schema"`
}
type Partitions struct {
	Root Root `json:"root"`
}
type Authz struct {
	Runtimes   Runtimes   `json:"runtimes"`
	Partitions Partitions `json:"partitions"`
}
