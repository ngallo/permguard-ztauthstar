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

import (
	"encoding/json"
	"fmt"
)

// NewManifest creates a new manifest.
func NewManifest(name string) (*Manifest, error) {
	manifest := &Manifest{
		Metadata: Metadata{
			Name:        name,
			Description: "This is a sample manifest",
		},
		Authz: Authz{
			Runtimes:   make(map[string]Runtime),
			Partitions: make(map[string]Partition),
		},
	}
	return manifest, nil
}

// ValidateManifest validates the manifest.
func ValidateManifest(manifest *Manifest) (bool, *Manifest, error) {
	data, err := json.Marshal(manifest)
    if err != nil {
        return false, nil, fmt.Errorf("[ztas] failed to serialize the manifest: %w", err)
    }
    return ValidateManifestData(data)
}

// ValidateManifestData validates the manifest data.
func ValidateManifestData(data []byte) (bool, *Manifest, error) {
    var manifest Manifest
    if err := json.Unmarshal([]byte(data), &manifest); err != nil {
        return false, nil, err
    }
	return true, &manifest, nil
}
