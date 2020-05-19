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

func TestCodeImplementsAccessor(t *testing.T) {
	o := NewCode("Test")
	assert.Implements(t, (*CodeAccessor)(nil), o)
}

func TestCodeDataType(t *testing.T) {
	o := NewCode("Test Code")
	dataType := o.DataType()
	assert.Equal(t, CodeDataType, dataType)
}

func TestCodeTypeInfo(t *testing.T) {
	o := NewCode("Test Code")
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.code", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.string", i.FQBaseName().String())
		}
	}
}

func TestCodeInvalid(t *testing.T) {
	assert.Panics(t, func() { NewCode(" Test Code") })
}

func TestNewCodeCollection(t *testing.T) {
	c := NewCodeCollection()
	assert.Equal(t, "FHIR.code", c.ItemTypeInfo().String())
}

func TestCodeNil(t *testing.T) {
	o := NewCodeNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.Equal(t, "", o.String())
}

func TestParseCode(t *testing.T) {
	o, err := ParseCode("Test Code")
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, o, "data type expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, "Test Code", o.String())
	}
}

func TestParseCodeInvalid(t *testing.T) {
	o, err := ParseCode(" Test Code")
	assert.Error(t, err, "error expected")
	assert.Nil(t, o, "no object expected")
}

func TestCodeValue(t *testing.T) {
	o := NewCode("Test")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, "Test", o.String())
}

func TestCodeEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewCode("test").Equal(NewString("test")))
	assert.Equal(t, true, NewCode("test").ValueEqual(NewString("test")))
	assert.Equal(t, true, NewCode("test").ValueEquivalent(NewString("test")))
}
