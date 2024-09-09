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

	azcopier "github.com/permguard/permguard-core/pkg/extensions/copier"
)

// SectionObjectInfo represents a child object info.
type SectionObjectInfo struct {
	obj			*Object
	numOfSects 	int
	err 		error
}

// GetObject returns the object.
func (s *SectionObjectInfo) GetObject() *Object {
	return s.obj
}

// GetNumberOfSections returns the number sections.
func (s *SectionObjectInfo) GetNumberOfSections() int {
	return s.numOfSects
}

// GetError returns the error.
func (s *SectionObjectInfo) GetError() error {
	return s.err
}

// NewSectionObjectInfo creates a new SectionObject.
func NewSectionObjectInfo(obj *Object, section int, err error) (*SectionObjectInfo, error) {
	return &SectionObjectInfo{
		obj: obj,
		numOfSects: section,
		err: err,
	}, nil
}

// MultiSectionsObjectInfo represents an object with multiple sections.
type MultiSectionsObjectInfo struct {
	path 			string
	objSecInfos 	[]*SectionObjectInfo
	numOfSects  	int
	err 			error
}

// NewMultiSectionsObjectInfo creates a new MultiSectionsObject.
func NewMultiSectionsObjectInfo(path string, numOfSections int, err error) (*MultiSectionsObjectInfo, error) {
	return &MultiSectionsObjectInfo{
		objSecInfos: make([]*SectionObjectInfo, 0),
		numOfSects: numOfSections,
		err: err,
	}, nil
}

// GetSectionObjectInfos returns the section object infos.
func (m *MultiSectionsObjectInfo) GetSectionObjectInfos() []*SectionObjectInfo {
	return azcopier.CopySlice(m.objSecInfos)
}

// GetSections returns the number of sections.
func (m *MultiSectionsObjectInfo) GetSections() int {
	return m.numOfSects
}

// GetError returns the error.
func (m *MultiSectionsObjectInfo) GetError() error {
	return m.err
}

// AddSectionObjectInfo adds a section object info.
func (m *MultiSectionsObjectInfo) AddSectionObjectInfo(obj *SectionObjectInfo) error {
	if obj == nil {
		return errors.New("object is nil")
	}
	m.objSecInfos = append(m.objSecInfos, obj)
	return nil
}

// AddSectionObjectInfoWithParams adds a section object info with parameters.
func (m *MultiSectionsObjectInfo) AddSectionObjectInfoWithParams(obj *Object, section int, err error) error {
	objSecInfo, err := NewSectionObjectInfo(obj, section, err)
	if err != nil {
		return err
	}
	return m.AddSectionObjectInfo(objSecInfo)
}
