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

func TestUCUMSystemURI(t *testing.T) {
	assert.Equal(t, "http://unitsofmeasure.org", UCUMSystemURI.String())
}

func TestQuantityImplementsAccessor(t *testing.T) {
	o := NewQuantityEmpty()
	assert.Implements(t, (*QuantityAccessor)(nil), o)
}

func TestQuantityImplementsNegator(t *testing.T) {
	o := NewQuantityEmpty()
	assert.Implements(t, (*Negator)(nil), o)
}

func TestQuantityDataType(t *testing.T) {
	o := NewQuantityEmpty()
	dataType := o.DataType()
	assert.Equal(t, QuantityDataType, dataType)
}

func TestQuantityTypeInfo(t *testing.T) {
	o := NewQuantityEmpty()
	i := o.TypeInfo()
	if assert.NotNil(t, i, "type info expected") {
		assert.Equal(t, "FHIR.Quantity", i.String())
		if assert.NotNil(t, i.FQBaseName(), "base name expected") {
			assert.Equal(t, "FHIR.Element", i.FQBaseName().String())
		}
	}
}

func TestNewQuantityCollection(t *testing.T) {
	c := NewQuantityCollection()
	assert.Equal(t, "FHIR.Quantity", c.ItemTypeInfo().String())
}

func TestEmptyQuantity(t *testing.T) {
	o := NewQuantityEmpty()
	assert.True(t, o.Empty(), "quantity is empty")
	assert.True(t, o.NilValue(), "quantity is empty")
	assert.Nil(t, o.Value())
	assert.Nil(t, o.Comparator())
	assert.Nil(t, o.Unit())
	assert.Nil(t, o.System())
	assert.Nil(t, o.Code())
}

func TestQuantity(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.False(t, o.Empty(), "quantity is not empty")
	assert.False(t, o.NilValue(), "quantity has value")
	if assert.NotNil(t, o.Value()) {
		assert.Equal(t, 47.1, o.Value().Float64())
	}
	if assert.NotNil(t, o.Comparator()) {
		assert.Equal(t, LessOrEqualThanQuantityComparator, o.Comparator())
	}
	if assert.NotNil(t, o.Unit()) {
		assert.Equal(t, "gram", o.Unit().String())
	}
	if assert.NotNil(t, o.System()) {
		assert.Equal(t, "http://unitsofmeasure.org", o.System().String())
	}
	if assert.NotNil(t, o.Code()) {
		assert.Equal(t, "g", o.Code().String())
	}
}

func TestQuantityValueOnly(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), nil, nil, nil, nil)
	assert.False(t, o.Empty(), "quantity is not empty")
	assert.False(t, o.NilValue(), "quantity has value")
}

func TestQuantityNilValue(t *testing.T) {
	o := NewQuantity(NewDecimalNil(), nil, nil, nil, nil)
	assert.False(t, o.Empty(), "quantity is not empty")
	assert.True(t, o.NilValue(), "quantity has nil value")
}

func TestQuantityComparatorOnly(t *testing.T) {
	o := NewQuantity(nil, LessOrEqualThanQuantityComparator, nil, nil, nil)
	assert.False(t, o.Empty(), "quantity is not empty")
	assert.True(t, o.NilValue(), "quantity has no value")
}

func TestQuantityUnitOnly(t *testing.T) {
	o := NewQuantity(nil, nil, NewString("gram"), nil, nil)
	assert.False(t, o.Empty(), "quantity is not empty")
}

func TestQuantitySystemOnly(t *testing.T) {
	o := NewQuantity(nil, nil, nil, UCUMSystemURI, nil)
	assert.False(t, o.Empty(), "quantity is not empty")
}

func TestQuantityCodeOnly(t *testing.T) {
	o := NewQuantity(nil, nil, nil, nil, NewCode("g"))
	assert.False(t, o.Empty(), "quantity is not empty")
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
		assert.Equal(t, "kilogram", o.Unit().String())
	}
}

func TestQuantityWithSystem(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetSystem(NewURI("test"))
	if assert.Same(t, o, n) {
		assert.Equal(t, "test", o.System().String())
	}
}

