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

func TestPositiveIntDataType(t *testing.T) {
	o := NewPositiveInt(4711)
	dataType := o.DataType()
	assert.Equal(t, PositiveIntDataType, dataType)
}

func TestPositiveIntTypeInfo(t *testing.T) {
	o := NewPositiveInt(1)
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.positiveInt", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.integer", i.FQBaseName().String())
		}
	}
}

func TestPositiveIntNil(t *testing.T) {
	o := NewPositiveIntNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.Equal(t, int32(1), o.Int())
	assert.Equal(t, "", o.String())
}

func TestPositiveIntIsZero(t *testing.T) {
	assert.Panics(t, func() { NewPositiveInt(0) })
}

func TestPositiveIntValue(t *testing.T) {
	o := NewPositiveInt(4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.Int()
	assert.Equal(t, int32(4711), value)
}

func TestPositiveIntEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewPositiveInt(1).Equal(newAccessorMock()))
	assert.Equal(t, false, NewPositiveInt(1).Equivalent(newAccessorMock()))
}

func TestPositiveIntEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewPositiveIntNil().Equal(NewPositiveInt(1)))
	assert.Equal(t, false, NewPositiveIntNil().Equivalent(NewPositiveInt(1)))
}

func TestPositiveIntEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewPositiveInt(1).Equal(NewPositiveIntNil()))
	assert.Equal(t, false, NewPositiveInt(1).Equivalent(NewPositiveIntNil()))
}

func TestPositiveIntEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewPositiveIntNil().Equal(NewPositiveIntNil()))
	assert.Equal(t, true, NewPositiveIntNil().Equivalent(NewPositiveIntNil()))
}

func TestPositiveIntEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewPositiveInt(8274).Equal(NewPositiveInt(8274)))
	assert.Equal(t, true, NewPositiveInt(8274).Equivalent(NewPositiveInt(8274)))
}

func TestPositiveIntEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewPositiveInt(8274).Equal(NewPositiveInt(8275)))
	assert.Equal(t, false, NewPositiveInt(8274).Equivalent(NewPositiveInt(8275)))
}

func TestPositiveIntEqualInteger(t *testing.T) {
	assert.Equal(t, true, NewPositiveInt(8274).Equal(NewInteger(8274)))
	assert.Equal(t, true, NewPositiveInt(8274).Equivalent(NewInteger(8274)))
}
