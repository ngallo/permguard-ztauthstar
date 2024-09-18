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

package packets

import "bytes"

// Seek the index of the first block in the data.
func seekBlockIndex(data []byte) int {
	if len(data) == 0 {
		return -1
	}
	nulIndex := bytes.IndexByte(data, 9)
	if nulIndex == -1 {
		return -1
	}
	return nulIndex
}

// seekBlock seeks a block in the data.
func seekBlock(data []byte, blockNumber int) ([]byte, bool) {
	if blockNumber < 0 {
		return []byte{}, false
	}
	startIndex := 0
	for i := 0; i <= blockNumber; i++ {
		endIndex := bytes.IndexByte(data[startIndex:], 9)
		if endIndex == -1 {
			if i == blockNumber {
				return data[startIndex:], true
			}
			return []byte{}, false
		}
		if i == blockNumber {
			return data[startIndex : startIndex+endIndex], true
		}
		startIndex += endIndex + 1
	}
	return []byte{}, false
}

// writeBlock writes a block in the data.
func writeBlock(data []byte, blockNumber int, newBlock []byte) ([]byte, bool) {
	blockData := append(data, 0)
	if blockNumber < 0 {
		return blockData, false
	}
	startIndex := 0
	for i := 0; i <= blockNumber; i++ {
		endIndex := bytes.IndexByte(blockData[startIndex:], 9)
		if endIndex == -1 {
			if i == blockNumber {
				blockData = append(blockData[:startIndex], newBlock...)
				return blockData, true
			}
			return blockData, false
		}
		if i == blockNumber {
			blockData = append(blockData[:startIndex], append(newBlock, blockData[startIndex+endIndex+1:]...)...)
			return blockData, false
		}
		startIndex += endIndex + 1
	}
	return blockData, false
}