func TestQuantityWithCode(t *testing.T) {
	o := NewQuantity(NewDecimalFloat64(47.1), LessOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	n := o.SetCode(NewCode("kg"))
	if assert.Same(t, o, n) {
		assert.Equal(t, "kg", o.Code().String())
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
		assert.Equal(t, "gram", n.Unit().String())
	}
	if assert.NotNil(t, n.System()) {
		assert.Equal(t, "http://unitsofmeasure.org", n.System().String())
	}
	if assert.NotNil(t, n.Code()) {
		assert.Equal(t, "g", n.Code().String())
	}
}

func TestQuantityEqualNil(t *testing.T) {
	assert.Equal(t, false, NewQuantityEmpty().Equal(nil))
}

func TestQuantityEqualTypeDiffers(t *testing.T) {
	assert.Equal(t, false, NewQuantityEmpty().Equal(newAccessorMock()))
	assert.Equal(t, false, NewQuantityEmpty().ValueEqual(newAccessorMock()))
	assert.Equal(t, false, NewQuantityEmpty().ValueEquivalent(newAccessorMock()))
}

func TestQuantityEqualEmpty(t *testing.T) {
	assert.Equal(t, true, NewQuantityEmpty().Equal(NewQuantityEmpty()))
	assert.Equal(t, true, NewQuantityEmpty().ValueEqual(NewQuantityEmpty()))
	assert.Equal(t, true, NewQuantityEmpty().ValueEquivalent(NewQuantityEmpty()))
}

func TestQuantityEqual(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, true, q1.Equal(q2))
	assert.Equal(t, true, q1.ValueEqual(q2))
	assert.Equal(t, true, q1.ValueEquivalent(q2))
}

func TestQuantityEqualValueDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.ValueEqual(q2))
	assert.Equal(t, false, q1.ValueEquivalent(q2))
}

func TestQuantityEqualComparatorDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), GreaterOrEqualThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, true, q1.ValueEqual(q2))
	assert.Equal(t, true, q1.ValueEquivalent(q2))
}

func TestQuantityEqualUnitDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("kilogram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, true, q1.ValueEqual(q2))
	assert.Equal(t, true, q1.ValueEquivalent(q2))
}

func TestQuantityEqualSystemDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), NewCode("test"), NewCode("g"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.ValueEqual(q2))
	assert.Equal(t, false, q1.ValueEquivalent(q2))
}

func TestQuantityEqualCodeDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("kg"))
	assert.Equal(t, false, q1.Equal(q2))
	assert.Equal(t, false, q1.ValueEqual(q2))
	assert.Equal(t, false, q1.ValueEquivalent(q2))
}

func TestQuantityEqualDecimal(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(NewDecimalFloat64(47.1)))
	assert.Equal(t, true, q1.ValueEqual(NewDecimalFloat64(47.1)))
	assert.Equal(t, true, q1.ValueEquivalent(NewDecimalFloat64(47.1)))
}

func TestQuantityEqualInteger(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(NewInteger(47)))
	assert.Equal(t, true, q1.ValueEqual(NewInteger(47)))
	assert.Equal(t, true, q1.ValueEquivalent(NewInteger(47)))
}

func TestQuantityEqualNotEqualDecimal(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(NewDecimalFloat64(47.1)))
	assert.Equal(t, false, q1.ValueEqual(NewDecimalFloat64(47.1)))
	assert.Equal(t, false, q1.ValueEquivalent(NewDecimalFloat64(47.1)))
}

func TestQuantityEqualNilInteger(t *testing.T) {
	q1 := NewQuantity(nil, LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(NewIntegerNil()))
	assert.Equal(t, true, q1.ValueEqual(NewIntegerNil()))
	assert.Equal(t, true, q1.ValueEquivalent(NewIntegerNil()))
}

func TestQuantityEqualNilMockNil(t *testing.T) {
	q1 := NewQuantity(nil, LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(newDecimalValueAccessorMock()))
	assert.Equal(t, true, q1.ValueEqual(newDecimalValueAccessorMock()))
	assert.Equal(t, true, q1.ValueEquivalent(newDecimalValueAccessorMock()))
}

func TestQuantityEquivalentDecimal(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.12), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, false, q1.Equal(NewDecimalFloat64(47.1)))
	assert.Equal(t, false, q1.ValueEqual(NewDecimalFloat64(47.1)))
	assert.Equal(t, true, q1.ValueEquivalent(NewDecimalFloat64(47.1)))
}

func TestQuantityStringEmpty(t *testing.T) {
	q := NewQuantity(nil, LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, nil)
	assert.Equal(t, "", q.String())
}

func TestQuantityStringValueOnly(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, nil)
	assert.Equal(t, "47.1", q.String())
}

func TestQuantityString(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, "47.1 g", q.String())
}

