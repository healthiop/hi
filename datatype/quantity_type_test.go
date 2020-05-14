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

func TestQuantityImplementsAccessor(t *testing.T) {
	o := NewEmptyQuantity()
	assert.Implements(t, (*QuantityAccessor)(nil), o)
}

func TestQuantityImplementsNegator(t *testing.T) {
	o := NewEmptyQuantity()
	assert.Implements(t, (*Negator)(nil), o)
}

func TestQuantityDataType(t *testing.T) {
	o := NewEmptyQuantity()
	dataType := o.DataType()
	assert.Equal(t, QuantityDataType, dataType)
}

func TestEmptyQuantity(t *testing.T) {
	o := NewEmptyQuantity()
	assert.Nil(t, o.Value())
	assert.Nil(t, o.Comparator())
	assert.Nil(t, o.Unit())
	assert.Nil(t, o.System())
	assert.Nil(t, o.Code())
}

func TestQuantity(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	if assert.NotNil(t, o.Value()) {
		assert.Equal(t, 47.1, o.Value().Float64())
	}
	if assert.NotNil(t, o.Comparator()) {
		assert.Equal(t, LessOrEqualThanQuantityComparator, o.Comparator())
	}
	if assert.NotNil(t, o.Unit()) {
		assert.Equal(t, "gram", o.Unit().Value())
	}
	if assert.NotNil(t, o.System()) {
		assert.Equal(t, "http://unitsofmeasure.org", o.System().Value())
	}
	if assert.NotNil(t, o.Code()) {
		assert.Equal(t, "g", o.Code().Value())
	}
}

func TestQuantityWithValue(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetValue(NewDecimalFloat64(-56.8))
	if assert.Same(t, o, n) {
		assert.Equal(t, -56.8, o.Value().Float64())
	}
}

func TestQuantityWithComparator(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetComparator(GreaterThanQuantityComparator)
	if assert.Same(t, o, n) {
		assert.Equal(t, GreaterThanQuantityComparator, o.Comparator())
	}
}

func TestQuantityWithUnit(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetUnit(NewString("kilogram"))
	if assert.Same(t, o, n) {
		assert.Equal(t, "kilogram", o.Unit().Value())
	}
}

func TestQuantityWithSystem(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetSystem(NewURI("test"))
	if assert.Same(t, o, n) {
		assert.Equal(t, "test", o.System().Value())
	}
}

func TestQuantityWithCode(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetCode(NewCode("kg"))
	if assert.Same(t, o, n) {
		assert.Equal(t, "kg", o.Code().Value())
	}
}

func TestQuantityNegate(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.Negate().(QuantityAccessor)
	if assert.NotNil(t, n.Value()) {
		assert.Equal(t, -47.1, n.Value().Float64())
	}
	if assert.NotNil(t, n.Comparator()) {
		assert.Equal(t, LessOrEqualThanQuantityComparator, n.Comparator())
	}
	if assert.NotNil(t, n.Unit()) {
		assert.Equal(t, "gram", n.Unit().Value())
	}
	if assert.NotNil(t, n.System()) {
		assert.Equal(t, "http://unitsofmeasure.org", n.System().Value())
	}
	if assert.NotNil(t, n.Code()) {
		assert.Equal(t, "g", n.Code().Value())
	}
}
