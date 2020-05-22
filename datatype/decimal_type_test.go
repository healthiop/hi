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
	"testing"
)

func TestDecimalImplementsAccessor(t *testing.T) {
	o := NewDecimalFloat64(4711.10)
	assert.Implements(t, (*DecimalAccessor)(nil), o)
}

func TestDecimalDataType(t *testing.T) {
	o := NewDecimalFloat64(4711.10)
	dataType := o.DataType()
	assert.Equal(t, DecimalDataType, dataType)
}

func TestDecimalTypeInfo(t *testing.T) {
	o := NewDecimalInt(0)
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.decimal", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
		}
	}
}

func TestDecimalNil(t *testing.T) {
	o := NewDecimalNil()
	assert.True(t, o.Nil(), "nil data type expected")
	assert.True(t, o.Empty(), "nil data type expected")
	assert.Equal(t, float64(0), o.Float64())
	assert.Equal(t, "", o.String())
}

func TestNewDecimal(t *testing.T) {
	o := NewDecimal(decimal.NewFromFloat(-4711.12))
	assert.True(t, IsNumber(o), "a decimal is a number")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	assert.Equal(t, -4711.12, o.Float64())
	assert.Equal(t, "-4711.12", o.String())
}

func TestNewDecimalInt(t *testing.T) {
	o := NewDecimalInt(-4711)
	assert.True(t, IsNumber(o), "a decimal is a number")
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.False(t, o.Empty(), "non-nil data type expected")
	assert.Equal(t, -4711.0, o.Float64())
	assert.Equal(t, "-4711", o.String())
}

func TestDecimalInt(t *testing.T) {
	o := NewDecimalFloat64(-4711.831)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, int32(-4711), o.Int())
}

func TestNewDecimalInt64(t *testing.T) {
	o := NewDecimalInt64(-4711)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, -4711.0, o.Float64())
}

func TestDecimalInt64(t *testing.T) {
	o := NewDecimalFloat64(-4711.831)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, int64(-4711), o.Int64())
}

func TestNewDecimalFloat32(t *testing.T) {
	o := NewDecimalFloat64(-4711.6)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, float32(-4711.6), o.Float32())
	assert.Equal(t, "-4711.6", o.String())
}

func TestNewDecimalFloat64(t *testing.T) {
	o := NewDecimalFloat64(-4711.678121)
	assert.False(t, o.Nil(), "non-nil data type expected")
	assert.Equal(t, -4711.678121, o.Float64())
}

func TestDecimalBigFloat(t *testing.T) {
	o, err := ParseDecimal("-4711.83123200")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, "-4711.831232", o.BigFloat().String())
	}
}

func TestDecimalDecimal(t *testing.T) {
	o, err := ParseDecimal("-4711.831232753400")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, int32(-12), o.Decimal().Exponent())
		assert.Equal(t, "-4711.831232753400", o.String())
	}
}

func TestParseDecimal(t *testing.T) {
	o, err := ParseDecimal("-83628.85")
	assert.Nil(t, err, "no error expected")
	if assert.NotNil(t, o, "value expected") {
		assert.False(t, o.Nil(), "non-nil data type expected")
		assert.Equal(t, -83628.85, o.Float64())
	}
}

func TestParseDecimalInvalid(t *testing.T) {
	o, err := ParseDecimal("82737u83")
	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}

func TestDecimalEqualNil(t *testing.T) {
	assert.Equal(t, false, NewDecimalInt(0).Equal(nil))
}

func TestDecimalEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewDecimalInt(0).Equal(newAccessorMock()))
	assert.Equal(t, false, NewDecimalInt(0).Equivalent(newAccessorMock()))
}

func TestDecimalEqualLeftNil(t *testing.T) {
	assert.Equal(t, false, NewDecimalNil().Equal(NewDecimalInt(0)))
	assert.Equal(t, false, NewDecimalNil().Equivalent(NewDecimalInt(0)))
}

func TestDecimalEqualRightNil(t *testing.T) {
	assert.Equal(t, false, NewDecimalInt(0).Equal(NewDecimalNil()))
	assert.Equal(t, false, NewDecimalInt(0).Equivalent(NewDecimalNil()))
}

func TestDecimalEqualBothNil(t *testing.T) {
	assert.Equal(t, true, NewDecimalNil().Equal(NewDecimalNil()))
	assert.Equal(t, true, NewDecimalNil().Equivalent(NewDecimalNil()))
}

func TestDecimalEqualEqual(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.21).Equal(NewDecimalFloat64(8274.21)))
	assert.Equal(t, true, NewDecimalFloat64(8274.21).Equivalent(NewDecimalFloat64(8274.21)))
}

func TestDecimalEqualEqualInteger(t *testing.T) {
	assert.Equal(t, false, NewDecimalFloat64(8274).Equal(NewInteger(8274)))
	assert.Equal(t, true, NewDecimalFloat64(8274).Equivalent(NewInteger(8274)))
}

func TestDecimalEqualNotEqual(t *testing.T) {
	assert.Equal(t, false, NewDecimalFloat64(8274.21).Equal(NewDecimalFloat64(8274.22)))
	assert.Equal(t, false, NewDecimalFloat64(8274.21).Equivalent(NewDecimalFloat64(8274.22)))
}

func TestDecimalEqualNotEqualInteger(t *testing.T) {
	assert.Equal(t, false, NewDecimalFloat64(8274).Equal(NewInteger(8275)))
	assert.Equal(t, false, NewDecimalFloat64(8274).Equivalent(NewInteger(8275)))
}

func TestDecimalEqualNotEqualQuantity(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(64.14), nil, nil, nil, NewCode("cm"))
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equal(q))
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equivalent(q))
}

func TestDecimalEqualNotEqualQuantityNil(t *testing.T) {
	q := NewQuantity(nil, nil, nil, nil, NewCode("cm"))
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equal(q))
	assert.Equal(t, false, NewDecimalFloat64(64.12).Equivalent(q))
}

func TestDecimalEquivalentLeft(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.6).Equivalent(NewDecimalFloat64(8274.67)))
}

func TestDecimalEquivalentRight(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.67).Equivalent(NewDecimalFloat64(8274.6)))
}

func TestDecimalEquivalentInteger(t *testing.T) {
	assert.Equal(t, true, NewDecimalFloat64(8274.61).Equivalent(NewInteger(8274)))
}
