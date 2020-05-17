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
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestIntegerImplementsAccessor(t *testing.T) {
	o := NewInteger(4711)
	assert.Implements(t, (*IntegerAccessor)(nil), o)
}

func TestIntegerImplementsNegator(t *testing.T) {
	o := NewInteger(4711)
	assert.Implements(t, (*Negator)(nil), o)
}

func TestIntegerDataType(t *testing.T) {
	o := NewInteger(4711)
	dataType := o.DataType()
	assert.Equal(t, IntegerDataType, dataType)
}

func TestIntegerTypeInfo(t *testing.T) {
	o := NewInteger(0)
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.integer", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
		}
	}
}

func TestNewIntegerCollection(t *testing.T) {
	c := NewIntegerCollection()
	assert.Equal(t, "FHIR.integer", c.ItemTypeInfo().String())
}

func TestIntegerNil(t *testing.T) {
	o := NewIntegerNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, int32(0), o.Int())
	assert.Equal(t, "", o.String())
}

func TestIntegerValue(t *testing.T) {
	o := NewInteger(-4711)
	assert.True(t, IsNumber(o), "an integer is a number")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	assert.Equal(t, int32(-4711), o.Int())
	assert.Equal(t, "-4711", o.String())
}

func TestInteger64Value(t *testing.T) {
	o := NewInteger(-4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.Int64()
	assert.Equal(t, int64(-4711), value)
}

func TestIntegerFloat32Value(t *testing.T) {
	o := NewInteger(-4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.Float32()
	assert.Equal(t, float32(-4711), value)
}

func TestIntegerFloat64Value(t *testing.T) {
	o := NewInteger(-4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.Float64()
	assert.Equal(t, float64(-4711), value)
}

func TestIntegerBigFloatValue(t *testing.T) {
	o := NewInteger(-4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.BigFloat()
	assert.Equal(t, big.NewFloat(-4711), value)
}

func TestIntegerDecimalValue(t *testing.T) {
	o := NewInteger(-4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	value := o.Decimal()
	expected := decimal.NewFromInt32(-4711)
	assert.True(t, expected.Equal(value), "expected %s, got %s", expected.String(), value.String())
}

func TestParseInteger(t *testing.T) {
	o, err := ParseInteger("-83628")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, int32(-83628), o.Int())
	}
}

func TestParseIntegerInvalid(t *testing.T) {
	o, err := ParseInteger("8273.3")
	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}

func TestIntegerNegatePos(t *testing.T) {
	o := NewInteger(8372)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, int32(8372), o.Int())
	if assert.Implements(t, (*IntegerAccessor)(nil), n) {
		assert.False(t, n.(IntegerAccessor).Nil(), "non-nil data type expected")
		assert.Equal(t, int32(-8372), n.(IntegerAccessor).Int())
	}
}

func TestIntegerNegateNeg(t *testing.T) {
	o := NewInteger(-8372)
	n := o.Negate()
	assert.NotSame(t, o, n)
	assert.Equal(t, int32(-8372), o.Int())
	if assert.Implements(t, (*IntegerAccessor)(nil), n) {
		assert.False(t, n.(IntegerAccessor).Nil(), "non-nil data type expected")
		assert.Equal(t, int32(8372), n.(IntegerAccessor).Int())
	}
}

func TestIntegerNegateNil(t *testing.T) {
	o := NewIntegerNil()
	n := o.Negate()
	assert.Same(t, o, n)
}

func TestIntegerEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewInteger(0).Equal(newAccessorMock()))
	assert.Equal(t, false, NewInteger(0).ValueEqual(newAccessorMock()))
	assert.Equal(t, false, NewInteger(0).ValueEquivalent(newAccessorMock()))
}

func TestIntegerEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewIntegerNil().Equal(NewInteger(0)))
	assert.Equal(t, false, NewIntegerNil().ValueEqual(NewInteger(0)))
	assert.Equal(t, false, NewIntegerNil().ValueEquivalent(NewInteger(0)))
}

func TestIntegerEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewInteger(0).Equal(NewIntegerNil()))
	assert.Equal(t, false, NewInteger(0).ValueEqual(NewIntegerNil()))
	assert.Equal(t, false, NewInteger(0).ValueEquivalent(NewIntegerNil()))
}

func TestIntegerEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewIntegerNil().Equal(NewIntegerNil()))
	assert.Equal(t, true, NewIntegerNil().ValueEqual(NewIntegerNil()))
	assert.Equal(t, true, NewIntegerNil().ValueEquivalent(NewIntegerNil()))
}

func TestIntegerEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewInteger(8274).Equal(NewInteger(8274)))
	assert.Equal(t, true, NewInteger(8274).ValueEqual(NewInteger(8274)))
	assert.Equal(t, true, NewInteger(8274).ValueEquivalent(NewInteger(8274)))
}

func TestIntegerEqualEqualDecimal(t *testing.T) {
	assert.Equal(t, true, NewInteger(8274).Equal(NewDecimalInt(8274)))
	assert.Equal(t, true, NewInteger(8274).ValueEqual(NewDecimalInt(8274)))
	assert.Equal(t, true, NewInteger(8274).ValueEquivalent(NewDecimalInt(8274)))
}

func TestIntegerEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewInteger(8274).Equal(NewInteger(8275)))
	assert.Equal(t, false, NewInteger(8274).ValueEqual(NewInteger(8275)))
	assert.Equal(t, false, NewInteger(8274).ValueEquivalent(NewInteger(8275)))
}

func TestIntegerEqualNotEqualDecimal(t *testing.T) {
	assert.Equal(t, false, NewInteger(8274).Equal(NewDecimalInt(8275)))
	assert.Equal(t, false, NewInteger(8274).ValueEqual(NewDecimalInt(8275)))
	assert.Equal(t, false, NewInteger(8274).ValueEquivalent(NewDecimalInt(8275)))
}

func TestIntegerEquivalent(t *testing.T) {
	assert.Equal(t, true, NewInteger(8274).ValueEquivalent(NewDecimalFloat64(8274.8237)))
}
