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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPacketWriterAndReader tests the packet writer and reader
func TestPacketWriterAndReader(t *testing.T) {
	assert := assert.New(t)

	packet := &Packet{}
	
	_, err := NewPacketWriter(packet)
	assert.Nil(err)
	// inputProtocol := &PacketProtocol{
	// 	Version:   int16(2 * i),
	// 	Operation: int16(10 * i),
	// }
	// err = writer.WriteProtocolSection(inputProtocol)
	// assert.Nil(err)

	// reader, err := NewPacketReader(packet)
	// assert.Nil(err)
	// var outputProtocol PacketProtocol
	// err = reader.ReadProtocolSection(&outputProtocol)
	// assert.Nil(err)

	// assert.Equal(inputProtocol.Version, outputProtocol.Version)
	// assert.Equal(inputProtocol.Operation, outputProtocol.Operation)
}
