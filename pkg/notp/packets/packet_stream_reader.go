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
	"errors"
)

// PacketReader is a readr of packets from the NOTP protocol.
type PacketReader struct {
	packet *Packet
}

// NewPacketReader creates a new packet readr.
func NewPacketReader(packet *Packet) (*PacketReader, error) {
	if packet == nil {
		return nil, errors.New("notp: nil packet")
	}
	if packet.Data == nil {
		packet.Data = []byte{}
	}
	return &PacketReader{
		packet: packet,
	}, nil
}

// ReadProtocol read a protocol packet.
func (w *PacketReader) ReadProtocol() (*ProtocolPacket, error) {
	if len(w.packet.Data) == 0 {
		return nil, errors.New("notp: missing protocol packet")
	}
	payload, _, _, err := readDataPacket(w.packet.Data)
	if err != nil {
		return nil, err
	}
	protocol := &ProtocolPacket{}
	err = protocol.Deserialize(payload)
	if err != nil {
		return nil, err
	}
	return protocol, nil
}

// DataPacketState is the state of a data packet.
type DataPacketState struct {
	CurrentOffset int
	IsComplete    bool
}

// ReadNextDataPacket read next data packet.
func (w *PacketReader) ReadNextDataPacket(state *DataPacketState) ([]byte, *DataPacketState, error) {
	return nil, nil, nil
}
