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
	"errors"
	"unsafe"
)

// PacketWriter is a writer of packets from the NOTP protocol.
type PacketWriter struct {
	packet *Packet
}

// NewPacketWriter creates a new packet writer.
func NewPacketWriter(packet *Packet) (*PacketWriter, error) {
	if packet == nil {
		return nil, errors.New("notp: nil packet")
	}
	return &PacketWriter{
		packet: packet,
	}, nil
}

// WriteProtocolSection writes the protocol section.
func (w *PacketWriter) WriteProtocolSection(protocol *Protocol) error {
	blockNumber := 0
	block, exist := seekBlock(w.packet.Data, blockNumber)
	if !exist {
		block = []byte{0}
	}
	buffer := bytes.NewBuffer(block)
	numbers := []int16{protocol.Version, protocol.Operation}
	for _, number := range numbers {
		numberSize := unsafe.Sizeof(number)
		if err := binary.Write(buffer, binary.LittleEndian, int16(numberSize)); err != nil {
			return err
		}
		if err := binary.Write(buffer, binary.LittleEndian, number); err != nil {
			return err
		}
	}
	buffer.WriteByte(0)
	newData := buffer.Bytes()

	packetData, _ := writeBlock(w.packet.Data, blockNumber, newData)
	w.packet.Data = packetData
	return nil
}

