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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestObjectManager tests the functions of ObjectManager.
func TestObjectManager(t *testing.T) {
	objectManager := NewObjectManager()

	// Test for CreateCommitObject and GetObjectInfo
	t.Run("Test CreateCommitObject and GetObjectInfo", func(t *testing.T) {
		assert := assert.New(t)
		commit := &Commit{
			Tree:    "3b18e17a0e8664d3dffab99ebf6d730ddc6e8649",
			Parents: []string{"a1b2c3d4e5f678901234567890abcdef12345678"},
			Info: CommitInfo{
				Date: time.Unix(1628704800, 0), // Example Unix timestamp
			},
			Message: "Initial commit",
		}

		// Create commit object
		commitObj, err := objectManager.CreateCommitObject(commit)
		assert.Nil(err)
		assert.NotEmpty(commitObj.OID, "OID should not be empty")
		assert.NotEmpty(commitObj.Content, "Commit content should not be empty")

		// Get object info
		objectInfo, err := objectManager.GetObjectInfo(commitObj)
		assert.Nil(err)
		assert.Equal(ObjectTypeCommit, objectInfo.Type, "Expected commit type")
		assert.NotNil(objectInfo.Instance, "Commit instance should not be nil")

		// Cast to commit and validate fields
		retrievedCommit := objectInfo.Instance.(*Commit)
		assert.Equal(commit.Tree, retrievedCommit.Tree, "Tree mismatch")
		assert.Equal(commit.Parents, retrievedCommit.Parents, "Parents mismatch")
		assert.Equal(commit.Info.Date.Unix(), retrievedCommit.Info.Date.Unix(), "Commit date mismatch")
		assert.Equal(commit.Message, retrievedCommit.Message, "Message mismatch")
	})

	// Test for CreateTreeObject and GetObjectInfo
	t.Run("Test CreateTreeObject and GetObjectInfo", func(t *testing.T) {
		assert := assert.New(t)
		tree := &Tree{
			Entries: []TreeEntry{
				{Mode: 0100644, Type: "blob", OID: "6eb715b073c6b28e03715129e03a0d52c8e21b73", Name: "README.md"},
				{Mode: 0100755, Type: "blob", OID: "a7fdb22705a5e6145b6a8b1fa947825c5e97a51c", Name: "script.sh"},
				{Mode: 040000, Type: "tree", OID: "a7fdb33705a5e6145b6a8b1fa947825c5e97a51c", Name: "src"},
			},
		}

		// Create tree object
		treeObj, err := objectManager.CreateTreeObject(tree)
		assert.Nil(err)
		assert.NotEmpty(treeObj.OID, "OID should not be empty")
		assert.NotEmpty(treeObj.Content, "Tree content should not be empty")

		// Get object info
		objectInfo, err := objectManager.GetObjectInfo(treeObj)
		assert.Nil(err)
		assert.Equal(ObjectTypeTree, objectInfo.Type, "Expected tree type")
		assert.NotNil(objectInfo.Instance, "Tree instance should not be nil")

		// Cast to tree and validate fields
		retrievedTree := objectInfo.Instance.(*Tree)
		assert.Equal(len(tree.Entries), len(retrievedTree.Entries), "Entries length mismatch")
		for i, entry := range tree.Entries {
			assert.Equal(entry.Mode, retrievedTree.Entries[i].Mode, "Mode mismatch for entry %d", i)
			assert.Equal(entry.Type, retrievedTree.Entries[i].Type, "Type mismatch for entry %d", i)
			assert.Equal(entry.OID, retrievedTree.Entries[i].OID, "OID mismatch for entry %d", i)
			assert.Equal(entry.Name, retrievedTree.Entries[i].Name, "Name mismatch for entry %d", i)
		}
	})

	// Test for CreateBlobObject and GetObjectInfo (new test for blob type)
	t.Run("Test CreateBlobObject and GetObjectInfo", func(t *testing.T) {
		assert := assert.New(t)
		blobData := []byte("This is the content of the blob object")

		// Create blob object
		blobObj, err := objectManager.CreateBlobObject(blobData)
		assert.Nil(err)
		assert.NotEmpty(blobObj.OID, "OID should not be empty")
		assert.NotEmpty(blobObj.Content, "Blob content should not be empty")

		// Get object info
		objectInfo, err := objectManager.GetObjectInfo(blobObj)
		assert.Nil(err)
		assert.Equal(ObjectTypeBlob, objectInfo.Type, "Expected blob type")
		assert.NotNil(objectInfo.Instance, "Blob instance should not be nil")

		// Validate the content of the blob
		retrievedBlob := objectInfo.Instance.([]byte)
		assert.Equal(blobData, retrievedBlob, "Blob content mismatch")
	})

	// Test for invalid data
	t.Run("Test invalid object", func(t *testing.T) {
		assert := assert.New(t)
		invalidObj := &Object{Content: []byte{}}
		_, err := objectManager.GetObjectInfo(invalidObj)
		assert.NotNil(err, "Expected error for empty object content")

		// Test for incorrect object type
		invalidObj.Content = []byte("xx 12\000some content")
		_, err = objectManager.GetObjectInfo(invalidObj)
		assert.NotNil(err, "Expected error for wrong object type")
		assert.Contains(err.Error(), "objects: unsupported object type ", "Expected objects: unsupported object type ")
	})
}
