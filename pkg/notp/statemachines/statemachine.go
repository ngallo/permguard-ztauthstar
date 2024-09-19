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

package statemachines

import (
	"errors"
)

// StateFunc represents a function that returns the next state function.
type StateFunc func(*StateMachineRuntime) (bool, StateFunc, error)

// stateInitial represents the initial state of the state machine.
func stateInitial(runtime *StateMachineRuntime) (bool, StateFunc, error) {
	return false, runtime.initialState, nil
}

// stateFInal represents the final state of the state machine.
func stateFinal(runtime *StateMachineRuntime) (bool, StateFunc, error) {
	return true, nil, nil
}

// StateMachineRuntime represents the runtime of a state machine.
type StateMachineRuntime struct {
	initialState StateFunc
}

// StateMachine represents a state machine.
type StateMachine struct {
	runtime *StateMachineRuntime
}

// Run runs the state machine.
func (m *StateMachine) Run() error {
    state := m.runtime.initialState
    var err error
	var final bool
    for state != nil {
        final, state, err = state(m.runtime)
        if err != nil {
            return err
        } else if final {
			break
		}
    }
    return nil
}

// NewStateMachine creates a new state machine.
func NewStateMachine(initialState StateFunc) (*StateMachine, error) {
	if initialState == nil {
		return nil, errors.New("notp: initial state cannot be nil")
	}
	return &StateMachine{
		runtime: &StateMachineRuntime{
			initialState: initialState,
		},
	}, nil
}
