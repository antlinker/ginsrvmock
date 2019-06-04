// Copyright 2017 Bo-Yi Wu.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// +build !jsoniter

package json

import (
	"encoding/json"
)

var (
	// Marshal  json扩展
	Marshal = json.Marshal
	// MarshalIndent  json扩展
	MarshalIndent = json.MarshalIndent
	// NewDecoder  json扩展
	NewDecoder = json.NewDecoder
	// NewEncoder  json扩展
	NewEncoder = json.NewEncoder
	// Unmarshal json扩展
	Unmarshal = json.Unmarshal
)
