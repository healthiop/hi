// Copyright (c) 2020-2021, Volker Schmidt (volker@volsch.eu)
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

func TestStringDataType(t *testing.T) {
	o := NewString("Test")
	dataType := o.DataType()
	assert.Equal(t, StringDataType, dataType)
}

func TestStringIsNotNumber(t *testing.T) {
	o := NewString("Test")
	assert.False(t, IsNumber(o), "a string is no number")
}

func TestStringTypeSpec(t *testing.T) {
	o := NewString("Test")
	i := o.TypeSpec()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.string", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
		}
	}
}

func TestStringNil(t *testing.T) {
	o := NewStringNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, "", o.String())
}

func TestStringInvalid(t *testing.T) {
	assert.Panics(t, func() { NewString(" Test\u0005String") })
}

func TestNewStringUnchecked(t *testing.T) {
	o := NewStringUnchecked(" Test\u0005String")
	assert.Equal(t, " Test\u0005String", o.String())
}

func TestParseString(t *testing.T) {
	o, err := ParseString("Test String\r\n")
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, o, "data type expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, "Test String\r\n", o.String())
	}
}

func TestParseStringInvalid(t *testing.T) {
	o, err := ParseString("Test\u0005String")
	assert.Error(t, err, "error expected")
	assert.Nil(t, o, "no object expected")
}

func TestStringValue(t *testing.T) {
	o := NewString("Test")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	value := o.String()
	assert.Equal(t, "Test", value)
}

func TestStringEqualNil(t *testing.T) {
	assert.Equal(t, false, NewString("").Equal(nil))
}

func TestStringEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewString("").Equal(newAccessorMock()))
	assert.Equal(t, false, NewString("").Equivalent(newAccessorMock()))
}

func TestStringEqualStringTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewString("test").Equal(NewURI("test")))
	assert.Equal(t, false, NewString("test").Equivalent(NewURI("test")))
}

func TestStringEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewStringNil().Equal(NewString("")))
	assert.Equal(t, false, NewStringNil().Equivalent(NewString("")))
}

func TestStringEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewString("").Equal(NewStringNil()))
	assert.Equal(t, false, NewString("").Equivalent(NewStringNil()))
}

func TestStringEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewStringNil().Equal(NewStringNil()))
	assert.Equal(t, true, NewStringNil().Equivalent(NewStringNil()))
}

func TestStringEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewString("test").Equal(NewString("test")))
	assert.Equal(t, true, NewString("test").Equivalent(NewString("test")))
}

func TestStringEquivalent(t *testing.T) {
	assert.Equal(t, true, NewString("TEST\nvalue").Equivalent(NewString("test VALUE")))
}

func TestStringEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewString("test1").Equal(NewString("test2")))
	assert.Equal(t, false, NewString("test1").Equivalent(NewString("test2")))
}
