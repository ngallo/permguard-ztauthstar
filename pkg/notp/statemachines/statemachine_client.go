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

// ClientAdvertisingState handles the client's actions during the advertising phase.
func ClientAdvertisingState(runtime *StateMachineExecutionContext) (isFinal bool, nextState StateTransitionFunc, err error) {
    return false, ClientNegotiatingState, nil
}

// ClientNegotiatingState handles the client's actions during the negotiation phase.
func ClientNegotiatingState(runtime *StateMachineExecutionContext) (isFinal bool, nextState StateTransitionFunc, err error) {
    return false, ClientObjectExchangeState, nil
}

// ClientObjectExchangeState handles the client's actions during the object exchange phase.
func ClientObjectExchangeState(runtime *StateMachineExecutionContext) (isFinal bool, nextState StateTransitionFunc, err error) {
    return false, FinalState, nil
}
