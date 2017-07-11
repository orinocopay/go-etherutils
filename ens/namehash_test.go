// Copyright 2017 Orinoco Payments
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

package ens

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameHashEmpty(t *testing.T) {
	expected := "0000000000000000000000000000000000000000000000000000000000000000"
	actual := NameHash("")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestNameHashTLD(t *testing.T) {
	expected := "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"
	actual := NameHash("eth")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestNameHashFooEth(t *testing.T) {
	expected := "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"
	actual := NameHash("foo.eth")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestNameHashFooDotEth(t *testing.T) {
	expected := "4143a5b2f547838d3b49982e3f2ec6a26415274e5b9c3ffeb21971bbfdfaa052"
	actual := NameHash("foo..eth")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestNameHashNickJohnsonEth(t *testing.T) {
	expected := "25cfe90ad9477590acf268bb3ad00ab18465ecce12760be3d8eac81c9f329995"
	actual := NameHash("nickjohnson.eth")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}
