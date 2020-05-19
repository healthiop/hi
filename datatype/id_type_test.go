// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package datatype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIDImplementsAccessor(t *testing.T) {
	o := NewID("Test")
	assert.Implements(t, (*IDAccessor)(nil), o)
}

func TestIDDataType(t *testing.T) {
	o := NewID("Test")
	dataType := o.DataType()
	assert.Equal(t, IDDataType, dataType)
}

func TestIDTypeInfo(t *testing.T) {
	o := NewID("Test")
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.id", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.string", i.FQBaseName().String())
		}
	}
}

func TestNewIDCollection(t *testing.T) {
	c := NewIDCollection()
	assert.Equal(t, "FHIR.id", c.ItemTypeInfo().String())
}

func TestIDNil(t *testing.T) {
	o := NewIDNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.Equal(t, "", o.String())
}

func TestIDInvalid(t *testing.T) {
	assert.Panics(t, func() { NewID(" Test ID") })
}

func TestIDValue(t *testing.T) {
	o := NewID("Test")
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.String()
	assert.Equal(t, "Test", value)
}

func TestParseID(t *testing.T) {
	o, err := ParseID("TestID")
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, o, "data type expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, "TestID", o.String())
	}
}

func TestParseIDInvalid(t *testing.T) {
	o, err := ParseID(" Test ID")
	assert.Error(t, err, "error expected")
	assert.Nil(t, o, "no object expected")
}

func TestIDEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewID("test").Equal(NewString("test")))
	assert.Equal(t, true, NewID("test").ValueEqual(NewString("test")))
	assert.Equal(t, true, NewID("test").ValueEquivalent(NewString("test")))
}
