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
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestClientStateMachineExecution verifies the correct execution of the client state machine.
func TestClientStateMachineExecution(t *testing.T) {
    assert := assert.New(t)

    // Create a new state machine starting at the advertising state
    stateMachine, err := NewStateMachine(ClientAdvertisingState)
    assert.Nil(err, "Failed to initialize the state machine")

    // Execute the state machine
    err = stateMachine.Execute()
    assert.Nil(err, "State machine execution encountered an error")

    // assert.Equal(expectedFinalState, stateMachine.CurrentState(), "State machine did not reach the expected final state")
}
