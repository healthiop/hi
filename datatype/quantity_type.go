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
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type QuantityComparator CodeAccessor

var (
	LessThanQuantityComparator           QuantityComparator = NewCode("<")
	LessOrEqualThanQuantityComparator    QuantityComparator = NewCode("<=")
	GreaterThanQuantityComparator        QuantityComparator = NewCode(">")
	GreaterOrEqualThanQuantityComparator QuantityComparator = NewCode(">=")
)

var UCUMSystemURI *URIType = NewURI("http://unitsofmeasure.org")

var quantityTypeInfo = newElementTypeInfo("Quantity")

var quantityCodeExpRegexp = regexp.MustCompile("^(.*[^\\d])([1-3])$")

type QuantityType struct {
	value      DecimalAccessor
	comparator QuantityComparator
	unit       StringAccessor
	system     URIAccessor
	code       CodeAccessor
}

type QuantityAccessor interface {
	ElementAccessor
	Stringifier
	DecimalValueAccessor
	Negator
	ArithmeticApplier

	Comparator() QuantityComparator
	Unit() StringAccessor
	System() URIAccessor
	Code() CodeAccessor
}

type QuantityModifier interface {
	QuantityAccessor

	SetValue(value *DecimalType) QuantityModifier
	SetComparator(value QuantityComparator) QuantityModifier
	SetUnit(value *StringType) QuantityModifier
	SetSystem(value *URIType) QuantityModifier
	SetCode(value *CodeType) QuantityModifier
}

func NewQuantityCollection() *CollectionType {
	return NewCollection(quantityTypeInfo)
}

func NewQuantityEmpty() *QuantityType {
	return &QuantityType{}
}

func NewQuantity(value DecimalAccessor, comparator QuantityComparator,
	unit StringAccessor, system URIAccessor, code CodeAccessor) *QuantityType {
	return &QuantityType{
		value:      value,
		comparator: comparator,
		unit:       unit,
		system:     system,
		code:       code,
	}
}

func (t *QuantityType) DataType() DataTypes {
	return QuantityDataType
}

func (t *QuantityType) Empty() bool {
	return t.value == nil &&
		t.comparator == nil &&
		t.unit == nil &&
		t.system == nil &&
		t.code == nil
}

func (t *QuantityType) NilValue() bool {
	v := t.Value()
	return v == nil || v.Nil()
}

func (t *QuantityType) Value() DecimalAccessor {
	return t.value
}

func (t *QuantityType) WithValue(accessor NumberAccessor) DecimalValueAccessor {
	var value DecimalAccessor
	if accessor != nil && !accessor.Nil() {
		value = NewDecimal(accessor.Decimal())
	}
	return NewQuantity(value, t.Comparator(), t.Unit(), t.System(), t.Code())
}

func (t *QuantityType) ArithmeticOpSupported(op ArithmeticOps) bool {
	return op == AdditionOp ||
		op == SubtractionOp ||
		op == MultiplicationOp ||
		op == DivisionOp
}

func (t *QuantityType) Comparator() QuantityComparator {
	return t.comparator
}

func (t *QuantityType) Unit() StringAccessor {
	return t.unit
}

func (t *QuantityType) System() URIAccessor {
	return t.system
}

func (t *QuantityType) Code() CodeAccessor {
	return t.code
}

func (t *QuantityType) SetValue(value *DecimalType) QuantityModifier {
	t.value = value
	return t
}

func (t *QuantityType) SetComparator(value QuantityComparator) QuantityModifier {
	t.comparator = value
	return t
}

func (t *QuantityType) SetUnit(value *StringType) QuantityModifier {
	t.unit = value
	return t
}

func (t *QuantityType) SetSystem(value *URIType) QuantityModifier {
	t.system = value
	return t
}

func (t *QuantityType) SetCode(value *CodeType) QuantityModifier {
	t.code = value
	return t
}

func (t *QuantityType) Negate() Accessor {
	return NewQuantity(t.value.Negate().(DecimalAccessor), t.comparator, t.unit, t.system, t.code)
}

func (e *QuantityType) TypeInfo() TypeInfoAccessor {
	return quantityTypeInfo
}

