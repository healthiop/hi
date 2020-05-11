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

func TestBooleanImplementsAccessor(t *testing.T) {
	o := NewBooleanType(false)
	assert.Implements(t, (*BooleanAccessor)(nil), o)
}

func TestBooleanDataType(t *testing.T) {
	o := NewBooleanType(false)
	dataType := o.DataType()
	assert.Equal(t, BooleanDataType, dataType)
}

func TestBooleanValue(t *testing.T) {
	o := NewBooleanType(true)
	value := o.Value()
	assert.Equal(t, true, value)
}

func TestParseBooleanValueTrue(t *testing.T) {
	o, err := ParseBooleanValue("true")

	assert.NotNil(t, o, "value expected")
	assert.Nil(t, err, "no error expected")
}

func TestParseBooleanValueFalse(t *testing.T) {
	o, err := ParseBooleanValue("false")

	assert.NotNil(t, o, "value expected")
	assert.Nil(t, err, "no error expected")
	if o != nil {
		assert.Equal(t, false, o.Value())
	}
}

func TestParseBooleanValueInvalid(t *testing.T) {
	o, err := ParseBooleanValue("0")

	assert.Nil(t, o, "value unexpected")
	assert.NotNil(t, err, "error expected")
}
