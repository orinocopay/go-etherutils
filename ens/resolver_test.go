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

func TestResolveEmpty(t *testing.T) {
	_, err := Resolve(client, "", rpcclient)
	assert.NotNil(t, err, "Resolved empty name")
}

func TestResolveNotPresent(t *testing.T) {
	_, err := Resolve(client, "sirnotappearinginthisregistry.eth", rpcclient)
	assert.NotNil(t, err, "Resolved name that does not exist")
	assert.Equal(t, "unregistered name", err.Error(), "Unexpected error")
}

func TestResolveNoResolver(t *testing.T) {
	_, err := Resolve(client, "noresolver.eth", rpcclient)
	assert.NotNil(t, err, "Resolved name without a resolver")
	assert.Equal(t, "no resolver", err.Error(), "Unexpected error")
}

func TestResolveBadResolver(t *testing.T) {
	_, err := Resolve(client, "resolvestozero.eth", rpcclient)
	assert.NotNil(t, err, "Resolved name with a bad resolver")
	assert.Equal(t, "no address", err.Error(), "Unexpected error")
}

func TestResolveTestEnsTest(t *testing.T) {
	expected := "a34c6bcae6f46ac6470443ccea67d937f6060c7e"
	actual, err := Resolve(client, "test.enstest.eth", rpcclient)
	assert.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}

func TestResolveNickJohnson(t *testing.T) {
	expected := "fdb33f8ac7ce72d7d4795dd8610e323b4c122fbb"
	actual, err := Resolve(client, "nickjohnson.eth", rpcclient)
	assert.Nil(t, err, "Error resolving name")
	assert.Equal(t, expected, hex.EncodeToString(actual[:]), "Did not receive expected result")
}
