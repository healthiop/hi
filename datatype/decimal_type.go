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
	"github.com/shopspring/decimal"
	"math/big"
)

var decimalTypeInfo = newElementTypeInfo("decimal")

var decimalNil = NewDecimalNil()

type DecimalType struct {
	nilValue bool
	value    decimal.Decimal
}

type DecimalAccessor interface {
	DecimalValueAccessor
	NumberAccessor
}

type DecimalValueAccessor interface {
	Accessor
	Value() DecimalAccessor
}

func NewDecimalCollection() *CollectionType {
	return NewCollection(decimalTypeInfo)
}

func NewDecimalNil() *DecimalType {
	return newDecimal(true, decimal.Zero)
}

func NewDecimalInt(value int32) *DecimalType {
	return newDecimal(false, decimal.NewFromInt32(value))
}

func NewDecimalInt64(value int64) *DecimalType {
	return newDecimal(false, decimal.NewFromInt(value))
}

func NewDecimalFloat64(value float64) *DecimalType {
	return newDecimal(false, decimal.NewFromFloat(value))
}

func ParseDecimal(value string) (*DecimalType, error) {
	if d, err := decimal.NewFromString(value); err != nil {
		return nil, fmt.Errorf("not a decimal: %s", value)
	} else {
		return newDecimal(false, d), nil
	}
}

func newDecimal(nilValue bool, value decimal.Decimal) *DecimalType {
	return &DecimalType{
		nilValue: nilValue,
		value:    value,
	}
}

func (t *DecimalType) DataType() DataTypes {
	return DecimalDataType
}

func (t *DecimalType) Empty() bool {
	return t.Nil()
}

func (t *DecimalType) Nil() bool {
	return t.nilValue
}

func (t *DecimalType) Int() int32 {
	return int32(t.value.IntPart())
}

func (t *DecimalType) Int64() int64 {
	return t.value.IntPart()
}

func (t *DecimalType) Float32() float32 {
	v, _ := t.value.Float64()
	return float32(v)
}

func (t *DecimalType) Float64() float64 {
	v, _ := t.value.Float64()
	return v
}

func (t *DecimalType) BigFloat() *big.Float {
	return t.value.BigFloat()
}

func (t *DecimalType) Decimal() decimal.Decimal {
	return t.value
}

func (t *DecimalType) Value() DecimalAccessor {
	return t
}

func (e *DecimalType) TypeInfo() TypeInfoAccessor {
	return decimalTypeInfo
}

func (t *DecimalType) Negate() Accessor {
	if t.nilValue {
		return t
	}
	return newDecimal(false, t.value.Neg())
}

func (t *DecimalType) Equal(accessor Accessor) bool {
	if accessor.DataType() != DecimalDataType {
		return false
	}
	return t.ValueEqual(accessor)
}

func (t *DecimalType) ValueEqual(accessor Accessor) bool {
	return decimalValueEqual(t, accessor)
}

func decimalValueEqual(t NumberAccessor, accessor Accessor) bool {
	var n NumberAccessor
	if IsNumber(accessor) {
		n = accessor.(NumberAccessor)
	} else if da, ok := accessor.(DecimalValueAccessor); ok {
		n = da.Value()
		if n == nil {
			n = decimalNil
		}
	} else {
		return false
	}

	return t.Nil() == n.Nil() && t.Decimal().Equal(n.Decimal())
}

func (t *DecimalType) ValueEquivalent(accessor Accessor) bool {
	return decimalValueEquivalent(t, accessor)
}

func decimalValueEquivalent(t NumberAccessor, accessor Accessor) bool {
	var n NumberAccessor
	if IsNumber(accessor) {
		n = accessor.(NumberAccessor)
	} else if da, ok := accessor.(DecimalValueAccessor); ok {
		n = da.Value()
		if n == nil {
			n = decimalNil
		}
	} else {
		return false
	}

	d1, d2 := leastPrecisionDecimal(t.Decimal(), n.Decimal())
	return t.Nil() == n.Nil() && d1.Equal(d2)
}

func (t *DecimalType) String() string {
	if t.nilValue {
		return ""
	}

	exp := t.value.Exponent()
	if exp >= 0 {
		return t.value.String()
	}
	return t.value.StringFixed(-exp)
}
