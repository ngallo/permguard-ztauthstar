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

package refsobjects

import (
	"errors"
)

// RefsObjPacketWriter is a writer of refs objects packet from the NOTP protocol.
type RefsObjPacketWriter struct {
	packet *RefsObjPacket
}

// NewRefsObjPacketWriter creates a new refs objects packet writer.
func NewRefsObjPacketWriter(packet *RefsObjPacket) (*RefsObjPacketWriter, error) {
	if packet == nil {
		return nil, errors.New("notp: nil packet")
	}
	return &RefsObjPacketWriter{
		packet: packet,
	}, nil
}
