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
	"fmt"
	"strconv"
	"strings"
	"time"
)

// SerializeCommit serializes a commit object.
func (m *ObjectManager) SerializeCommit(commit *Commit) ([]byte, error) {
	if commit == nil {
		return nil, errors.New("objects: commit is nil")
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("tree %s\n", commit.tree))
	sb.WriteString(fmt.Sprintf("parent %s\n", commit.parent))
	sb.WriteString(fmt.Sprintf("author %d %s %s\n", commit.info.authorTimestamp.Unix(), commit.info.authorTimestamp.Format("-0700"), commit.info.author))
	sb.WriteString(fmt.Sprintf("committer %d %s %s\n", commit.info.committerTimestamp.Unix(), commit.info.committerTimestamp.Format("-0700"), commit.info.committer))
	sb.WriteString(commit.message)
	return []byte(sb.String()), nil
}

// parseIdentity parses the identity line.
func (m *ObjectManager) parseIdentity(line string) (string, time.Time) {
	parts := strings.Split(line, " ")
	var name string
	var date time.Time

	if len(parts) >= 2 {
		unixTime, _ := strconv.ParseInt(parts[0], 10, 64)
		loc, _ := time.Parse(parts[1], parts[0])
		date = time.Unix(unixTime, 0).In(loc.Location())
	}
	if len(parts) >= 3 {
		name = strings.Join(parts[2:], " ")
	}
	return name, date
}

// DeserializeCommit deserializes a commit object.
func (m *ObjectManager) DeserializeCommit(data []byte) (*Commit, error) {
	if data == nil {
		return nil, errors.New("objects: data is nil")
	}
	inputStr := string(data)
	lines := strings.Split(inputStr, "\n")
	commit := &Commit{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, "tree ") {
			commit.tree = strings.TrimPrefix(line, "tree ")
		} else if strings.HasPrefix(line, "parent ") {
			commit.parent = strings.TrimPrefix(line, "parent ")
		} else if strings.HasPrefix(line, "author ") {
			author, date := m.parseIdentity(strings.TrimPrefix(line, "author "))
			commit.info.author = author
			commit.info.authorTimestamp = date
		} else if strings.HasPrefix(line, "committer ") {
			committer, date := m.parseIdentity(strings.TrimPrefix(line, "committer "))
			commit.info.committer = committer
			commit.info.committerTimestamp = date
		} else if i == len(lines)-1 {
			commit.message = line
		}
	}
	return commit, nil
}
