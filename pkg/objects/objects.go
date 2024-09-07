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
	"time"
)

const (
	// ObjectTypeCommit is the object type for a commit.
	ObjectTypeCommit = "commit"
	// ObjectTypeTree is the object type for a tree.
	ObjectTypeTree   = "tree"
)

// Object represents the object.
type Object struct {
	OID   	string
	Content []byte
}

// ObjectInfo is the object info.
type ObjectInfo struct {
	OID			string
	Type		string
	Instance	any
}

// CommitInfo represents the author or committer of the commit.
type CommitInfo struct {
	Date  time.Time
}

// Commit represents a commit object.
type Commit struct {
	Tree      string
	Parents   []string
	Info CommitInfo
	Message   string
}

// TreeEntry represents a single entry in a tree object.
type TreeEntry struct {
	Mode uint32
	Type string
	OID string
	Name string
}

// Tree represents a tree object.
type Tree struct {
	Entries []TreeEntry
}
