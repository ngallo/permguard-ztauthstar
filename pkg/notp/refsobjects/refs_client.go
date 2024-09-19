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
	notppackets "github.com/permguard/permguard-abs-language/pkg/notp/packets"
	notpsmachine "github.com/permguard/permguard-abs-language/pkg/notp/statemachines"
)

// ClientAdvertisingState handles the client's actions during the advertising phase.
func ClientAdvertisingState(runtime *notpsmachine.StateMachineRuntimeContext) (isFinal bool, nextState notpsmachine.StateTransitionFunc, err error) {
	str := "ADVERTISING"
	packet := &RefsObjPacket{
		Packet: notppackets.Packet{
			Data: []byte(str),
		},
	}
	runtime.TransmitPacket(&packet.Packet)
	return false, ClientNegotiatingState, nil
}

// ClientNegotiatingState handles the client's actions during the negotiation phase.
func ClientNegotiatingState(runtime *notpsmachine.StateMachineRuntimeContext) (isFinal bool, nextState notpsmachine.StateTransitionFunc, err error) {
	str := "NEGOTIATING"
	packet := &RefsObjPacket{
		Packet: notppackets.Packet{
			Data: []byte(str),
		},
	}
	runtime.TransmitPacket(&packet.Packet)
	return false, ClientObjectExchangeState, nil
}

// ClientObjectExchangeState handles the client's actions during the object exchange phase.
func ClientObjectExchangeState(runtime *notpsmachine.StateMachineRuntimeContext) (isFinal bool, nextState notpsmachine.StateTransitionFunc, err error) {
	str := "OBJECT_EXCHANGE"
	packet := &RefsObjPacket{
		Packet: notppackets.Packet{
			Data: []byte(str),
		},
	}
	runtime.TransmitPacket(&packet.Packet)
	return false, notpsmachine.FinalState, nil
}
