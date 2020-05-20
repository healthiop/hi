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
	"strconv"
)

var integerTypeInfo = newElementTypeInfo("integer")

type IntegerType struct {
	PrimitiveType
	value int32
}

type IntegerAccessor interface {
	NumberAccessor
}

func IsInteger(accessor Accessor) bool {
	dt := accessor.DataType()
	return dt == IntegerDataType ||
		dt == PositiveIntDataType ||
		dt == UnsignedIntDataType
}

func NewIntegerCollection() *CollectionType {
	return NewCollection(integerTypeInfo)
}

func NewIntegerNil() *IntegerType {
	return newInteger(true, 0)
}

func NewInteger(value int32) *IntegerType {
	return newInteger(false, value)
}

func ParseInteger(value string) (*IntegerType, error) {
	if i, err := strconv.Atoi(value); err != nil {
		return nil, fmt.Errorf("not an integer: %s", value)
	} else {
		return NewInteger(int32(i)), nil
	}
}

func newInteger(nilValue bool, value int32) *IntegerType {
	return &IntegerType{
		PrimitiveType: PrimitiveType{
			nilValue: nilValue,
		},
		value: value,
	}
}

func (t *IntegerType) DataType() DataTypes {
	return IntegerDataType
}

func (t *IntegerType) Int() int32 {
	return t.value
}

func (t *IntegerType) Int64() int64 {
	return int64(t.value)
}

func (t *IntegerType) Float32() float32 {
	return float32(t.value)
}

func (t *IntegerType) Float64() float64 {
	return float64(t.value)
}

func (t *IntegerType) BigFloat() *big.Float {
	return big.NewFloat(float64(t.value))
}

func (t *IntegerType) Decimal() decimal.Decimal {
	return decimal.NewFromInt32(t.value)
}

func (t *IntegerType) NilValue() bool {
	return t.Nil()
}

func (t *IntegerType) Value() DecimalAccessor {
	if t.nilValue {
		return NewDecimalNil()
	}
	return NewDecimalInt(t.value)
}

func (t *IntegerType) WithValue(accessor NumberAccessor) DecimalValueAccessor {
	if accessor == nil || IsInteger(accessor) {
		return accessor
	}

	if accessor.Nil() {
		return NewIntegerNil()
	}
	return NewInteger(accessor.Int())
}

func (t *IntegerType) ArithmeticOpSupported(ArithmeticOps) bool {
	return true
}

func (t *IntegerType) Negate() Accessor {
	if t.nilValue {
		return t
	}
	return NewInteger(-t.value)
}

func (t *IntegerType) TypeInfo() TypeInfoAccessor {
	return integerTypeInfo
}

func (t *IntegerType) Equal(accessor Accessor) bool {
	if accessor == nil || !IsInteger(accessor) {
		return false
	}
	return t.ValueEqual(accessor)
}

func (t *IntegerType) ValueEqual(accessor Accessor) bool {
	if IsInteger(accessor) {
		o := accessor.(IntegerAccessor)
		return t.Nil() == o.Nil() && t.Int() == o.Int()
	}

	return decimalValueEqual(t, accessor)
}

func (t *IntegerType) ValueEquivalent(accessor Accessor) bool {
	if IsInteger(accessor) {
		o := accessor.(IntegerAccessor)
		return t.Nil() == o.Nil() && t.Int() == o.Int()
	}

	return decimalValueEquivalent(t, accessor)
}

func (t *IntegerType) String() string {
	if t.nilValue {
		return ""
	}

	return strconv.FormatInt(int64(t.value), 10)
}

func (t *IntegerType) Calc(operand DecimalValueAccessor, op ArithmeticOps) (DecimalValueAccessor, error) {
	if t.Nil() || operand == nil || operand.NilValue() {
		return nil, nil
	}

	if !t.ArithmeticOpSupported(op) || !operand.ArithmeticOpSupported(op) {
		return nil, fmt.Errorf("arithmetic operator not supported: %c", op)
	}

	if IsInteger(operand) {
		operandValue := operand.(IntegerAccessor).Int()
		switch op {
		case AdditionOp:
			return NewInteger(t.Int() + operandValue), nil
		case SubtractionOp:
			return NewInteger(t.Int() - operandValue), nil
		case MultiplicationOp:
			return NewInteger(t.Int() * operandValue), nil
		case DivisionOp:
			if operandValue == 0 {
				return nil, nil
			}
			return NewDecimalFloat64(float64(t.Int()) / float64(operandValue)), nil
		case DivOp:
			if operandValue == 0 {
				return nil, nil
			}
			return NewInteger(t.Int() / operandValue), nil
		case ModOp:
			if operandValue == 0 {
				return nil, nil
			}
			return NewInteger(t.Int() % operandValue), nil
		default:
			panic(fmt.Sprintf("Unhandled operator: %d", op))
		}
	}

	return operand.WithValue(decimalCalc(t, operand.Value(), op)), nil
}
