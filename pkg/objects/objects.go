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
	"errors"
	"time"

	azcopier "github.com/permguard/permguard-core/pkg/extensions/copier"
	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
)

const (
	// ObjectTypeCommit is the object type for a commit.
	ObjectTypeCommit = "commit"
	// ObjectTypeTree is the object type for a tree.
	ObjectTypeTree = "tree"
	// ObjectTypeBlob is the object type for a blob.
	ObjectTypeBlob = "blob"
	// ZeroOID  represents the zero oid
	ZeroOID = "0000000000000000000000000000000000000000000000000000000000000000"
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

// NewObject creates a new object.
func NewObject(content []byte) *Object {
	return &Object{
		oid:     azcrypto.ComputeSHA256(content),
		content: content,
	}
}

// ObjectInfo is the object info.
type ObjectInfo struct {
	object   *Object
	otype    string
	instance any
}

// GetOID returns the OID of the object.
func (o *ObjectInfo) GetOID() string {
	if o.object == nil {
		return ""
	}
	return o.object.oid
}

// GetObject returns the object.
func (o *ObjectInfo) GetObject() *Object {
	return o.object
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
	parent  string
	info    CommitInfo
	message string
}

// GetTree returns the tree of the commit.
func (c *Commit) GetTree() string {
	return c.tree
}

// GetParent return the parent of the commit.
func (c *Commit) GetParent() string {
	return c.parent
}

// GetInfo returns the info of the commit.
func (c *Commit) GetInfo() CommitInfo {
	return c.info
}

// GetMessage returns the message of the commit.
func (c *Commit) GetMessage() string {
	return c.message
}

// NewCommit creates a new commit object.
func NewCommit(tree string, parents string, timestamp time.Time, message string) *Commit {
	if timestamp == (time.Time{}) {
		timestamp = time.Now()
	}
	return &Commit{
		tree:   tree,
		parent: parents,
		info: CommitInfo{
			date: timestamp,
		},
		message: message,
	}
}

// TreeEntry represents a single entry in a tree object.
type TreeEntry struct {
	otype string
	oid   string
	oname string
}

// NewTreeEntry creates a new tree entry.
func NewTreeEntry(otype, oid, oname string) *TreeEntry {
	return &TreeEntry{
		otype: otype,
		oid:   oid,
		oname: oname,
	}
}

// GetType returns the type of the tree entry.
func (t *TreeEntry) GetType() string {
	return t.otype
}

// GetOID returns the OID of the tree entry.
func (t *TreeEntry) GetOID() string {
	return t.oid
}

// GetOName returns the object name of the tree entry.
func (t *TreeEntry) GetOName() string {
	return t.oname
}

// Tree represents a tree object.
type Tree struct {
	entries []TreeEntry
}

// NewTree creates a new tree object.
func NewTree() *Tree {
	return &Tree{
		entries: make([]TreeEntry, 0),
	}
}

// GetEntries returns the entries of the tree.
func (t *Tree) GetEntries() []TreeEntry {
	return azcopier.CopySlice(t.entries)
}

// AddEntry adds an entry to the tree.
func (t *Tree) AddEntry(entry *TreeEntry) error {
	if entry == nil {
		return errors.New("objects: tree entry is nil")
	}
	for _, e := range t.entries {
		if e.GetOName() == entry.GetOName() {
			return errors.New("objects: tree entry already exists")
		}
	}
	t.entries = append(t.entries, *entry)
	return nil
}
