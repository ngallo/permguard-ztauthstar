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

package packets

const (
	PacketNullByte = 0x01
)

// Packet represents a packet.
type Packet struct {
	Data []byte
}

// Packetable represents a packet that can be serialized and deserialized.
type Packetable interface {
	GetType() int32
	GetData() []byte
	Serialize() error
	Deserialize() error
}

// DataPacket represents the packet data section.
type DataPacket struct {
	data []byte
}

// GetType returns the type of the packet.
func (p *DataPacket) GetType() int32 {
	return 1
}

// GetData returns the data of the packet.
func (p *DataPacket) GetData() []byte {
	return p.data
}
