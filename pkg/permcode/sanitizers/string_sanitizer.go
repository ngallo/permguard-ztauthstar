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
//
// SPDX-License-Identifier: Apache-2.0

package sanitizers

import (
	"strings"
)

// SanitizeString sanitizes a string.
func SanitizeString(value string) string {
	return strings.ToLower(value)
}

// SanitizeWilcardString sanitizes a wildcard string.
func SanitizeWilcardString(value string) string {
	sanitizedValue := SanitizeString(value)
	if len(value) == 0 {
		sanitizedValue = "*"
	}
	for strings.Contains(sanitizedValue, "**") {
		sanitizedValue = strings.ReplaceAll(sanitizedValue, "**", "*")
	}
	return sanitizedValue
}
