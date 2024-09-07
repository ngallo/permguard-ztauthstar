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
	"bytes"
	"errors"
	"fmt"
	"strings"

	azcrypto "github.com/permguard/permguard-core/pkg/extensions/crypto"
)

// ObjectManager is the manager for policies.
type ObjectManager struct {
}

// NewObjectManager creates a new ObjectManager.
func NewObjectManager() *ObjectManager {
	return &ObjectManager{}
}

// CreateObject creates an object.
func (m *ObjectManager) createOject(objectType string, content []byte) (Object, error) {
	length := len(content)
	var buffer bytes.Buffer
	buffer.WriteString(objectType)
	buffer.WriteString(" ")
	buffer.WriteString(fmt.Sprintf("%d", length))
	buffer.WriteByte(0)
	buffer.Write(content)
	objContent :=  buffer.Bytes()
	return Object{
		OID:     azcrypto.ComputeSHA256(content),
		Content: objContent,
	}, nil
}

// CreateCommitObject creates a commit object.
func (m *ObjectManager) CreateCommitObject(commit *Commit) (Object, error) {
	commitBytes, err := m.SerializeCommit(commit)
	if err != nil {
		return Object{}, err
	}
	return m.createOject(ObjectTypeCommit, commitBytes)
}

// CreateTreeObject creates a tree object.
func (m *ObjectManager) CreateTreeObject(tree *Tree) (Object, error) {
	treeBytes, err := m.SerializeTree(tree)
	if err != nil {
		return Object{}, err
	}
	return m.createOject(ObjectTypeTree, treeBytes)
}

// GetObjectInfo gets the object info.
func (m *ObjectManager) GetObjectInfo(object Object) (*ObjectInfo, error) {
	if len(object.Content) == 0 {
		return nil, errors.New("objects: object content is empty")
	}
	headerParts := strings.SplitN(string(object.Content), " ", 2)
	if len(headerParts) < 2 {
		return nil, errors.New("objects: object is not a commit")
	}
	objectType := headerParts[0]
	contentStartIndex := bytes.Index(object.Content, []byte{0}) + 1
	if contentStartIndex <= 0 || contentStartIndex >= len(object.Content) {
		return nil, errors.New("objects: invalid object content")
	}
	objContent := object.Content[contentStartIndex:]
	var instance any
	switch objectType {
	case ObjectTypeCommit:
		commit, err := m.DeserializeCommit(objContent)
		if err != nil {
			return nil, err
		}
		instance = commit
	case ObjectTypeTree:
		tree, err := m.DeserializeTree(objContent)
		if err != nil {
			return nil, err
		}
		instance = tree
	default:
		return nil, fmt.Errorf("objects: unsupported object type %s", objectType)
	}
	return &ObjectInfo{
		OID:      object.OID,
		Type:     objectType,
		Instance: instance,
	}, nil
}