func (t *QuantityType) Equal(accessor Accessor) bool {
	if accessor == nil {
		return false
	}
	if o, ok := accessor.(QuantityAccessor); !ok {
		return false
	} else {
		return Equal(t.Value(), o.Value()) &&
			Equal(t.Comparator(), o.Comparator()) &&
			Equal(t.Unit(), o.Unit()) &&
			Equal(t.System(), o.System()) &&
			Equal(t.Code(), o.Code())
	}
}

func (t *QuantityType) ValueEqual(accessor Accessor) bool {
	return quantityValueEqual(t, accessor, false)
}

func (t *QuantityType) ValueEquivalent(accessor Accessor) bool {
	return quantityValueEqual(t, accessor, true)
}

func quantityValueEqual(t QuantityAccessor, accessor Accessor, equivalent bool) bool {
	if o, ok := accessor.(QuantityAccessor); ok {
		return Equal(t.Value(), o.Value()) &&
			Equal(t.System(), o.System()) &&
			Equal(t.Code(), o.Code())
	}
	if d, ok := accessor.(DecimalValueAccessor); ok {
		v1 := t.Value()
		if v1 == nil {
			v1 = decimalNil
		}
		v2 := d.Value()
		if v2 == nil {
			v2 = decimalNil
		}
		if equivalent {
			return v1.ValueEquivalent(v2)
		}
		return v1.ValueEqual(v2)
	}
	return false
}

func (t *QuantityType) String() string {
	var b strings.Builder
	b.Grow(32)
	if t.value != nil {
		b.WriteString(t.value.String())
	}
	if t.code != nil {
		if b.Len() > 0 {
			b.WriteByte(' ')
		}
		b.WriteString(t.code.String())
	}
	return b.String()
}

func (t *QuantityType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if t.NilValue() || operand == nil || operand.NilValue() {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("arithmetic operator not supported: %c", op)
	}

	var code CodeAccessor = t.Code()
	if q, ok := operand.(QuantityAccessor); ok {
		var err error
		code, err = mergeQuantityCodes(t, q, op)
		if err != nil {
			return nil, err
		}
	}

	value, _ := t.Value().Calc(operand.Value(), op)
	return NewQuantity(value.Value(), t.Comparator(), t.Unit(), t.System(), code), nil
}

func mergeQuantityCodes(l QuantityAccessor, r QuantityAccessor, op ArithmeticOps) (CodeAccessor, error) {
	if !ValueEqual(l.System(), r.System()) {
		return nil, fmt.Errorf("systems are not equal: %s != %s",
			StringValue(l.System()), StringValue(r.System()))
	}

	leftUnit, leftExp := extractQuantityCodeExp(l.Code())
	rightUnit, rightExp := extractQuantityCodeExp(r.Code())

	if leftUnit != rightUnit {
		return nil, fmt.Errorf("code units are not equal: %s != %s",
			leftUnit, rightUnit)
	}

	if len(leftUnit) == 0 {
		return NewCodeNil(), nil
	}

	exp := leftExp
	switch op {
	case AdditionOp, SubtractionOp:
		if leftExp != rightExp {
			return nil, fmt.Errorf("code units exponents are not equal: %d != %d",
				leftExp, rightExp)
		}
	case MultiplicationOp:
		exp = leftExp + rightExp
	case DivisionOp:
		exp = leftExp - rightExp
	}

	if exp < 1 || exp > 3 {
		return nil, fmt.Errorf("resulting code unit exponent is invalid (must be between 1 and 3): %d", exp)
	}

	if exp == 1 {
		return NewCode(leftUnit), nil
	}
	return NewCode(leftUnit + strconv.FormatInt(int64(exp), 10)), nil
}

func extractQuantityCodeExp(code CodeAccessor) (string, int) {
	if code == nil {
		return "", 0
	}

	value := code.String()
	if len(value) < 2 {
		return value, 1
	}

	parts := quantityCodeExpRegexp.FindStringSubmatch(value)
	if parts == nil {
		return value, 1
	}

	exp, _ := strconv.Atoi(parts[2])
	return parts[1], exp
}
