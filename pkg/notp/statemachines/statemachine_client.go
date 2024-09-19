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

// stateClientAdvertising represents the state of the client during the advertising phase.
func stateClientAdvertising(runtime *StateMachineRuntime) (bool, StateFunc, error) {
	return false, stateClientNegotiating, nil
}

// stateClientNegotiating represents the state of the client during the negotiation phase.
func stateClientNegotiating(runtime *StateMachineRuntime) (bool, StateFunc, error) {
	return false, stateClientObjectExchange, nil
}

// stateClientObjectExchange represents the state of the client during the object exchange phase.
func stateClientObjectExchange(runtime *StateMachineRuntime) (bool, StateFunc, error) {
	return false, stateFinal, nil
}
