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
	"strconv"
	"strings"
)

// SerializeTree serializes a tree object.
func (m *ObjectManager) SerializeTree(tree *Tree) ([]byte, error) {
	if tree == nil {
		return nil, fmt.Errorf("objects: tree is nil")
	}
	var sb strings.Builder
	for _, entry := range tree.entries {
		sb.WriteString(fmt.Sprintf("%06o %s %s %s\n", entry.mode, entry.otype, entry.oid, entry.name))
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
		parts := strings.SplitN(line, " ", 4)
		if len(parts) != 4 {
			return nil, fmt.Errorf("invalid entry format: %s", line)
		}
		mode, err := strconv.ParseUint(parts[0], 8, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid mode: %s", parts[0])
		}
		entry := TreeEntry{
			mode:  uint32(mode),
			otype: parts[1],
			oid:   parts[2],
			name:  parts[3],
		}
		tree.entries = append(tree.entries, entry)
	}
	return tree, nil
}
