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

// SamplePacket represents a sample packet.
type SamplePacket struct {
	Text string
}

// GetType returns the type of the packet.
func (p *SamplePacket) GetType() int32 {
	return -1
}

// Serialize serializes the packet.
func (p *SamplePacket) Serialize() ([]byte, error) {
	return []byte(p.Text), nil
}

// Deserialize deserializes the packet.
func (p *SamplePacket) Deserialize(data []byte) error {
	p.Text = string(data)
	return nil
}

// TestPacketWriterAndReader tests the packet writer and reader
func TestPacketWriterAndReader(t *testing.T) {
	assert := assert.New(t)

	packet := &Packet{}

	writer, err := NewPacketWriter(packet)
	assert.Nil(err)

	protocol := &ProtocolPacket{ Version: 10 }
	err = writer.WriteProtocol(protocol)
	assert.Nil(err)

	data1 := &SamplePacket{ Text: "fd1d3938-2988-4df3-9b83-cc278b69cab0" }
	err = writer.AppendDataPacket(data1)
	assert.Nil(err)

	data2 := &SamplePacket{ Text: "3ecd7285-8406-4647-8e8f-92d87348636d" }
	err = writer.AppendDataPacket(data2)
	assert.Nil(err)

	data3 := &SamplePacket{ Text: "83ce2f5b-f5c4-4bd7-85de-69291f1f80d4" }
	err = writer.AppendDataPacket(data3)
	assert.Nil(err)
}
