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

// Package transport implements the transport layer of the NOTP protocol.
package transport

import (
	"errors"

	notppackets "github.com/permguard/permguard-abs-language/pkg/notp/packets"
)

// PacketSender defines the interface for sending packets over the transport layer.
type PacketSender func(packet *notppackets.Packet) error

// PacketReceiver defines the interface for receiving packets from the transport layer.
type PacketReceiver func() (*notppackets.Packet, error)

// TransportLayer represents the transport layer responsible for packet transmission in the NOTP protocol.
type TransportLayer struct {
	PacketSender   PacketSender
	PacketReceiver PacketReceiver
}

// TransmitPacket sends a packet through the transport layer.
func (t *TransportLayer) TransmitPacket(packet *notppackets.Packet) error {
	if t.PacketSender == nil {
		return errors.New("notp: transport layer does not have a defined packet sender")
	}
	return t.PacketSender(packet)
}

// ReceivePacket retrieves a packet from the transport layer.
func (t *TransportLayer) ReceivePacket() (*notppackets.Packet, error) {
	if t.PacketReceiver == nil {
		return nil, errors.New("notp: transport layer does not have a defined packet receiver")
	}
	return t.PacketReceiver()
}

// NewTransportLayer creates and initializes a new transport layer.
func NewTransportLayer(packetSender PacketSender, packetReceiver PacketReceiver) (*TransportLayer, error) {
	if packetSender == nil {
		return nil, errors.New("notp: PacketSender cannot be nil")
	}
	if packetReceiver == nil {
		return nil, errors.New("notp: PacketReceiver cannot be nil")
	}
	return &TransportLayer{
		PacketSender:   packetSender,
		PacketReceiver: packetReceiver,
	}, nil
}
