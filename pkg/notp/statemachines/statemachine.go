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

// StateTransitionFunc defines a function responsible for transitioning to the next state in the state machine.
type StateTransitionFunc func(runtime *StateMachineExecutionContext) (isFinal bool, nextState StateTransitionFunc, err error)

// InitialState defines the initial state of the state machine.
func InitialState(runtime *StateMachineExecutionContext) (bool, StateTransitionFunc, error) {
    return false, runtime.InitialState, nil
}

// FinalState defines the final state of the state machine.
func FinalState(runtime *StateMachineExecutionContext) (bool, StateTransitionFunc, error) {
    return true, nil, nil
}

// StateMachineExecutionContext holds the execution context of the state machine.
type StateMachineExecutionContext struct {
    InitialState StateTransitionFunc
}

// StateMachine orchestrates the execution of state transitions.
type StateMachine struct {
    executionContext *StateMachineExecutionContext
}

// Execute starts and runs the state machine through its states until termination.
func (m *StateMachine) Execute() error {
    state := m.executionContext.InitialState
    var err error
    var isFinal bool
    for state != nil {
        isFinal, state, err = state(m.executionContext)
        if err != nil {
            return err
        } else if isFinal {
            break
        }
    }
    return nil
}

// NewStateMachine creates and initializes a new state machine with the given initial state.
func NewStateMachine(initialState StateTransitionFunc) (*StateMachine, error) {
    if initialState == nil {
        return nil, errors.New("notp: initial state cannot be nil")
    }
    return &StateMachine{
        executionContext: &StateMachineExecutionContext{
            InitialState: initialState,
        },
    }, nil
}
