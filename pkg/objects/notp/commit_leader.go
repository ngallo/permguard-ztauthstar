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

// LeaderStateMachine represents the leader's state machine.
type LeaderStateMachine struct {
	notpsmachine.StateMachine
}

// NewLeaderStateMachine creates and initializes a new state machine with the given initial state and transport layer.
func NewLeaderStateMachine(transportLayer *notptransport.TransportLayer) (*LeaderStateMachine, error) {
	leaderStateMachine := &LeaderStateMachine{}
	stateMachine, err := notpsmachine.NewStateMachine(leaderStateMachine.LeaderAdvertisingState, transportLayer)
	if err != nil {
		return nil, err
	}
	leaderStateMachine.StateMachine = *stateMachine
	return leaderStateMachine, nil
}

// LeaderAdvertisingState handles the leader's actions during the advertising phase.
func (s *LeaderStateMachine) LeaderAdvertisingState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, s.LeaderNegotiatingState, nil
}

// LeaderNegotiatingState handles the leader's actions during the negotiation phase.
func (s *LeaderStateMachine) LeaderNegotiatingState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, s.LeaderObjectExchangeState, nil
}

// LeaderObjectExchangeState handles the leader's actions during the object exchange phase.
func (s *LeaderStateMachine) LeaderObjectExchangeState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, notpsmachine.FinalState, nil
}
