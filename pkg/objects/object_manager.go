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
	"strconv"
	"strings"
)

const (
	// PacketNullByte is the null byte used to separate data in the packet.
	PacketNullByte = 0xFF
)

// ObjectManager is the manager for policies.
type ObjectManager struct {
}

// NewObjectManager creates a new ObjectManager.
func NewObjectManager() (*ObjectManager, error) {
	return &ObjectManager{}, nil
}

// CreateObject creates an object.
func (m *ObjectManager) createOject(objectType string, content []byte) (*Object, error) {
	length := len(content)
	var buffer bytes.Buffer
	buffer.WriteString(objectType)
	buffer.WriteString(" ")
	buffer.WriteString(fmt.Sprintf("%d", length))
	buffer.WriteByte(PacketNullByte)
	buffer.Write(content)
	objContent := buffer.Bytes()
	return NewObject(objContent), nil
}

// CreateCommitObject creates a commit object.
func (m *ObjectManager) CreateCommitObject(commit *Commit) (*Object, error) {
	commitBytes, err := m.SerializeCommit(commit)
	if err != nil {
		return nil, err
	}
	return m.createOject(ObjectTypeCommit, commitBytes)
}

// CreateTreeObject creates a tree object.
func (m *ObjectManager) CreateTreeObject(tree *Tree) (*Object, error) {
	treeBytes, err := m.SerializeTree(tree)
	if err != nil {
		return nil, err
	}
	if len(treeBytes) == 0 {
		return nil, errors.New("objects: data is empty")
	}
	return m.createOject(ObjectTypeTree, treeBytes)
}

// CreateBlobObject creates a blob object.
func (m *ObjectManager) CreateBlobObject(data []byte) (*Object, error) {
	if len(data) == 0 {
		return nil, errors.New("objects: data is empty")
	}
	return m.createOject(ObjectTypeBlob, data)
}

// CreateObjectFormData create the object form data.
func (m *ObjectManager) CreateObjectFormData(binaryData []byte) (*Object, error) {
	return NewObject(binaryData), nil
}

// GetObjectInfo gets the object info.
func (m *ObjectManager) GetObjectInfo(object *Object) (*ObjectInfo, error) {
	if object == nil {
		return nil, errors.New("objects: object is nil")
	}
	objContent := object.content
	nulIndex := bytes.IndexByte(objContent, PacketNullByte)
	if nulIndex == -1 {
		return nil, fmt.Errorf("objects: invalid object format: no NUL separator found")
	}
	header := string(objContent[:nulIndex])
	headerParts := strings.SplitN(header, " ", 2)
	if len(headerParts) != 2 {
		return nil, fmt.Errorf("objects: invalid object header format")
	}
	objectType := headerParts[0]
	length, err := strconv.Atoi(headerParts[1])
	if err != nil {
		return nil, fmt.Errorf("objects: invalid length: %v", err)
	}
	start := nulIndex + 1
	end := start + length
	content := objContent[start:end]
	if len(content) != length {
		return nil, fmt.Errorf("objects: content length mismatch: expected %d, got %d", length, len(content))
	}
	var instance any
	switch objectType {
	case ObjectTypeCommit:
		commit, err := m.DeserializeCommit(content)
		if err != nil {
			return nil, err
		}
		instance = commit
	case ObjectTypeTree:
		tree, err := m.DeserializeTree(content)
		if err != nil {
			return nil, err
		}
		instance = tree
	case ObjectTypeBlob:
		instance = content
	default:
		return nil, fmt.Errorf("objects: unsupported object type %s", objectType)
	}
	return &ObjectInfo{
		object:   object,
		otype:    objectType,
		instance: instance,
	}, nil
}

// buildCommitHistory builds the commit history.
func (m *ObjectManager) buildCommitHistory(fromCommitID string, toCommitID string, match bool, history []Commit, objFunc func(string) (*Object, error)) (bool, []Commit, error) {
	if fromCommitID == ZeroOID && toCommitID == ZeroOID {
		match = true
		return match, history, nil
	}
	var commitObj *Object
	var err error
	if fromCommitID != ZeroOID {
		commitObj, err = objFunc(fromCommitID)
		if err != nil {
			return false, nil, err
		}
	}
	var commit *Commit
	if commitObj != nil {
		commitObjInfo, err := m.GetObjectInfo(commitObj)
		if err != nil {
			return false, nil, err
		}
		var ok bool
		commit, ok = commitObjInfo.GetInstance().(*Commit)
		if !ok {
			return false, nil, fmt.Errorf("objects: invalid object type")
		}
		if commit != nil {
			history = append(history, *commit)
		}
	}
	if commitObj == nil || commit == nil {
		return match, history, nil
	}
	if commitObj.GetOID() == toCommitID {
		match = true
		return match, history, nil
	}
	return m.buildCommitHistory(commit.GetParent(), toCommitID, match, history, objFunc)
}

// BuildCommitHistory builds the commit history.
func (m *ObjectManager) BuildCommitHistory(fromCommitID string, toCommitID string, reverse bool, objFunc func(string) (*Object, error)) (bool, []Commit, error) {
	if fromCommitID == ZeroOID && toCommitID != ZeroOID {
		return false, nil, fmt.Errorf("objects: invalid from commit ID")
	}
	match, history, err := m.buildCommitHistory(fromCommitID, toCommitID, false, []Commit{}, objFunc)
	if err != nil && reverse {
		for i, j := 0, len(history)-1; i < j; i, j = i+1, j-1 {
			history[i], history[j] = history[j], history[i]
		}
	}
	return match, history, err
}
