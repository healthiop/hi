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

func TestIntegerImplementsAccessor(t *testing.T) {
	o := NewInteger(4711)
	assert.Implements(t, (*IntegerAccessor)(nil), o)
}

func TestIntegerDataType(t *testing.T) {
	o := NewInteger(4711)
	dataType := o.DataType()
	assert.Equal(t, IntegerDataType, dataType)
}

func TestIntegerValue(t *testing.T) {
	o := NewInteger(-4711)
	value := o.Int()
	assert.Equal(t, int32(-4711), value)
}

func TestInteger64Value(t *testing.T) {
	o := NewInteger(-4711)
	value := o.Int64()
	assert.Equal(t, int64(-4711), value)
}

func TestIntegerFloat32Value(t *testing.T) {
	o := NewInteger(-4711)
	value := o.Float32()
	assert.Equal(t, float32(-4711), value)
}

func TestIntegerFloat64Value(t *testing.T) {
	o := NewInteger(-4711)
	value := o.Float64()
	assert.Equal(t, float64(-4711), value)
}

func TestParseIntegerValue(t *testing.T) {
	o, err := ParseInteger("-83628")
	assert.NotNil(t, o, "value expected")
	assert.Nil(t, err, "no error expected")
	if o != nil {
		assert.Equal(t, int32(-83628), o.Int())
	}
}

func TestParseIntegerValueInvalid(t *testing.T) {
	o, err := ParseInteger("8273.3")
	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}
