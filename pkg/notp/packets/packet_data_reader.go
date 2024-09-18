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
	"fmt"
)

// PacketReader is a reader of packets from the NOTP protocol.
type PacketReader struct {
	packet *Packet
}

// NewPacketReader creates a new packet reader.
func NewPacketReader(packet *Packet) (*PacketReader, error) {
	if packet == nil {
		return nil, errors.New("notp: nil packet")
	}
	return &PacketReader{
		packet: packet,
	}, nil
}

// ReadProtocolSection reads the protocol section.
func (r *PacketReader) ReadProtocolSection(protocol *Protocol) error {
	blockNumber := 0
	data, exist := seekBlock(r.packet.Data, blockNumber)
	if !exist {
		return fmt.Errorf("notp: block %d not found", blockNumber)
	}
	buf := bytes.NewBuffer(data)

	numbers := []int16{0, 0}
	for i := range numbers {
		var numberSize int16
		if err := binary.Read(buf, binary.LittleEndian, &numberSize); err != nil {
			return err
		}
		if err := binary.Read(buf, binary.LittleEndian, &numbers[i]); err != nil {
			return err
		}
	}

	protocol.Version = numbers[0]
	protocol.Operation = numbers[1]

	return nil
}

