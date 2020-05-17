// Copyright (c) 2020, Volker Schmidt (volker@volsch.eu)
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source URI must retain the above copyright notice, this
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

func TestURIImplementsAccessor(t *testing.T) {
	o := NewURI("test")
	assert.Implements(t, (*URIAccessor)(nil), o)
}

func TestURIDataType(t *testing.T) {
	o := NewURI("TestURI")
	dataType := o.DataType()
	assert.Equal(t, URIDataType, dataType)
}

func TestURITypeInfo(t *testing.T) {
	o := NewURI("TestURI")
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.uri", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
		}
	}
}

func TestNewURICollection(t *testing.T) {
	c := NewURICollection()
	assert.Equal(t, "FHIR.uri", c.ItemTypeInfo().String())
}

func TestURINil(t *testing.T) {
	o := NewURINil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, "", o.String())
}

func TestURIInvalid(t *testing.T) {
	assert.Panics(t, func() { NewURI(" Test URI") })
}

func TestParseURI(t *testing.T) {
	o, err := ParseURI("TestURI")
	assert.NoError(t, err, "no error expected")
	if assert.NotNil(t, o, "data type expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, "TestURI", o.String())
	}
}

func TestParseURIInvalid(t *testing.T) {
	o, err := ParseURI("Test URI")
	assert.Error(t, err, "error expected")
	assert.Nil(t, o, "no object expected")
}

func TestURIValue(t *testing.T) {
	o := NewURI("test")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	value := o.String()
	assert.Equal(t, "test", value)
}

func TestURIEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewURI("").Equal(newAccessorMock()))
	assert.Equal(t, false, NewURI("").ValueEqual(newAccessorMock()))
	assert.Equal(t, false, NewURI("").ValueEquivalent(newAccessorMock()))
}

func TestURIEqualStringTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewURI("test").Equal(NewString("test")))
	assert.Equal(t, false, NewURI("test").ValueEqual(NewString("test")))
	assert.Equal(t, false, NewURI("test").ValueEquivalent(NewString("test")))
}

func TestURIEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewURINil().Equal(NewURI("")))
	assert.Equal(t, false, NewURINil().ValueEqual(NewURI("")))
	assert.Equal(t, false, NewURINil().ValueEquivalent(NewURI("")))
}

func TestURIEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewURI("").Equal(NewURINil()))
	assert.Equal(t, false, NewURI("").ValueEqual(NewURINil()))
	assert.Equal(t, false, NewURI("").ValueEquivalent(NewURINil()))
}

func TestURIEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewURINil().Equal(NewURINil()))
	assert.Equal(t, true, NewURINil().ValueEqual(NewURINil()))
	assert.Equal(t, true, NewURINil().ValueEquivalent(NewURINil()))
}

func TestURIEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewURI("test").Equal(NewURI("test")))
	assert.Equal(t, true, NewURI("test").ValueEqual(NewURI("test")))
	assert.Equal(t, true, NewURI("test").ValueEquivalent(NewURI("test")))
}

func TestURIEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewURI("test1").Equal(NewURI("test2")))
	assert.Equal(t, false, NewURI("test1").ValueEqual(NewURI("test2")))
	assert.Equal(t, false, NewURI("test1").ValueEquivalent(NewURI("test2")))
}

func TestURIEqualNotEquivalent(t *testing.T) {
	assert.Equal(t, false, NewURI("TEST").ValueEquivalent(NewURI("test")))
}
