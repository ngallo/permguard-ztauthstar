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

package objects

import (
	"fmt"
	"strings"
)

// SerializeTree serializes a tree object.
func (m *ObjectManager) SerializeTree(tree *Tree) ([]byte, error) {
	if tree == nil {
		return nil, fmt.Errorf("objects: tree is nil")
	}
	var sb strings.Builder
	for _, entry := range tree.entries {
		sb.WriteString(fmt.Sprintf("%s %s %s\n", entry.otype, entry.oid, entry.oname))
	}
	return []byte(sb.String()), nil
}

// DeserializeTree deserializes a tree object.
func (m *ObjectManager) DeserializeTree(data []byte) (*Tree, error) {
	if data == nil {
		return nil, fmt.Errorf("objects: data is nil")
	}
	inputStr := string(data)
	lines := strings.Split(strings.TrimSpace(inputStr), "\n")
	tree := &Tree{}
	for _, line := range lines {
		parts := strings.SplitN(line, " ", 5)
		if len(parts) != 5 {
			return nil, fmt.Errorf("objects: invalid entry format: %s", line)
		}
		entry := TreeEntry{
			otype: parts[1],
			oid:   parts[2],
			oname: parts[3],
		}
		tree.entries = append(tree.entries, entry)
	}
	return tree, nil
}
