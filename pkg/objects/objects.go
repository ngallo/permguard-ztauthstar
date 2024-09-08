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
	ObjectTypeTree = "tree"
	// ObjectTypeBlob is the object type for a blob.
	ObjectTypeBlob = "blob"
)

// Object represents the object.
type Object struct {
	oid     string
	content []byte
}

// GetOID returns the OID of the object.
func (o *Object) GetOID() string {
	return o.oid
}

// GetContent returns the content of the object.
func (o *Object) GetContent() []byte {
	return o.content
}

// ObjectInfo is the object info.
type ObjectInfo struct {
	oid      string
	otype    string
	instance any
}

// GetOID returns the OID of the object.
func (o *ObjectInfo) GetOID() string {
	return o.oid
}

// GetType returns the type of the object.
func (o *ObjectInfo) GetType() string {
	return o.otype
}

// GetInstance returns the instance of the object.
func (o *ObjectInfo) GetInstance() any {
	return o.instance
}

// CommitInfo represents the author or committer of the commit.
type CommitInfo struct {
	date time.Time
}

// GetDate returns the date of the commit info.
func (c *CommitInfo) GetDate() time.Time {
	return c.date
}

// Commit represents a commit object.
type Commit struct {
	tree    string
	parents []string
	info    CommitInfo
	message string
}

// GetTree returns the tree of the commit.
func (c *Commit) GetTree() string {
	return c.tree
}

// GetParents returns the parents of the commit.
func (c *Commit) GetParents() []string {
	return c.parents
}

// GetInfo returns the info of the commit.
func (c *Commit) GetInfo() CommitInfo {
	return c.info
}

// GetMessage returns the message of the commit.
func (c *Commit) GetMessage() string {
	return c.message
}

// TreeEntry represents a single entry in a tree object.
type TreeEntry struct {
	mode  uint32
	otype string
	oid   string
	name  string
}

// GetMode returns the mode of the tree entry.
func (t *TreeEntry) GetMode() uint32 {
	return t.mode
}

// GetType returns the type of the tree entry.
func (t *TreeEntry) GetType() string {
	return t.otype
}

// GetOID returns the OID of the tree entry.
func (t *TreeEntry) GetOID() string {
	return t.oid
}

// GetName returns the name of the tree entry.
func (t *TreeEntry) GetName() string {
	return t.name
}

// Tree represents a tree object.
type Tree struct {
	entries []TreeEntry
}

// GetEntries returns the entries of the tree.
func (t *Tree) GetEntries() []TreeEntry {
	return t.entries
}
