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
	notpsmachine "github.com/permguard/permguard-abs-language/pkg/notp/statemachines"
)

// ServerAdvertisingState handles the server's actions during the advertising phase.
func ServerAdvertisingState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, ServerNegotiatingState, nil
}

// ServerNegotiatingState handles the server's actions during the negotiation phase.
func ServerNegotiatingState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, ServerObjectExchangeState, nil
}

// ServerObjectExchangeState handles the server's actions during the object exchange phase.
func ServerObjectExchangeState(runtime *notpsmachine.StateMachineRuntimeContext) (bool, notpsmachine.StateTransitionFunc, error) {
	return false, notpsmachine.FinalState, nil
}
