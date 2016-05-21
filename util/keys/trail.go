// Copyright © 2016 Abcum Ltd
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

package keys

import (
	"fmt"
	"strings"
	"time"
)

// Trail ...
type Trail struct {
	KV   string      // KV
	NS   string      // NS
	DB   string      // DB
	TB   string      // TB
	TK   string      // •
	ID   interface{} // ID
	Time time.Time   // Time
}

// init initialises the key
func (k *Trail) init() *Trail {
	if k.TK == "" {
		k.TK = "~"
	}
	if k.Time.IsZero() {
		k.Time = time.Now()
	}
	return k
}

// Encode encodes the key into binary
func (k *Trail) Encode() []byte {
	k.init()
	return encode(k.KV, k.NS, k.DB, k.TB, k.TK, k.ID, k.Time)
}

// Decode decodes the key from binary
func (k *Trail) Decode(data []byte) {
	k.init()
	decode(data, &k.KV, &k.NS, &k.DB, &k.TB, &k.TK, &k.ID, &k.Time)
}

// String returns a string representation of the key
func (k *Trail) String() string {
	k.init()
	return "/" + strings.Join([]string{k.KV, k.NS, k.DB, k.TB, k.TK, fmt.Sprintf("%v", k.ID), k.Time.Format(time.RFC3339Nano)}, "/")
}