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

package notp

import (
	notpsmachine "github.com/permguard/permguard-notp-protocol/pkg/notp/statemachines"
	notptransport "github.com/permguard/permguard-notp-protocol/pkg/notp/transport"
)

// FollowerStateMachine represents the follower's state machine.
type FollowerStateMachine struct {
	notpsmachine.StateMachine
}

// NewFollowerStateMachine creates and initializes a new state machine with the given initial state and transport layer.
func NewFollowerStateMachine(transportLayer *notptransport.TransportLayer) (*FollowerStateMachine, error) {
	followerStateMachine := &FollowerStateMachine{}
	stateMachine, err := notpsmachine.NewStateMachine(followerStateMachine.FollowerAdvertisingState, transportLayer)
	if err != nil {
		return nil, err
	}
	followerStateMachine.StateMachine = *stateMachine
	return followerStateMachine, nil
}

// FollowerAdvertisingState handles the follower's actions during the advertising phase.
func (s *FollowerStateMachine) FollowerAdvertisingState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, s.FollowerNegotiatingState, nil
}

// FollowerNegotiatingState handles the follower's actions during the negotiation phase.
func (s *FollowerStateMachine) FollowerNegotiatingState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, s.FollowerObjectExchangeState, nil
}

// FollowerObjectExchangeState handles the follower's actions during the object exchange phase.
func (s *FollowerStateMachine) FollowerObjectExchangeState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, notpsmachine.FinalState, nil
}
