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
	objInfo 	*ObjectInfo
	section 	int
	err 		error
}

// GetObjectInfo returns the object.
func (s *SectionObjectInfo) GetObjectInfo() *ObjectInfo {
	return s.objInfo
}

// GetSection returns the section.
func (s *SectionObjectInfo) GetSection() int {
	return s.section
}

// GetError returns the error.
func (s *SectionObjectInfo) GetError() error {
	return s.err
}

// NewSectionObjectInfo creates a new SectionObject.
func NewSectionObjectInfo(objInfo *ObjectInfo, section int, err error) (*SectionObjectInfo, error) {
	return &SectionObjectInfo{
		objInfo: objInfo,
		section: section,
		err: err,
	}, nil
}

// MultiSectionsObjectInfo represents an object with multiple sections.
type MultiSectionsObjectInfo struct {
	objInfos 	[]*SectionObjectInfo
	sections  	int
	err 		error
}

// NewMultiSectionsObjectInfo creates a new MultiSectionsObject.
func NewMultiSectionsObjectInfo(sections int, err error) (*MultiSectionsObjectInfo, error) {
	return &MultiSectionsObjectInfo{
		objInfos: make([]*SectionObjectInfo, 0),
		sections: sections,
		err: err,
	}, nil
}

// GetObjectInfos returns the objects.
func (m *MultiSectionsObjectInfo) GetObjectInfos() []*SectionObjectInfo {
	return azcopier.CopySlice(m.objInfos)
}

// GetSections returns the number of sections.
func (m *MultiSectionsObjectInfo) GetSections() int {
	return m.sections
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
	m.objInfos = append(m.objInfos, obj)
	return nil
}
