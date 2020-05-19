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

func TestUnsignedIntImplementsAccessor(t *testing.T) {
	o := NewUnsignedInt(4711)
	assert.Implements(t, (*UnsignedIntAccessor)(nil), o)
}

func TestUnsignedIntDataType(t *testing.T) {
	o := NewUnsignedInt(4711)
	dataType := o.DataType()
	assert.Equal(t, UnsignedIntDataType, dataType)
}

func TestUnsignedIntTypeInfo(t *testing.T) {
	o := NewUnsignedInt(0)
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.unsignedInt", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.integer", i.FQBaseName().String())
		}
	}
}

func TestNewUnsignedIntCollection(t *testing.T) {
	c := NewUnsignedIntCollection()
	assert.Equal(t, "FHIR.unsignedInt", c.ItemTypeInfo().String())
}

func TestUnsignedIntNil(t *testing.T) {
	o := NewUnsignedIntNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.Equal(t, int32(0), o.Int())
	assert.Equal(t, "", o.String())
}

func TestUnsignedIntIsNegative(t *testing.T) {
	assert.Panics(t, func() { NewUnsignedInt(-1) })
}

func TestUnsignedIntValue(t *testing.T) {
	o := NewUnsignedInt(4711)
	assert.False(t, o.Nil(), "nil data type expected")
	value := o.Int()
	assert.Equal(t, int32(4711), value)
}

func TestUnsignedIntValueIsZero(t *testing.T) {
	o := NewUnsignedInt(0)
	assert.False(t, o.Nil(), "nil data type expected")
	value := o.Int()
	assert.Equal(t, int32(0), value)
}

func TestUnsignedIntEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewUnsignedInt(0).Equal(newAccessorMock()))
	assert.Equal(t, false, NewUnsignedInt(0).ValueEqual(newAccessorMock()))
	assert.Equal(t, false, NewUnsignedInt(0).ValueEquivalent(newAccessorMock()))
}

func TestUnsignedIntEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewUnsignedIntNil().Equal(NewUnsignedInt(0)))
	assert.Equal(t, false, NewUnsignedIntNil().ValueEqual(NewUnsignedInt(0)))
	assert.Equal(t, false, NewUnsignedIntNil().ValueEquivalent(NewUnsignedInt(0)))
}

func TestUnsignedIntEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewUnsignedInt(0).Equal(NewUnsignedIntNil()))
	assert.Equal(t, false, NewUnsignedInt(0).ValueEqual(NewUnsignedIntNil()))
	assert.Equal(t, false, NewUnsignedInt(0).ValueEquivalent(NewUnsignedIntNil()))
}

func TestUnsignedIntEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewUnsignedIntNil().Equal(NewUnsignedIntNil()))
	assert.Equal(t, true, NewUnsignedIntNil().ValueEqual(NewUnsignedIntNil()))
	assert.Equal(t, true, NewUnsignedIntNil().ValueEquivalent(NewUnsignedIntNil()))
}

func TestUnsignedIntEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewUnsignedInt(8274).Equal(NewUnsignedInt(8274)))
	assert.Equal(t, true, NewUnsignedInt(8274).ValueEqual(NewUnsignedInt(8274)))
	assert.Equal(t, true, NewUnsignedInt(8274).ValueEquivalent(NewUnsignedInt(8274)))
}

func TestUnsignedIntEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewUnsignedInt(8274).Equal(NewUnsignedInt(8275)))
	assert.Equal(t, false, NewUnsignedInt(8274).ValueEqual(NewUnsignedInt(8275)))
	assert.Equal(t, false, NewUnsignedInt(8274).ValueEquivalent(NewUnsignedInt(8275)))
}

func TestUnsignedIntEqualInteger(t *testing.T) {
	assert.Equal(t, true, NewUnsignedInt(8274).Equal(NewInteger(8274)))
	assert.Equal(t, true, NewUnsignedInt(8274).ValueEqual(NewInteger(8274)))
	assert.Equal(t, true, NewUnsignedInt(8274).ValueEquivalent(NewInteger(8274)))
}