func TestQuantityStringCodeOnly(t *testing.T) {
	q := NewQuantity(nil, LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.Equal(t, "g", q.String())
}

func TestQuantityWithValueNil(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	r := q.WithValue(nil)
	e := NewQuantity(nil, LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.True(t, e.Equal(r))
}

func TestQuantityWithValueNilValue(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	r := q.WithValue(NewDecimalNil())
	e := NewQuantity(nil, LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.True(t, e.Equal(r))
}

func TestQuantityWithValueNotNil(t *testing.T) {
	q := NewQuantity(NewDecimalFloat64(47.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	r := q.WithValue(NewDecimalFloat64(12.1))
	e := NewQuantity(NewDecimalFloat64(12.1), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	assert.True(t, e.Equal(r))
}

func TestQuantityArithmeticOpSupported(t *testing.T) {
	d := NewQuantityEmpty()
	assert.True(t, d.ArithmeticOpSupported(AdditionOp), "addition must be supported")
	assert.True(t, d.ArithmeticOpSupported(SubtractionOp), "subtraction must be supported")
	assert.True(t, d.ArithmeticOpSupported(MultiplicationOp), "multiplication must be supported")
	assert.True(t, d.ArithmeticOpSupported(DivisionOp), "division must be supported")
	assert.False(t, d.ArithmeticOpSupported(DivOp), "DIV must not be supported")
	assert.False(t, d.ArithmeticOpSupported(ModOp), "MOD must not be supported")
}

func TestExtractQuantityCodeExpNil(t *testing.T) {
	unit, exp := extractQuantityCodeExp(nil)
	assert.Equal(t, "", unit)
	assert.Equal(t, 0, exp)
}

func TestExtractQuantityCodeExpSingleCharUnit(t *testing.T) {
	unit, exp := extractQuantityCodeExp(NewCode("m"))
	assert.Equal(t, "m", unit)
	assert.Equal(t, 1, exp)
}

func TestExtractQuantityCodeExpMultiCharUnit(t *testing.T) {
	unit, exp := extractQuantityCodeExp(NewCode("cm"))
	assert.Equal(t, "cm", unit)
	assert.Equal(t, 1, exp)
}

func TestExtractQuantityCodeExpExp(t *testing.T) {
	unit, exp := extractQuantityCodeExp(NewCode("m2"))
	assert.Equal(t, "m", unit)
	assert.Equal(t, 2, exp)
}

func TestMergeQuantityCodesSystemNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), nil, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), nil, NewCode("g"))
	code, err := mergeQuantityCodes(q1, q2, AdditionOp)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, NewCode("g"), code)
}

func TestMergeQuantityCodesSystemsDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), NewURI("http://test.com"), NewCode("g"))
	code, err := mergeQuantityCodes(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityCodesCodesDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	code, err := mergeQuantityCodes(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityCodesCodesNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCodeNil())
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, nil)
	code, err := mergeQuantityCodes(q1, q2, AdditionOp)
	assert.NoError(t, err, "no error expected")
	assert.Equal(t, NewCodeNil(), code)
}

func TestMergeQuantityCodesAdditionExpsDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	code, err := mergeQuantityCodes(q1, q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityCodesSubtractionExpsDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	code, err := mergeQuantityCodes(q1, q2, SubtractionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityCodesMultiplicationExpsError(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	code, err := mergeQuantityCodes(q1, q2, MultiplicationOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestMergeQuantityCodesDivisionExpError(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	q2 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	code, err := mergeQuantityCodes(q1, q2, MultiplicationOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, code, "no code expected")
}

func TestQuantityCalcAddition(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(21.7), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	r, err := q1.Calc(q2, AdditionOp)
	e := NewQuantity(NewDecimalFloat64(68.9), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcAdditionValue(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	r, err := q1.Calc(NewInteger(20), AdditionOp)
	e := NewQuantity(NewDecimalFloat64(67.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcSubtraction(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(21.7), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	r, err := q1.Calc(q2, SubtractionOp)
	e := NewQuantity(NewDecimalFloat64(25.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcMultiplication(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(47.2), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(21.7), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	r, err := q1.Calc(q2, MultiplicationOp)
	e := NewQuantity(NewDecimalFloat64(1024.24), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcDivision(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48.75), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m3"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	r, err := q1.Calc(q2, DivisionOp)
	e := NewQuantity(NewDecimalFloat64(19.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m2"))
	assert.NoError(t, err, "no error expected")
	assert.True(t, e.Equal(r))
}

func TestQuantityCalcNotSupportedOp(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48.75), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m3"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	r, err := q1.Calc(q2, DivOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, r, "no result expected")
}

func TestQuantityCalcCodesDiffer(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(48.75), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	r, err := q1.Calc(q2, AdditionOp)
	assert.Error(t, err, "error expected")
	assert.Nil(t, r, "no result expected")
}

func TestQuantityCalcLeftNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalNil(), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalFloat64(2.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	r, err := q1.Calc(q2, AdditionOp)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, r, "empty result expected")
}

func TestQuantityCalcRightNil(t *testing.T) {
	q1 := NewQuantity(NewDecimalFloat64(2.5), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("m"))
	q2 := NewQuantity(NewDecimalNil(), LessThanQuantityComparator,
		NewString("gram"), UCUMSystemURI, NewCode("g"))
	r, err := q1.Calc(q2, AdditionOp)
	assert.NoError(t, err, "no error expected")
	assert.Nil(t, r, "empty result expected")
}
