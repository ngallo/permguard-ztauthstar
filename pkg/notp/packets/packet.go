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

import (
	"bytes"
	"encoding/binary"
)

const (
	// PacketNullByte is the null byte used to separate data in the packet.
	PacketNullByte = 0x01
)

// Packet represents a packet.
type Packet struct {
	Data []byte
}

// Packetable represents a packet that can be serialized and deserialized.
type Packetable interface {
	GetType() int32
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}

// writeStreamDataPacket writes a stream data packet to the buffer.
func writeStreamDataPacket(data []byte, packetType *int32, packetStream *int32, payload []byte) ([]byte, error) {
	size := int32(len(payload))
	if packetType != nil {
		buf := new(bytes.Buffer)
		if err := binary.Write(buf, binary.LittleEndian, *packetType); err != nil {
			return nil, err
		}
		data = append(data, buf.Bytes()...)
	}
	if packetStream != nil {
		buf := new(bytes.Buffer)
		if err := binary.Write(buf, binary.LittleEndian, *packetStream); err != nil {
			return nil, err
		}
		data = append(data, buf.Bytes()...)
	}
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, size); err != nil {
		return nil, err
	}
	data = append(data, buf.Bytes()...)
	data = append(data, PacketNullByte)
	data = append(data, payload...)
	return data, nil
}

// writeDataPacket writes a data packet to the buffer.
func writeDataPacket(data []byte, payload []byte) ([]byte, error) {
	return writeStreamDataPacket(data, nil, nil, payload)
}
